package web

import (
	"regexp"
	"strings"
	"golang.org/x/net/html"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// TOCEntry represents a table of contents entry
type TOCEntry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Level int    `json:"level"`
}

// ExtractHTMLTOC extracts table of contents from HTML content
// Looks for H2 elements with ID attributes
func ExtractHTMLTOC(content string) ([]TOCEntry, error) {
	var toc []TOCEntry
	
	// Wrap content in a proper HTML structure for parsing
	htmlContent := "<html><body>" + content + "</body></html>"
	
	// Parse HTML
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return toc, err
	}
	
	// Walk the HTML tree to find sections with IDs
	var walker func(*html.Node)
	walker = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "section" {
			// Extract ID attribute from section
			var id string
			for _, attr := range n.Attr {
				if attr.Key == "id" {
					id = attr.Val
					break
				}
			}
			
			// If section has ID, create TOC entry from ID
			if id != "" {
				// Generate title from ID (e.g., "key-features" → "Key Features")
				title := generateTitleFromID(id)
				toc = append(toc, TOCEntry{
					ID:    id,
					Title: title,
					Level: 2,
				})
			}
		}
		
		// Recursively walk children
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			walker(child)
		}
	}
	
	walker(doc)
	return toc, nil
}

// ExtractMarkdownTOC extracts table of contents from markdown AST
// Uses the goldmark parser to find heading nodes
func ExtractMarkdownTOC(renderer goldmark.Markdown, source []byte) ([]TOCEntry, error) {
	var toc []TOCEntry
	
	// Parse markdown to AST
	reader := text.NewReader(source)
	doc := renderer.Parser().Parse(reader)
	
	// Walk AST to find heading nodes
	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		
		// Only process H2 headings (level 2)
		if heading, ok := n.(*ast.Heading); ok && heading.Level == 2 {
			// Extract heading text
			var title strings.Builder
			for child := heading.FirstChild(); child != nil; child = child.NextSibling() {
				if text, ok := child.(*ast.Text); ok {
					title.Write(text.Segment.Value(source))
				}
			}
			
			// Generate ID from title (similar to goldmark's auto-heading ID)
			id := generateHeadingID(title.String())
			
			if id != "" && title.String() != "" {
				toc = append(toc, TOCEntry{
					ID:    id,
					Title: formatTitle(title.String()),
					Level: 2,
				})
			}
		}
		
		return ast.WalkContinue, nil
	})
	
	return toc, nil
}


// generateHeadingID generates a heading ID from text (similar to goldmark's algorithm)
func generateHeadingID(text string) string {
	// Convert to lowercase
	id := strings.ToLower(text)
	
	// Replace spaces with hyphens
	id = strings.ReplaceAll(id, " ", "-")
	
	// Remove non-alphanumeric characters except hyphens
	reg := regexp.MustCompile(`[^a-z0-9\-]`)
	id = reg.ReplaceAllString(id, "")
	
	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile(`-+`)
	id = reg.ReplaceAllString(id, "-")
	
	// Trim leading/trailing hyphens
	id = strings.Trim(id, "-")
	
	return id
}

// generateTitleFromID converts section IDs to readable titles
// e.g., "key-features" → "Key Features", "dns-settings" → "DNS Settings"
func generateTitleFromID(id string) string {
	// Replace hyphens with spaces
	title := strings.ReplaceAll(id, "-", " ")
	title = strings.ReplaceAll(title, "_", " ")
	
	// Apply smart title casing
	return formatTitle(title)
}

// formatTitle formats titles for display
// Converts "pricing" to "Pricing", preserves acronyms like "DNS Settings"
func formatTitle(title string) string {
	title = strings.TrimSpace(title)
	if title == "" {
		return title
	}
	
	// Smart title case: preserve acronyms (all caps words)
	words := strings.Fields(title)
	for i, word := range words {
		if len(word) > 0 {
			// Check if word is all uppercase (likely an acronym)
			isAcronym := true
			for _, char := range word {
				if char >= 'a' && char <= 'z' {
					isAcronym = false
					break
				}
			}
			
			// Preserve acronyms, otherwise apply title case
			if isAcronym && len(word) > 1 {
				words[i] = word // Keep as-is for acronyms
			} else {
				words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			}
		}
	}
	
	return strings.Join(words, " ")
}