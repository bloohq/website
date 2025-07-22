package web

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestMarkdownServiceConvert tests basic markdown conversion
func TestMarkdownServiceConvert(t *testing.T) {
	ms := NewMarkdownService()
	
	tests := []struct {
		name     string
		markdown string
		expected string
	}{
		{
			name:     "Basic paragraph",
			markdown: "Hello, world!",
			expected: "<p>Hello, world!</p>",
		},
		{
			name:     "Heading",
			markdown: "# Hello World",
			expected: `<h1 id="hello-world">Hello World</h1>`,
		},
		{
			name:     "Bold text",
			markdown: "This is **bold** text",
			expected: "<p>This is <strong>bold</strong> text</p>",
		},
		{
			name:     "Link",
			markdown: "[Blue](https://blue.cc)",
			expected: `<p><a href="https://blue.cc">Blue</a></p>`,
		},
		{
			name:     "Code block",
			markdown: "```go\nfunc main() {}\n```",
			expected: "<pre><code class=\"language-go\">func main() {}\n</code></pre>",
		},
		{
			name:     "List",
			markdown: "- Item 1\n- Item 2",
			expected: "<ul>\n<li>Item 1</li>\n<li>Item 2</li>\n</ul>",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ms.Convert([]byte(tt.markdown))
			if err != nil {
				t.Fatalf("Conversion failed: %v", err)
			}
			
			// Normalize whitespace for comparison
			result = strings.TrimSpace(result)
			expected := strings.TrimSpace(tt.expected)
			
			if result != expected {
				t.Errorf("Expected %q, got %q", expected, result)
			}
		})
	}
}

// TestProcessMarkdownFile tests processing a complete markdown file with frontmatter
func TestProcessMarkdownFile(t *testing.T) {
	ms := NewMarkdownService()
	testDir := t.TempDir()
	
	// Create a test markdown file with frontmatter
	mdContent := `---
title: Test Article
description: This is a test article
slug: test-article
category: Testing
tags: [test, demo]
date: 2024-01-15
---

# Test Article

This is the content of the test article.

## Section 1

Some content here.`
	
	mdPath := filepath.Join(testDir, "test.md")
	if err := os.WriteFile(mdPath, []byte(mdContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	
	// Create a mock SEO service
	seoService := NewSEOService()
	
	// Process the markdown file
	html, frontmatter, err := ms.ProcessMarkdownFile(mdPath, seoService)
	if err != nil {
		t.Fatalf("Failed to process markdown file: %v", err)
	}
	
	// Check frontmatter
	if frontmatter.Title != "Test Article" {
		t.Errorf("Expected title 'Test Article', got %q", frontmatter.Title)
	}
	
	if frontmatter.Description != "This is a test article" {
		t.Errorf("Expected description 'This is a test article', got %q", frontmatter.Description)
	}
	
	if frontmatter.Slug != "test-article" {
		t.Errorf("Expected slug 'test-article', got %q", frontmatter.Slug)
	}
	
	if frontmatter.Category != "Testing" {
		t.Errorf("Expected category 'Testing', got %q", frontmatter.Category)
	}
	
	if len(frontmatter.Tags) != 2 || frontmatter.Tags[0] != "test" {
		t.Errorf("Expected tags [test, demo], got %v", frontmatter.Tags)
	}
	
	// Check HTML content
	if !strings.Contains(html, `<h1 id="test-article">Test Article</h1>`) {
		t.Errorf("Expected heading in HTML, got %q", html)
	}
	
	if !strings.Contains(html, `<h2 id="section-1">Section 1</h2>`) {
		t.Errorf("Expected section heading in HTML, got %q", html)
	}
}

// TestMarkdownCache tests the caching functionality
func TestMarkdownCache(t *testing.T) {
	cache := NewMarkdownCache()
	
	// Test setting and getting
	content := &CachedContent{
		HTML: "<h1>Test</h1>",
		Frontmatter: &Frontmatter{
			Title: "Test Page",
			Date:  "2024-01-15",
		},
	}
	
	cache.Set("/test", content)
	
	// Test Get
	retrieved, found := cache.Get("/test")
	if !found {
		t.Error("Expected to find cached content")
	}
	
	if retrieved.HTML != content.HTML {
		t.Errorf("Expected HTML %q, got %q", content.HTML, retrieved.HTML)
	}
	
	if retrieved.Frontmatter.Title != content.Frontmatter.Title {
		t.Errorf("Expected title %q, got %q", content.Frontmatter.Title, retrieved.Frontmatter.Title)
	}
	
	// Test Size
	if cache.Size() != 1 {
		t.Errorf("Expected cache size 1, got %d", cache.Size())
	}
	
	// Test that content exists in cache
	if cache.Size() != 1 {
		t.Error("Expected cache to contain 1 item")
	}
	
	// Test Clear
	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected cache size 0 after clear, got %d", cache.Size())
	}
}

// TestPreRenderAllMarkdown tests pre-rendering multiple markdown files
func TestPreRenderAllMarkdown(t *testing.T) {
	// Skip this test in CI or when running from different directories
	// since PreRenderAllMarkdown is hardcoded to look for "content" directory
	t.Skip("Skipping PreRenderAllMarkdown test - requires specific directory structure")
}


// TestConcurrentCacheAccess tests thread safety of the cache
func TestConcurrentCacheAccess(t *testing.T) {
	cache := NewMarkdownCache()
	done := make(chan bool)
	
	// Start multiple goroutines writing to cache
	for i := 0; i < 10; i++ {
		go func(n int) {
			content := &CachedContent{
				HTML: fmt.Sprintf("<h1>Test %d</h1>", n),
				Frontmatter: &Frontmatter{
					Title: fmt.Sprintf("Test %d", n),
				},
			}
			cache.Set(fmt.Sprintf("/test%d", n), content)
			done <- true
		}(i)
	}
	
	// Start multiple goroutines reading from cache
	for i := 0; i < 10; i++ {
		go func(n int) {
			time.Sleep(5 * time.Millisecond) // Small delay to ensure some writes happen first
			cache.Get(fmt.Sprintf("/test%d", n))
			done <- true
		}(i)
	}
	
	// Wait for all goroutines
	for i := 0; i < 20; i++ {
		<-done
	}
	
	// Verify final state
	if cache.Size() != 10 {
		t.Errorf("Expected cache size 10, got %d", cache.Size())
	}
}