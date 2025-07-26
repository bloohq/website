package web

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// NavItem represents a navigation item
type NavItem struct {
	ID         string      `json:"id"`
	Name       interface{} `json:"name"` // Can be string or map[string]string for multi-language
	Href       string      `json:"href,omitempty"`
	External   bool        `json:"external,omitempty"`
	Expanded   bool        `json:"expanded,omitempty"`
	Children   []NavItem   `json:"children,omitempty"`
	OriginalID string      `json:"-"` // Store original directory name for sorting (not sent to JSON)
}

// Navigation holds the complete navigation structure
type Navigation struct {
	Sections []NavItem `json:"sections"`
}

// DirMetadata represents _dir.yml file content
type DirMetadata struct {
	Title string `yaml:"title"`
}

// NavigationService handles all navigation-related operations
type NavigationService struct {
	navigation *Navigation
	seoService *SEOService
}

// GetLocalizedName extracts the correct language name from a NavItem's Name field
func (ns *NavigationService) GetLocalizedName(nameField interface{}, lang string) string {
	if nameField == nil {
		return ""
	}

	// If it's already a string (for dynamic content), return as-is
	if str, ok := nameField.(string); ok {
		return str
	}

	// If it's a map (from nav.json with language objects)
	if nameMap, ok := nameField.(map[string]interface{}); ok {
		// Try the requested language first
		if name, exists := nameMap[lang]; exists {
			if str, ok := name.(string); ok {
				return str
			}
		}
		// Fallback to English
		if name, exists := nameMap[DefaultLanguage]; exists {
			if str, ok := name.(string); ok {
				return str
			}
		}
		// Fallback to first available language
		for _, name := range nameMap {
			if str, ok := name.(string); ok {
				return str
			}
		}
	}

	return ""
}

// LocalizeNavigation creates a localized copy of navigation for a specific language
func (ns *NavigationService) LocalizeNavigation(nav *Navigation, lang string) *Navigation {
	if nav == nil {
		return nil
	}

	localizedNav := &Navigation{
		Sections: make([]NavItem, len(nav.Sections)),
	}

	for i, section := range nav.Sections {
		localizedNav.Sections[i] = ns.localizeNavItem(section, lang)
	}

	return localizedNav
}

// localizeNavItem recursively localizes a NavItem and its children
func (ns *NavigationService) localizeNavItem(item NavItem, lang string) NavItem {
	localizedItem := NavItem{
		ID:         item.ID,
		Name:       ns.GetLocalizedName(item.Name, lang), // Convert to localized string
		Href:       item.Href,
		External:   item.External,
		Expanded:   item.Expanded,
		OriginalID: item.OriginalID,
	}

	// Localize children recursively
	if len(item.Children) > 0 {
		localizedItem.Children = make([]NavItem, len(item.Children))
		for i, child := range item.Children {
			localizedItem.Children[i] = ns.localizeNavItem(child, lang)
		}
	}

	return localizedItem
}

// NewNavigationService creates a new navigation service
func NewNavigationService(seoService *SEOService) *NavigationService {
	service := &NavigationService{
		seoService: seoService,
	}

	// Load static navigation (now includes docs, API, and legal sections)
	if err := service.LoadNavigation(); err != nil {
		log.Printf("Error loading navigation: %v", err)
	}

	return service
}

// LoadNavigation loads navigation data from JSON file
func (ns *NavigationService) LoadNavigation() error {
	data, err := os.ReadFile("data/nav.json")
	if err != nil {
		return err
	}

	ns.navigation = &Navigation{}
	return json.Unmarshal(data, ns.navigation)
}

// Note: Dynamic navigation generation removed - docs, API, and legal sections
// are now hardcoded in nav.json for better translation support

