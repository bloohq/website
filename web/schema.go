package web

import (
	"encoding/json"
	"html/template"
	"log"
	"strings"
	"time"
)

// SchemaService handles structured data generation for SEO
type SchemaService struct {
	metadata *Metadata
	siteURL  string
}

// NewSchemaService creates a new schema service instance
func NewSchemaService(metadata *Metadata, siteURL string) *SchemaService {
	return &SchemaService{
		metadata: metadata,
		siteURL:  strings.TrimSuffix(siteURL, "/"),
	}
}

// GenerateSchema generates appropriate schema based on page type
func (s *SchemaService) GenerateSchema(pageType, path string, frontmatter *Frontmatter) template.JS {
	var schemaData []map[string]interface{}
	
	// Always include organization schema
	schemaData = append(schemaData, s.generateOrganizationSchema())
	
	// Add page-specific schemas
	switch pageType {
	case "platform":
		schemaData = append(schemaData, s.generateSoftwareApplicationSchema())
	case "pricing":
		schemaData = append(schemaData, s.generateProductSchema())
	case "insights":
		if frontmatter != nil {
			schemaData = append(schemaData, s.generateArticleSchema(frontmatter, path))
		}
	case "faq":
		schemaData = append(schemaData, s.generateFAQSchema())
	}
	
	// Convert to JSON
	jsonData, err := json.Marshal(schemaData)
	if err != nil {
		log.Printf("Error marshaling schema data: %v", err)
		return template.JS("[]")
	}
	
	return template.JS(jsonData)
}

// GetPageType determines the page type for schema generation
func (s *SchemaService) GetPageType(path string) string {
	switch {
	case strings.HasPrefix(path, "/platform"):
		return "platform"
	case path == "/pricing":
		return "pricing"
	case strings.HasPrefix(path, "/insights/"):
		return "insights"
	case strings.Contains(path, "/faq") || path == "/resources/faq":
		return "faq"
	default:
		return "general"
	}
}

// generateOrganizationSchema creates organization structured data
func (s *SchemaService) generateOrganizationSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"@context": "https://schema.org",
		"@type": "Organization",
		"name": "Blue",
		"url": s.siteURL,
		"logo": s.siteURL + "/logo/blue-logo.png",
		"description": "Blue is a powerful platform to create, manage, and scale processes for modern teams.",
		"sameAs": []string{
			"https://twitter.com/bluehq",
			"https://www.linkedin.com/company/blue-teamwork",
			"https://www.youtube.com/@HeyBlueTeam",
		},
		"contactPoint": map[string]interface{}{
			"@type": "ContactPoint",
			"contactType": "customer support",
			"email": "support@blue.cc",
		},
	}
	
	if s.metadata != nil && s.metadata.Site.Description != "" {
		schema["description"] = s.metadata.Site.Description
	}
	
	return schema
}

// generateSoftwareApplicationSchema creates software application structured data
func (s *SchemaService) generateSoftwareApplicationSchema() map[string]interface{} {
	return map[string]interface{}{
		"@context": "https://schema.org",
		"@type": "SoftwareApplication",
		"name": "Blue",
		"applicationCategory": "BusinessApplication",
		"operatingSystem": "Web, iOS, Android",
		"description": "Process management platform for modern teams",
		"screenshot": s.siteURL + "/product/dashboard-screenshot.png",
		"featureList": []string{
			"Process Management",
			"Team Collaboration", 
			"Workflow Automation",
			"Custom Fields",
			"API Access",
			"Real-time Updates",
			"File Attachments",
			"Custom Permissions",
		},
		"offers": map[string]interface{}{
			"@type": "Offer",
			"price": "7.00",
			"priceCurrency": "USD",
			"priceSpecification": map[string]interface{}{
				"@type": "UnitPriceSpecification",
				"price": "7.00",
				"priceCurrency": "USD",
				"unitText": "per user per month",
			},
		},
		"aggregateRating": map[string]interface{}{
			"@type": "AggregateRating",
			"ratingValue": "4.8",
			"reviewCount": "156",
			"bestRating": "5",
			"worstRating": "1",
		},
	}
}

// generateArticleSchema creates article structured data
func (s *SchemaService) generateArticleSchema(frontmatter *Frontmatter, path string) map[string]interface{} {
	schema := map[string]interface{}{
		"@context": "https://schema.org",
		"@type": "Article",
		"headline": frontmatter.Title,
		"description": frontmatter.Description,
		"url": s.siteURL + path,
		"author": map[string]interface{}{
			"@type": "Person",
			"name": "Blue Team",
		},
		"publisher": map[string]interface{}{
			"@type": "Organization",
			"name": "Blue",
			"logo": map[string]interface{}{
				"@type": "ImageObject",
				"url": s.siteURL + "/logo/blue-logo.png",
			},
		},
	}
	
	// Add dates if available
	if frontmatter.Date != "" {
		// Parse date and format as ISO 8601
		if t, err := time.Parse("2006-01-02", frontmatter.Date); err == nil {
			isoDate := t.Format(time.RFC3339)
			schema["datePublished"] = isoDate
			schema["dateModified"] = isoDate
		}
	}
	
	// Add article image if available
	if frontmatter.Image != "" {
		schema["image"] = s.siteURL + frontmatter.Image
	} else {
		schema["image"] = s.siteURL + "/og.png"
	}
	
	// Add category as articleSection
	if frontmatter.Category != "" {
		schema["articleSection"] = frontmatter.Category
	}
	
	// Add tags as keywords
	if len(frontmatter.Tags) > 0 {
		schema["keywords"] = strings.Join(frontmatter.Tags, ", ")
	}
	
	return schema
}

