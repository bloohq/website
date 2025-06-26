package web

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// NavItem represents a navigation item
type NavItem struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Href     string    `json:"href,omitempty"`
	Expanded bool      `json:"expanded,omitempty"`
	Children []NavItem `json:"children,omitempty"`
}

// Navigation holds the complete navigation structure
type Navigation struct {
	Sections []NavItem `json:"sections"`
	Legal    []NavItem `json:"legal"`
}

// PageMetadata represents metadata for a specific page
type PageMetadata struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

// SiteMetadata represents global site metadata
type SiteMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Language    string `json:"language"`
	Author      string `json:"author"`
}

// MetadataDefaults represents default metadata values
type MetadataDefaults struct {
	TitleSuffix string   `json:"title_suffix"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

// Metadata holds the complete metadata structure
type Metadata struct {
	Site     SiteMetadata                `json:"site"`
	Pages    map[string]PageMetadata     `json:"pages"`
	Defaults MetadataDefaults            `json:"defaults"`
}

// Router handles file-based routing for HTML pages
type Router struct {
	pagesDir      string
	layoutsDir    string
	componentsDir string
	navigation    *Navigation
	metadata      *Metadata
}

// NewRouter creates a new router instance
func NewRouter(pagesDir string) *Router {
	router := &Router{
		pagesDir:      pagesDir,
		layoutsDir:    "layouts",
		componentsDir: "components",
	}
	
	// Load navigation data
	if err := router.loadNavigation(); err != nil {
		log.Printf("Error loading navigation: %v", err)
	}
	
	// Load metadata
	if err := router.loadMetadata(); err != nil {
		log.Printf("Error loading metadata: %v", err)
	}
	
	return router
}

// loadNavigation loads navigation data from JSON file
func (r *Router) loadNavigation() error {
	data, err := os.ReadFile("data/nav.json")
	if err != nil {
		return err
	}
	
	r.navigation = &Navigation{}
	return json.Unmarshal(data, r.navigation)
}

// loadMetadata loads metadata from JSON file
func (r *Router) loadMetadata() error {
	data, err := os.ReadFile("data/metadata.json")
	if err != nil {
		return err
	}
	
	r.metadata = &Metadata{}
	return json.Unmarshal(data, r.metadata)
}

// PageData holds data for template rendering
type PageData struct {
	Title       string
	Content     template.HTML
	Navigation  *Navigation
	PageMeta    *PageMetadata
	SiteMeta    *SiteMetadata
	Description string
	Keywords    []string
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Skip public file requests
	if strings.HasPrefix(req.URL.Path, "/public/") {
		return
	}

	// Get the requested path
	path := req.URL.Path
	
	// Redirect .html URLs to clean URLs
	if strings.HasSuffix(path, ".html") && path != "/" {
		cleanURL := strings.TrimSuffix(path, ".html")
		http.Redirect(w, req, cleanURL, http.StatusMovedPermanently)
		return
	}
	
	// Convert URL path to file path
	var filePath string
	if path == "/" {
		// Root path maps to index.html
		filePath = filepath.Join(r.pagesDir, "index.html")
	} else {
		// Remove leading slash
		cleanPath := strings.TrimPrefix(path, "/")
		
		if strings.HasSuffix(cleanPath, "/") {
			// Directory path, look for index.html
			filePath = filepath.Join(r.pagesDir, cleanPath, "index.html")
		} else {
			// Try as direct .html file first
			filePath = filepath.Join(r.pagesDir, cleanPath+".html")
			
			// If that doesn't exist, try as directory with index.html
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				indexPath := filepath.Join(r.pagesDir, cleanPath, "index.html")
				if _, err := os.Stat(indexPath); err == nil {
					// Redirect to trailing slash version
					http.Redirect(w, req, path+"/", http.StatusMovedPermanently)
					return
				}
			}
		}
	}

	// Check if page file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.NotFound(w, req)
		return
	}

	// Read the page content
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error reading page", http.StatusInternalServerError)
		log.Printf("Page reading error: %v", err)
		return
	}

	// Prepare template files - start with layout
	templateFiles := []string{
		filepath.Join(r.layoutsDir, "main.html"),
	}
	
	// Auto-scan all component templates
	componentFiles, err := filepath.Glob(filepath.Join(r.componentsDir, "*.html"))
	if err != nil {
		http.Error(w, "Error loading components", http.StatusInternalServerError)
		log.Printf("Component scanning error: %v", err)
		return
	}
	
	// Add all component files
	templateFiles = append(templateFiles, componentFiles...)

	// Create template with custom functions
	tmpl := template.New("main.html").Funcs(template.FuncMap{
		"toJSON": func(v any) template.JS {
			data, _ := json.Marshal(v)
			return template.JS(data)
		},
	})
	
	// Parse all template files  
	tmpl, err = tmpl.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		log.Printf("Template parsing error: %v", err)
		return
	}

	// Prepare page data
	pageData := r.preparePageData(path, template.HTML(contentBytes))

	// Set content type and execute main layout
	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.ExecuteTemplate(w, "main.html", pageData); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
		log.Printf("Template execution error: %v", err)
		return
	}
}

// preparePageData creates PageData with metadata for the given path
func (r *Router) preparePageData(path string, content template.HTML) PageData {
	// Get page key for metadata lookup
	pageKey := r.getPageKey(path)
	
	// Get page metadata
	var pageMeta *PageMetadata
	var title, description string
	var keywords []string
	
	if r.metadata != nil {
		// Check if specific page metadata exists
		if meta, exists := r.metadata.Pages[pageKey]; exists {
			pageMeta = &meta
			title = meta.Title
			description = meta.Description
			keywords = meta.Keywords
		} else {
			// Use defaults
			title = r.getFallbackTitle(path) + r.metadata.Defaults.TitleSuffix
			description = r.metadata.Defaults.Description
			keywords = r.metadata.Defaults.Keywords
		}
	} else {
		// Fallback if no metadata loaded
		title = r.getFallbackTitle(path)
		description = "Blue - Powerful platform to create, manage, and scale processes for modern teams."
		keywords = []string{"blue", "process management", "team collaboration"}
	}
	
	var siteMeta *SiteMetadata
	if r.metadata != nil {
		siteMeta = &r.metadata.Site
	}
	
	return PageData{
		Title:       title,
		Content:     content,
		Navigation:  r.navigation,
		PageMeta:    pageMeta,
		SiteMeta:    siteMeta,
		Description: description,
		Keywords:    keywords,
	}
}

// getPageKey converts URL path to metadata key
func (r *Router) getPageKey(path string) string {
	if path == "/" {
		return "home"
	}
	
	// Remove leading/trailing slashes
	cleanPath := strings.Trim(path, "/")
	return cleanPath
}

// getFallbackTitle creates a fallback title from URL path
func (r *Router) getFallbackTitle(path string) string {
	if path == "/" {
		return "Home"
	}
	
	// Remove leading slash and convert to title case
	cleanPath := strings.TrimPrefix(path, "/")
	cleanPath = strings.TrimSuffix(cleanPath, "/")
	
	// Replace slashes with spaces and title case
	parts := strings.Split(cleanPath, "/")
	for i, part := range parts {
		// Simple title case replacement for strings.Title
		words := strings.Split(strings.ReplaceAll(part, "-", " "), " ")
		for j, word := range words {
			if len(word) > 0 {
				words[j] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
		}
		parts[i] = strings.Join(words, " ")
	}
	
	return strings.Join(parts, " - ")
}