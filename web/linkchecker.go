package web

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/html"
)

// LinkChecker validates internal links across the site
type LinkChecker struct {
	markdownService *MarkdownService
	htmlService     *HTMLService
	seoService      *SEOService
	validRoutes     map[string]bool
	publicFiles     map[string]bool
	brokenLinks     map[string][]BrokenLink
	brokenLinksMutex sync.Mutex
	totalLinks      int64
	validLinks      int64
	concurrency     int
	logger          *Logger
	startTime       time.Time
}

// BrokenLink represents a broken link found in the site
type BrokenLink struct {
	URL         string
	SourcePath  string
	LineContext string
}

// NewLinkChecker creates a new link checker instance
func NewLinkChecker(markdownService *MarkdownService, htmlService *HTMLService, seoService *SEOService) *LinkChecker {
	concurrency := runtime.NumCPU()
	if concurrency < 2 {
		concurrency = 2
	}
	if concurrency > 8 {
		concurrency = 8 // Cap at 8 to avoid overwhelming the system
	}
	
	return &LinkChecker{
		markdownService: markdownService,
		htmlService:     htmlService,
		seoService:      seoService,
		validRoutes:     make(map[string]bool),
		publicFiles:     make(map[string]bool),
		brokenLinks:     make(map[string][]BrokenLink),
		concurrency:     concurrency,
	}
}

// FileContent represents content to be processed
type FileContent struct {
	Path    string
	Content string
}

// LinkJob represents a link validation job
type LinkJob struct {
	SourcePath string
	Link       string
}

// CheckAllLinks checks all internal links in the site using concurrent processing
func (lc *LinkChecker) CheckAllLinks() error {
	// Build valid routes from cached content
	lc.buildValidRoutes()

	// Create channels for file processing
	fileJobs := make(chan FileContent, 100)
	linkJobs := make(chan LinkJob, 1000)
	
	var wg sync.WaitGroup
	
	// Start link extraction workers
	for i := 0; i < lc.concurrency; i++ {
		wg.Add(1)
		go lc.linkExtractionWorker(fileJobs, linkJobs, &wg)
	}
	
	// Start link validation workers
	var validationWg sync.WaitGroup
	for i := 0; i < lc.concurrency; i++ {
		validationWg.Add(1)
		go lc.linkValidationWorker(linkJobs, &validationWg)
	}

	// Send all markdown content for processing
	go func() {
		defer close(fileJobs)
		
		// Process markdown content
		for path, content := range lc.markdownService.GetAllCachedContent() {
			fileJobs <- FileContent{Path: path, Content: content.HTML}
		}
		
		// Process HTML content
		for path, content := range lc.htmlService.GetAllCachedContent() {
			fileJobs <- FileContent{Path: path, Content: content.HTML}
		}
	}()
	
	// Wait for link extraction to complete
	wg.Wait()
	close(linkJobs)
	
	// Wait for link validation to complete
	validationWg.Wait()

	// Print results
	lc.printResults()

	return nil
}

// linkExtractionWorker extracts links from file content concurrently
func (lc *LinkChecker) linkExtractionWorker(fileJobs <-chan FileContent, linkJobs chan<- LinkJob, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for fileContent := range fileJobs {
		// Extract HTML links
		htmlLinks := lc.extractLinksFromHTML(fileContent.Content)
		
		// Extract markdown links
		mdLinks := lc.extractLinksFromMarkdown(fileContent.Content)
		
		// Combine all links
		allLinks := append(htmlLinks, mdLinks...)
		
		// Send each link for validation
		for _, link := range allLinks {
			linkJobs <- LinkJob{
				SourcePath: fileContent.Path,
				Link:       link,
			}
		}
	}
}

// linkValidationWorker validates links concurrently
func (lc *LinkChecker) linkValidationWorker(linkJobs <-chan LinkJob, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range linkJobs {
		lc.validateSingleLink(job.SourcePath, job.Link)
	}
}

// buildValidRoutes builds a map of all valid routes
func (lc *LinkChecker) buildValidRoutes() {
	// Add all cached markdown routes
	for path := range lc.markdownService.GetAllCachedContent() {
		lc.validRoutes[path] = true
		// Also add with trailing slash
		if !strings.HasSuffix(path, "/") {
			lc.validRoutes[path+"/"] = true
		}
	}

	// Add all cached HTML routes
	for path := range lc.htmlService.GetAllCachedContent() {
		lc.validRoutes[path] = true
		// Also add with trailing slash
		if !strings.HasSuffix(path, "/") {
			lc.validRoutes[path+"/"] = true
		}
	}

	// Explicitly add common routes that might be handled specially by router
	lc.validRoutes["/"] = true

	// Add common API and special routes
	lc.validRoutes["/health"] = true
	lc.validRoutes["/api/status/current"] = true
	lc.validRoutes["/api/status/history"] = true
	lc.validRoutes["/platform/status"] = true

	// Scan public directory for static files
	lc.scanPublicFiles()
}

// scanPublicFiles scans the public directory and adds all files as valid routes
func (lc *LinkChecker) scanPublicFiles() {
	err := filepath.WalkDir("public", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // Continue on errors
		}

		if !d.IsDir() {
			// Convert file path to URL path
			urlPath := "/" + path
			lc.publicFiles[urlPath] = true
		}

		return nil
	})

	if err != nil {
		// Fallback to common files if scanning fails
		lc.publicFiles["/favicon.ico"] = true
		lc.publicFiles["/css/style.css"] = true
	}
}