// GenerateContentNavigation creates navigation tree from content directory
func (ns *NavigationService) GenerateContentNavigation(contentDir, baseURL string) (*Navigation, error) {
	var sections []NavItem

	// Read the content directory
	entries, err := os.ReadDir(contentDir)
	if err != nil {
		return nil, err
	}

	// Process each directory/file
	for _, entry := range entries {
		if entry.IsDir() {
			navItem, err := ns.processDirectory(filepath.Join(contentDir, entry.Name()), entry.Name(), baseURL, contentDir)
			if err != nil {
				log.Printf("Error processing directory %s: %v", entry.Name(), err)
				continue
			}
			if navItem != nil {
				sections = append(sections, *navItem)
			}
		} else if strings.HasSuffix(entry.Name(), ".md") {
			// Handle individual markdown files at root level
			fileName := strings.TrimSuffix(entry.Name(), ".md")
			fileTitle := CleanTitle(fileName)

			// Try to get title from frontmatter
			if filePath := filepath.Join(contentDir, entry.Name()); filePath != "" {
				if data, err := os.ReadFile(filePath); err == nil {
					if frontmatter, _, err := ns.seoService.ParseFrontmatter(data); err == nil && frontmatter != nil && frontmatter.Title != "" {
						fileTitle = frontmatter.Title
					}
				}
			}

			href := baseURL + "/" + CleanID(fileName)

			// Extract content type for unique ID
			pathParts := strings.Split(contentDir, "/")
			contentType := ""
			if len(pathParts) >= 2 {
				contentType = pathParts[1]
			}

			sections = append(sections, NavItem{
				ID:         contentType + "-" + CleanID(fileName),
				Name:       fileTitle,
				Href:       href,
				OriginalID: fileName,
			})
		}
	}

	// Sort sections by numeric prefix
	ns.sortNavItems(sections)

	return &Navigation{Sections: sections}, nil
}

// processDirectory processes a content directory and creates NavItem
func (ns *NavigationService) processDirectory(dirPath, dirName, baseURL, rootContentDir string) (*NavItem, error) {
	// Read directory metadata
	title := CleanTitle(dirName)
	dirMetaPath := filepath.Join(dirPath, "_dir.yml")
	if data, err := os.ReadFile(dirMetaPath); err == nil {
		var dirMeta DirMetadata
		if err := yaml.Unmarshal(data, &dirMeta); err == nil && dirMeta.Title != "" {
			title = dirMeta.Title
		}
	}

	// Create nav item with path-based ID for uniqueness
	// Extract content type from path (e.g., "docs", "api")
	pathParts := strings.Split(rootContentDir, "/")
	contentType := ""
	if len(pathParts) >= 2 {
		contentType = pathParts[1]
	}

	// Create unique ID by prefixing with content type
	uniqueID := contentType + "-" + CleanID(dirName)

	navItem := &NavItem{
		ID:         uniqueID,
		Name:       title,
		Expanded:   false,
		Children:   []NavItem{}, // Always initialize to prevent undefined errors
		OriginalID: dirName,
	}

	// Read directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return navItem, nil
	}

	var children []NavItem

	// Process subdirectories and files
	for _, entry := range entries {
		if entry.Name() == "_dir.yml" {
			continue
		}

		if entry.IsDir() {
			// Recursive subdirectory
			childNav, err := ns.processDirectory(filepath.Join(dirPath, entry.Name()), entry.Name(), baseURL, rootContentDir)
			if err != nil {
				log.Printf("Error processing subdirectory %s: %v", entry.Name(), err)
				continue
			}
			if childNav != nil {
				children = append(children, *childNav)
			}
		} else if strings.HasSuffix(entry.Name(), ".md") {
			// Markdown file
			fileName := strings.TrimSuffix(entry.Name(), ".md")
			fileTitle := CleanTitle(fileName)

			// Try to get title from frontmatter
			if filePath := filepath.Join(dirPath, entry.Name()); filePath != "" {
				if data, err := os.ReadFile(filePath); err == nil {
					if frontmatter, _, err := ns.seoService.ParseFrontmatter(data); err == nil && frontmatter != nil && frontmatter.Title != "" {
						fileTitle = frontmatter.Title
					}
				}
			}

			// Create relative path for href
			relDir := strings.TrimPrefix(dirPath, rootContentDir+"/")

			// Get content type from root directory
			contentTypeName := strings.Split(rootContentDir, "/")[1] // e.g., "content/docs" -> "docs"

			// Remove the content type prefix
			if strings.HasPrefix(relDir, contentTypeName+"/") {
				relDir = strings.TrimPrefix(relDir, contentTypeName+"/")
			} else if contentTypeName == "api" && strings.HasPrefix(relDir, "api/") {
				relDir = strings.TrimPrefix(relDir, "api/")
			} else if contentTypeName == "legal" && strings.HasPrefix(relDir, "legal/") {
				relDir = strings.TrimPrefix(relDir, "legal/")
			}

			// Clean numeric prefixes from directory path
			relDir = CleanDirectoryPath(relDir)

			href := baseURL + "/" + relDir + "/" + CleanID(fileName)

			// Create unique ID using parent directory context
			fileID := uniqueID + "-" + CleanID(fileName)

			children = append(children, NavItem{
				ID:         fileID,
				Name:       fileTitle,
				Href:       href,
				OriginalID: fileName,
			})
		}
	}

	// Sort children by numeric prefix
	ns.sortNavItems(children)

	if len(children) > 0 {
		navItem.Children = children
	}

	return navItem, nil
}

