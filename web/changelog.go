package web

import (
	"encoding/json"
	"log"
	"os"
)

// ChangelogEntry represents a single changelog item
type ChangelogEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ChangelogMonth represents a month's worth of changelog entries
type ChangelogMonth struct {
	Month   string            `json:"month"`
	Year    string            `json:"year"`
	Changes []ChangelogEntry  `json:"changes"`
}

// ChangelogService handles changelog data loading and management
type ChangelogService struct {
	changelog []ChangelogMonth
}

// NewChangelogService creates a new changelog service instance
func NewChangelogService() *ChangelogService {
	return &ChangelogService{}
}

// LoadChangelog loads changelog data from JSON file
func (c *ChangelogService) LoadChangelog() error {
	data, err := os.ReadFile("data/changelog.json")
	if err != nil {
		log.Printf("Warning: Could not load changelog data: %v", err)
		return err
	}
	
	if err := json.Unmarshal(data, &c.changelog); err != nil {
		log.Printf("Error parsing changelog JSON: %v", err)
		return err
	}
	
	log.Printf("Successfully loaded %d changelog entries", len(c.changelog))
	return nil
}

// GetChangelog returns the loaded changelog data
func (c *ChangelogService) GetChangelog() []ChangelogMonth {
	return c.changelog
}

// GetChangelogForYear returns changelog entries for a specific year
func (c *ChangelogService) GetChangelogForYear(year string) []ChangelogMonth {
	var yearEntries []ChangelogMonth
	for _, entry := range c.changelog {
		if entry.Year == year {
			yearEntries = append(yearEntries, entry)
		}
	}
	return yearEntries
}

// GetAvailableYears returns all years that have changelog entries
func (c *ChangelogService) GetAvailableYears() []string {
	yearMap := make(map[string]bool)
	var years []string
	
	for _, entry := range c.changelog {
		if !yearMap[entry.Year] {
			years = append(years, entry.Year)
			yearMap[entry.Year] = true
		}
	}
	
	return years
}