// extractLinksFromHTML extracts links from HTML content
func (lc *LinkChecker) extractLinksFromHTML(htmlContent string) []string {
	var links []string
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return links
	}

	var extract func(*html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				for _, attr := range n.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			case "img":
				for _, attr := range n.Attr {
					if attr.Key == "src" && !strings.HasPrefix(attr.Val, "data:") {
						links = append(links, attr.Val)
					}
				}
			case "video":
				for _, attr := range n.Attr {
					if attr.Key == "src" {
						links = append(links, attr.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(doc)

	return links
}

// extractLinksFromMarkdown extracts markdown-style links
func (lc *LinkChecker) extractLinksFromMarkdown(content string) []string {
	var links []string

	// Match markdown links [text](url)
	mdLinkRegex := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	matches := mdLinkRegex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) > 2 {
			links = append(links, match[2])
		}
	}

	// Match markdown images ![alt](url)
	mdImageRegex := regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	imageMatches := mdImageRegex.FindAllStringSubmatch(content, -1)
	for _, match := range imageMatches {
		if len(match) > 2 {
			links = append(links, match[2])
		}
	}

	return links
}

// validateSingleLink validates a single link (thread-safe version)
func (lc *LinkChecker) validateSingleLink(sourcePath, link string) {
	// Increment total links counter
	atomic.AddInt64(&lc.totalLinks, 1)

	// Skip external links
	if strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") || strings.HasPrefix(link, "//") {
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Skip mailto and other protocols
	if strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "tel:") || strings.HasPrefix(link, "javascript:") {
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Skip plain email addresses (without mailto: prefix)
	if strings.Contains(link, "@") && !strings.Contains(link, "/") {
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Skip empty links and anchors on same page
	if link == "" || link == "#" {
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Parse the link
	u, err := url.Parse(link)
	if err != nil {
		lc.addBrokenLink(sourcePath, link, "Invalid URL format")
		return
	}

	// Handle anchor-only links
	if strings.HasPrefix(link, "#") {
		// TODO: Validate anchor exists in current page
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Get the path without fragment and query
	checkPath := u.Path
	if checkPath == "" {
		checkPath = "/"
	}

	// Check if it's a valid route
	if lc.validRoutes[checkPath] {
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Check if it's a public file
	if strings.HasPrefix(checkPath, "/") || lc.publicFiles[checkPath] {
		atomic.AddInt64(&lc.validLinks, 1)
		return
	}

	// Check for redirects
	if redirectTo, _, shouldRedirect := lc.seoService.CheckRedirect(checkPath); shouldRedirect {
		// Link redirects to valid location
		if lc.validRoutes[redirectTo] {
			atomic.AddInt64(&lc.validLinks, 1)
			return
		}
	}

	// Link is broken
	lc.addBrokenLink(sourcePath, link, "Page not found")
}

// addBrokenLink adds a broken link to the results (thread-safe)
func (lc *LinkChecker) addBrokenLink(sourcePath, link, context string) {
	brokenLink := BrokenLink{
		URL:         link,
		SourcePath:  sourcePath,
		LineContext: context,
	}

	lc.brokenLinksMutex.Lock()
	defer lc.brokenLinksMutex.Unlock()
	
	if _, exists := lc.brokenLinks[link]; !exists {
		lc.brokenLinks[link] = []BrokenLink{}
	}
	lc.brokenLinks[link] = append(lc.brokenLinks[link], brokenLink)
}

// printResults prints the link checking results and generates CSV
func (lc *LinkChecker) printResults() {
	// Get counters from atomic variables
	total := atomic.LoadInt64(&lc.totalLinks)
	valid := atomic.LoadInt64(&lc.validLinks)
	brokenCount := total - valid

	if brokenCount == 0 {
		lc.logger.Log(LogCheck, "✅", "Completed", fmt.Sprintf("%d/%d links valid", valid, total), time.Since(lc.startTime))
	} else {
		lc.logger.Log(LogCheck, "✅", "Completed", fmt.Sprintf("%d/%d links valid, %d broken", valid, total, brokenCount), time.Since(lc.startTime))

		// Generate CSV file
		lc.generateCSV()
	}
}

// generateCSV creates a CSV file with broken links and their sources
func (lc *LinkChecker) generateCSV() {
	file, err := os.Create("404.csv")
	if err != nil {
		log.Printf("Failed to create 404.csv: %v", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Broken URL", "Source Page", "References"}
	if err := writer.Write(header); err != nil {
		log.Printf("Failed to write CSV header: %v", err)
		return
	}

	// Thread-safe access to broken links
	lc.brokenLinksMutex.Lock()
	brokenLinksMap := make(map[string][]BrokenLink)
	for url, links := range lc.brokenLinks {
		brokenLinksMap[url] = links
	}
	lc.brokenLinksMutex.Unlock()

	// Sort broken links by URL for consistent output
	var brokenURLs []string
	for url := range brokenLinksMap {
		brokenURLs = append(brokenURLs, url)
	}
	sort.Strings(brokenURLs)

	// Write each broken link with its sources
	for _, url := range brokenURLs {
		sources := brokenLinksMap[url]

		// Group sources and count references
		sourceCount := make(map[string]int)
		for _, source := range sources {
			sourceCount[source.SourcePath]++
		}

		// Write one row per unique source page
		for sourcePage, count := range sourceCount {
			record := []string{
				url,
				sourcePage,
				fmt.Sprintf("%d", count),
			}
			if err := writer.Write(record); err != nil {
				log.Printf("Failed to write CSV record: %v", err)
			}
		}
	}

	// CSV export notification removed for cleaner output
}

// RunLinkChecker runs the link checker with the provided services
func RunLinkChecker(markdownService *MarkdownService, htmlService *HTMLService, seoService *SEOService, logger *Logger) error {
	checker := NewLinkChecker(markdownService, htmlService, seoService)
	checker.logger = logger
	checker.startTime = time.Now()
	return checker.CheckAllLinks()
}
