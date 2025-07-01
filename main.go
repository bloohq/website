package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"blue-website/web"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	// Validate required environment variables
	requiredEnvVars := []string{
		"CLOUDFLARE_ACCOUNT_ID",
		"CLOUDFLARE_DATABASE_ID",
		"CLOUDFLARE_API_KEY",
	}
	
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Printf("Warning: %s environment variable not set, status monitoring disabled", envVar)
		}
	}

	// Generate search index at startup
	fmt.Println("🔍 Generating search index...")
	if err := web.GenerateSearchIndex(); err != nil {
		log.Printf("⚠️  Warning: Failed to generate search index: %v", err)
	} else {
		fmt.Println("✅ Search index generated successfully")
	}

	// Serve static files from public directory
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// File-based routing handler
	router := web.NewRouter("pages")
	
	// Initialize status monitoring if environment variables are set
	if os.Getenv("CLOUDFLARE_API_KEY") != "" {
		fmt.Println("🏥 Initializing status monitoring...")
		
		// Create D1 client
		d1Client := web.NewD1Client()
		
		// Create health checker
		healthChecker := web.NewHealthChecker(d1Client)
		
		// Initialize database and load historical data
		if err := healthChecker.Initialize(); err != nil {
			log.Printf("⚠️  Warning: Failed to initialize status monitoring: %v", err)
		} else {
			fmt.Println("✅ Status monitoring initialized")
			
			// Set the health checker in the router
			router.SetStatusChecker(healthChecker)
			
			// Start background health checks
			go func() {
				ticker := time.NewTicker(5 * time.Minute)
				defer ticker.Stop()
				
				// Run initial check only if needed
				log.Println("🔍 Checking if initial health checks are needed...")
				healthChecker.CheckAllServicesIfNeeded()
				
				// Run periodic checks
				for range ticker.C {
					log.Println("⏰ Running scheduled health checks...")
					healthChecker.CheckAllServices()
				}
			}()
		}
	} else {
		log.Println("⚠️  Status monitoring disabled (missing environment variables)")
	}
	
	http.Handle("/", router)

	// Get port from environment variable, default to 8080 for local development
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("🚀 Blue Website server starting on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

 