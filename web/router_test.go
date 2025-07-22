package web

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRouterServeHTTP tests the main routing functionality
func TestRouterServeHTTP(t *testing.T) {
	// Create test directories
	testDir := t.TempDir()
	pagesDir := filepath.Join(testDir, "pages")
	componentsDir := filepath.Join(testDir, "components")
	layoutsDir := filepath.Join(testDir, "layouts")
	contentDir := filepath.Join(testDir, "content")
	
	// Create directories
	for _, dir := range []string{pagesDir, componentsDir, layoutsDir, contentDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}
	
	// Create test files
	createTestFile(t, filepath.Join(pagesDir, "index.html"), "<h1>Home Page</h1>")
	createTestFile(t, filepath.Join(pagesDir, "about.html"), "<h1>About Page</h1>")
	createTestFile(t, filepath.Join(pagesDir, "404.html"), "<h1>404 - Page Not Found</h1>")
	createTestFile(t, filepath.Join(layoutsDir, "main.html"), `{{define "main.html"}}<!DOCTYPE html><html><body>{{.Content}}</body></html>{{end}}`)
	
	// Create router with test directories
	router := &Router{
		pagesDir:      pagesDir,
		layoutsDir:    layoutsDir,
		componentsDir: componentsDir,
		contentDir:    contentDir,
		tocExcludedPaths: []string{"/changelog"},
	}
	
	// Initialize minimal services to prevent nil pointer errors
	router.seoService = NewSEOService()
	router.markdownService = NewMarkdownService()
	router.contentService = NewContentService(contentDir)
	router.navigationService = NewNavigationService(router.seoService)
	router.htmlService = NewHTMLService(pagesDir, layoutsDir, componentsDir, router.markdownService)
	router.schemaService = NewSchemaService(nil, "https://example.com")
	
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Homepage",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "Home Page",
		},
		{
			name:           "About page",
			path:           "/about",
			expectedStatus: http.StatusOK,
			expectedBody:   "About Page",
		},
		{
			name:           "Non-existent page",
			path:           "/non-existent",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 - Page Not Found",
		},
		{
			name:           "Health endpoint",
			path:           "/health",
			expectedStatus: http.StatusOK,
			expectedBody:   "\"status\":\"ok\"",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()
			
			router.ServeHTTP(w, req)
			
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			
			if !strings.Contains(string(body), tt.expectedBody) {
				t.Errorf("Expected body to contain %q, got %q", tt.expectedBody, string(body))
			}
		})
	}
}

// TestRouterRedirects tests URL redirects
func TestRouterRedirects(t *testing.T) {
	// Create test directories
	testDir := t.TempDir()
	router := createTestRouter(testDir)
	
	// Set up redirects in SEO service
	router.seoService.redirects = &Redirects{
		Redirects: map[string]string{
			"/old-page": "/new-page",
		},
		Rules: RedirectRules{
			StatusCode: http.StatusMovedPermanently,
		},
	}
	
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedLocation string
	}{
		{
			name:           "HTML extension redirect",
			path:           "/about.html",
			expectedStatus: http.StatusMovedPermanently,
			expectedLocation: "/about",
		},
		{
			name:           "Custom redirect",
			path:           "/old-page",
			expectedStatus: http.StatusMovedPermanently,
			expectedLocation: "/new-page",
		},
		{
			name:           "API docs redirect",
			path:           "/api-docs/test",
			expectedStatus: http.StatusMovedPermanently,
			expectedLocation: "/api/test",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()
			
			router.ServeHTTP(w, req)
			
			resp := w.Result()
			
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			
			location := resp.Header.Get("Location")
			if location != tt.expectedLocation {
				t.Errorf("Expected location %q, got %q", tt.expectedLocation, location)
			}
		})
	}
}

// TestLoadComponentTemplates tests component template loading
func TestLoadComponentTemplates(t *testing.T) {
	testDir := t.TempDir()
	componentsDir := filepath.Join(testDir, "components")
	os.MkdirAll(componentsDir, 0755)
	
	// Create test component files
	createTestFile(t, filepath.Join(componentsDir, "header.html"), "{{define \"header\"}}Header{{end}}")
	createTestFile(t, filepath.Join(componentsDir, "footer.html"), "{{define \"footer\"}}Footer{{end}}")
	
	router := &Router{
		componentsDir: componentsDir,
	}
	
	files, err := router.loadComponentTemplates()
	if err != nil {
		t.Fatalf("Failed to load component templates: %v", err)
	}
	
	if len(files) != 2 {
		t.Errorf("Expected 2 component files, got %d", len(files))
	}
	
	// Verify files are HTML files
	for _, file := range files {
		if !strings.HasSuffix(file, ".html") {
			t.Errorf("Expected HTML file, got %s", file)
		}
	}
}

// TestRouterCaching tests that cached content is served correctly
func TestRouterCaching(t *testing.T) {
	testDir := t.TempDir()
	router := createTestRouter(testDir)
	
	// Create a test page file first
	pagesDir := filepath.Join(testDir, "pages")
	createTestFile(t, filepath.Join(pagesDir, "cached.html"), "<h1>Original Content</h1>")
	
	// Pre-cache the page with different content
	cachedContent := &CachedContent{
		HTML: "<!DOCTYPE html><html><body><h1>Cached Content</h1></body></html>",
	}
	router.htmlService.cache.Set("/cached", cachedContent)
	
	req := httptest.NewRequest("GET", "/cached", nil)
	w := httptest.NewRecorder()
	
	router.ServeHTTP(w, req)
	
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	
	// Should serve cached content, not original file content
	if !strings.Contains(string(body), "Cached Content") {
		t.Errorf("Expected cached content, got %q", string(body))
	}
	
	if strings.Contains(string(body), "Original Content") {
		t.Errorf("Should not contain original content when serving from cache")
	}
}

// Helper functions

func createTestRouter(testDir string) *Router {
	pagesDir := filepath.Join(testDir, "pages")
	componentsDir := filepath.Join(testDir, "components")
	layoutsDir := filepath.Join(testDir, "layouts")
	contentDir := filepath.Join(testDir, "content")
	
	// Create directories
	for _, dir := range []string{pagesDir, componentsDir, layoutsDir, contentDir} {
		os.MkdirAll(dir, 0755)
	}
	
	// Create basic layout
	createTestFile(nil, filepath.Join(layoutsDir, "main.html"), 
		`{{define "main.html"}}<!DOCTYPE html><html><body>{{.Content}}</body></html>{{end}}`)
	
	router := &Router{
		pagesDir:      pagesDir,
		layoutsDir:    layoutsDir,
		componentsDir: componentsDir,
		contentDir:    contentDir,
		tocExcludedPaths: []string{"/changelog"},
	}
	
	// Initialize services
	router.seoService = NewSEOService()
	router.markdownService = NewMarkdownService()
	router.contentService = NewContentService(contentDir)
	router.navigationService = NewNavigationService(router.seoService)
	router.htmlService = NewHTMLService(pagesDir, layoutsDir, componentsDir, router.markdownService)
	router.schemaService = NewSchemaService(nil, "https://example.com")
	
	return router
}

func createTestFile(t *testing.T, path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		if t != nil {
			t.Fatalf("Failed to create test file %s: %v", path, err)
		}
		panic(err)
	}
}