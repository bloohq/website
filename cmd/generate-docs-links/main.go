package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type DocLink struct {
	Title string
	URL   string
	Path  string
}

func main() {
	baseURL := "https://blue.cc"
	contentDir := "../../content/en"
	outputFile := "documentation-links.md"

	var allLinks []string

	// User Documentation
	docsLinks := scanDirectory(filepath.Join(contentDir, "docs"), baseURL, "/docs/")
	if len(docsLinks) > 0 {
		allLinks = append(allLinks, "## User Documentation")
		allLinks = append(allLinks, formatLinks(docsLinks)...)
		allLinks = append(allLinks, "")
	}

	// API Documentation
	apiLinks := scanDirectory(filepath.Join(contentDir, "api"), baseURL, "/api/")
	if len(apiLinks) > 0 {
		allLinks = append(allLinks, "## API Documentation")
		allLinks = append(allLinks, formatLinks(apiLinks)...)
		allLinks = append(allLinks, "")
	}

	// Legal Documents
	legalLinks := scanDirectory(filepath.Join(contentDir, "legal"), baseURL, "/legal/")
	if len(legalLinks) > 0 {
		allLinks = append(allLinks, "## Legal Documents")
		allLinks = append(allLinks, formatLinks(legalLinks)...)
		allLinks = append(allLinks, "")
	}

	// Platform Pages (hardcoded as they're likely HTML pages)
	allLinks = append(allLinks, "## Platform Pages")
	allLinks = append(allLinks, "- Status Page: https://blue.cc/platform/status")
	allLinks = append(allLinks, "- Changelog: https://blue.cc/platform/changelog")
	allLinks = append(allLinks, "- Roadmap: https://blue.cc/platform/roadmap")
	allLinks = append(allLinks, "")

	// Write to file
	content := "# Blue Documentation Links\n\n" + strings.Join(allLinks, "\n")
	err := os.WriteFile(outputFile, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated %s with %d sections\n", outputFile, 4)
}

func scanDirectory(dir string, baseURL string, urlPrefix string) []DocLink {
	var links []DocLink

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		// Get relative path from content dir
		relPath := strings.TrimPrefix(path, dir)
		relPath = strings.TrimPrefix(relPath, "/")
		
		// Convert to URL path
		urlPath := strings.TrimSuffix(relPath, ".md")
		urlPath = strings.ReplaceAll(urlPath, "/index", "")
		
		// Remove numeric prefixes (e.g., "1.introduction" -> "introduction")
		parts := strings.Split(urlPath, "/")
		for i, part := range parts {
			if dotIndex := strings.Index(part, "."); dotIndex > 0 {
				// Check if before dot is a number
				if isNumericPrefix(part[:dotIndex]) {
					parts[i] = part[dotIndex+1:]
				}
			}
			// Replace spaces with hyphens in URL paths
			parts[i] = strings.ReplaceAll(parts[i], " ", "-")
		}
		urlPath = strings.Join(parts, "/")

		// Create title from filename
		title := filepath.Base(path)
		title = strings.TrimSuffix(title, ".md")
		if dotIndex := strings.Index(title, "."); dotIndex > 0 && isNumericPrefix(title[:dotIndex]) {
			title = title[dotIndex+1:]
		}
		title = strings.ReplaceAll(title, "-", " ")
		title = strings.ReplaceAll(title, "_", " ")
		title = toTitleCase(title)

		// Build full URL
		fullURL := baseURL + urlPrefix + urlPath

		links = append(links, DocLink{
			Title: title,
			URL:   fullURL,
			Path:  relPath,
		})

		return nil
	})

	// Sort links by path for consistent ordering
	sort.Slice(links, func(i, j int) bool {
		return links[i].Path < links[j].Path
	})

	return links
}

func formatLinks(links []DocLink) []string {
	var formatted []string
	for _, link := range links {
		formatted = append(formatted, fmt.Sprintf("- %s: %s", link.Title, link.URL))
	}
	return formatted
}

func isNumericPrefix(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return len(s) > 0
}

func toTitleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			// Special cases for common acronyms
			upperWord := strings.ToUpper(word)
			if upperWord == "API" || upperWord == "CSV" || upperWord == "ID" || upperWord == "URL" {
				words[i] = upperWord
			} else {
				words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			}
		}
	}
	return strings.Join(words, " ")
}