// sortNavItems sorts navigation items by numeric prefix
func (ns *NavigationService) sortNavItems(items []NavItem) {
	// Sort by extracting numeric prefix from original names
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			// Extract numeric prefixes for comparison using original IDs
			num1 := ExtractNumericPrefix(items[i].OriginalID)
			num2 := ExtractNumericPrefix(items[j].OriginalID)

			if num1 > num2 {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
}

// GetNavigationForPath returns the appropriate navigation based on the URL path
func (ns *NavigationService) GetNavigationForPath(path string) *Navigation {
	return ns.GetNavigationForPathWithLanguage(path, DefaultLanguage)
}

// GetNavigationForPathWithLanguage returns the appropriate navigation based on the URL path and language
func (ns *NavigationService) GetNavigationForPathWithLanguage(path, lang string) *Navigation {
	// Return localized static navigation (now includes docs, API, and legal sections from nav.json)
	if ns.navigation == nil {
		return &Navigation{}
	}

	// Create a localized copy of the static navigation and set expansion state based on current path
	nav := ns.LocalizeNavigation(ns.navigation, lang)

	// Set expansion state for sections based on current path
	if nav != nil && nav.Sections != nil {
		for i := range nav.Sections {
			section := &nav.Sections[i]
			switch section.ID {
			case "documentation":
				section.Expanded = strings.HasPrefix(path, "/docs")
			case "api-reference":
				section.Expanded = strings.HasPrefix(path, "/api")
			case "legal":
				section.Expanded = strings.HasPrefix(path, "/legal")
			}
		}
	}

	return nav
}

// GetFirstItemInDirectory finds the first item (by numeric prefix) in a directory
// Returns the URL of the first item, or empty string if not found
func (ns *NavigationService) GetFirstItemInDirectory(path string) string {
	// Remove trailing slash and determine content type
	cleanPath := strings.TrimSuffix(path, "/")
	contentType, isContentPath := GetContentTypeFromPath(cleanPath)
	if !isContentPath {
		return ""
	}

	// Find the appropriate section in the main navigation
	if ns.navigation == nil {
		return ""
	}

	var targetSection *NavItem
	sectionID := ""
	switch contentType.Name {
	case "docs":
		sectionID = "documentation"
	case "api":
		sectionID = "api-reference"
	case "legal":
		sectionID = "legal"
	default:
		return ""
	}

	// Find the section in the main navigation
	for i := range ns.navigation.Sections {
		if ns.navigation.Sections[i].ID == sectionID {
			targetSection = &ns.navigation.Sections[i]
			break
		}
	}

	if targetSection == nil || len(targetSection.Children) == 0 {
		return ""
	}

	// If it's a root content type path (e.g., /docs, /api, /legal)
	if cleanPath == contentType.URLPrefix {
		// Return the first section's first item
		if len(targetSection.Children) > 0 {
			firstSection := targetSection.Children[0]
			if firstSection.Href != "" {
				return firstSection.Href
			}
			// If section has children, return first child
			if len(firstSection.Children) > 0 {
				return firstSection.Children[0].Href
			}
		}
		return ""
	}

	// For deeper paths, we need to navigate the tree
	// Remove the content type prefix to get the relative path
	relativePath := strings.TrimPrefix(cleanPath, contentType.URLPrefix)
	relativePath = strings.TrimPrefix(relativePath, "/")

	// Find the matching directory in the navigation tree
	return ns.findFirstItemInPath(targetSection.Children, relativePath, contentType.URLPrefix)
}

// findFirstItemInPath recursively searches the navigation tree for a matching path
func (ns *NavigationService) findFirstItemInPath(sections []NavItem, targetPath string, baseURL string) string {
	if targetPath == "" {
		// We're at the target directory, return first item
		if len(sections) > 0 {
			firstItem := sections[0]
			if firstItem.Href != "" {
				return firstItem.Href
			}
			// If first item has children, return first child
			if len(firstItem.Children) > 0 {
				return firstItem.Children[0].Href
			}
		}
		return ""
	}

	// Split the path to get the first segment
	pathParts := strings.Split(targetPath, "/")
	if len(pathParts) == 0 {
		return ""
	}

	firstSegment := pathParts[0]
	remainingPath := strings.Join(pathParts[1:], "/")

	// Find the matching section
	for _, section := range sections {
		if section.ID == firstSegment {
			if remainingPath == "" {
				// We've found the target directory
				if len(section.Children) > 0 {
					return section.Children[0].Href
				}
				return ""
			}
			// Continue searching in children
			return ns.findFirstItemInPath(section.Children, remainingPath, baseURL)
		}
	}

	return ""
}