// generateProductSchema creates product/offer structured data for pricing
func (s *SchemaService) generateProductSchema() map[string]interface{} {
	return map[string]interface{}{
		"@context": "https://schema.org",
		"@type": "Product",
		"name": "Blue Core Subscription",
		"description": "Everything you need to create, manage, and scale your processes",
		"brand": map[string]interface{}{
			"@type": "Brand",
			"name": "Blue",
		},
		"offers": map[string]interface{}{
			"@type": "Offer",
			"url": s.siteURL + "/pricing",
			"priceCurrency": "USD",
			"price": "7.00",
			"priceValidUntil": "2025-12-31",
			"availability": "https://schema.org/InStock",
			"priceSpecification": map[string]interface{}{
				"@type": "UnitPriceSpecification",
				"price": "7.00",
				"priceCurrency": "USD",
				"unitText": "per user per month",
				"billingIncrement": 1,
				"billingDuration": "P1M",
			},
		},
		"aggregateRating": map[string]interface{}{
			"@type": "AggregateRating",
			"ratingValue": "4.8",
			"reviewCount": "156",
		},
	}
}

// generateFAQSchema creates FAQ structured data
func (s *SchemaService) generateFAQSchema() map[string]interface{} {
	// Selected FAQs from different categories for better SEO coverage
	faqs := []map[string]interface{}{
		// General
		{
			"@type": "Question",
			"name": "What is Blue?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Blue is a general purpose process management platform that can be used to manage a variety of processes across organizations. It helps staff to understand what they have to do, managers to have oversight, and for everyone to stay aligned.",
			},
		},
		{
			"@type": "Question",
			"name": "How is Blue different from Asana, Monday.com, or Trello?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "First of all, Blue is way cheaper. It's also not venture-funded, so our customers are our main priority. It's completely founder-led and we focus on a level of simplicity that they perhaps don't.",
			},
		},
		// Pricing
		{
			"@type": "Question",
			"name": "How much does Blue cost?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Blue has one main plan with monthly and yearly options: Monthly plan: $7/month/user. Yearly plan: $70/year/user (so you get two months free)",
			},
		},
		{
			"@type": "Question",
			"name": "How does the Blue free trial work?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "The free trial provides full access to all Blue features for 7 days. No credit card is required to start. If you wish to continue, you can upgrade to our paid plan at the end of the free trial and you will also have an option to extend your trial for another 7 days.",
			},
		},
		// Features
		{
			"@type": "Question",
			"name": "What integrations does Blue support?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Blue has 5,000+ integrations via Make.com, Zapier, Pabbly Connect, Albato, and Boost.Space. We also have a full API for custom integrations and granular Webhooks.",
			},
		},
		{
			"@type": "Question",
			"name": "Does Blue have Mobile Apps?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Yes, native iOS and Android apps are available to access Blue on the go. Core functionality like tasks, comments, and calendars sync across devices.",
			},
		},
		// Security
		{
			"@type": "Question",
			"name": "Is Blue GDPR compliant?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Yes, Blue is fully GDPR-compliant. For details, see our GDPR policy at https://blue.cc/legal/gdpr",
			},
		},
		{
			"@type": "Question",
			"name": "How is data encrypted in Blue?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Blue uses bank-level encryption for both data at rest and in transit. Data at rest is encrypted with 256-bit AES encryption. Data in transit is protected with TLS 1.2+ (or higher) protocols. This ensures your information is secure at all times.",
			},
		},
		// Getting Started
		{
			"@type": "Question",
			"name": "Can I import my existing data into Blue?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "Yes, Blue supports bulk data import from CSV files. This allows you to migrate up to 200,000 records simultaneously and is perfect for migration from manual Excel workflows or other work management systems.",
			},
		},
		// Support
		{
			"@type": "Question",
			"name": "What support options are available?",
			"acceptedAnswer": map[string]interface{}{
				"@type": "Answer",
				"text": "We offer email support at support@blue.cc. We also offer paid support packages for large-scale imports, full-team onboarding, priority chat support, and regular training.",
			},
		},
	}
	
	return map[string]interface{}{
		"@context": "https://schema.org",
		"@type": "FAQPage",
		"mainEntity": faqs,
	}
}