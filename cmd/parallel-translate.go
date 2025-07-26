package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// RateLimiter controls the rate of API requests
type RateLimiter struct {
	ticker   *time.Ticker
	tokens   chan struct{}
	stopChan chan struct{}
}

// NewRateLimiter creates a new rate limiter that allows 'requestsPerMinute' requests per minute
func NewRateLimiter(requestsPerMinute int) *RateLimiter {
	interval := time.Minute / time.Duration(requestsPerMinute)

	rl := &RateLimiter{
		ticker:   time.NewTicker(interval),
		tokens:   make(chan struct{}, 5), // Buffer of 5 tokens for burst capability
		stopChan: make(chan struct{}),
	}

	// Pre-fill with initial tokens
	for i := 0; i < 5; i++ {
		select {
		case rl.tokens <- struct{}{}:
		default:
		}
	}

	// Start the token generator
	go rl.generateTokens()

	return rl
}

// Wait blocks until a token is available
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

// Stop stops the rate limiter
func (rl *RateLimiter) Stop() {
	close(rl.stopChan)
	rl.ticker.Stop()
}

// generateTokens adds tokens to the bucket at the specified rate
func (rl *RateLimiter) generateTokens() {
	for {
		select {
		case <-rl.ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
				// Token bucket is full, skip
			}
		case <-rl.stopChan:
			return
		}
	}
}

// TranslationJob represents a single file translation task
type TranslationJob struct {
	SourceFile string
	TargetLang string
	SourceDir  string
	TargetDir  string
	Content    string
	Attempt    int
}

// TranslationResult represents the result of a translation job
type TranslationResult struct {
	Job        TranslationJob
	Success    bool
	Error      error
	Duration   time.Duration
	OutputPath string
}

// ProgressTracker tracks translation progress across all languages
type ProgressTracker struct {
	mu                sync.RWMutex
	totalJobs         int64
	completedJobs     int64
	activeWorkers     int64
	languageProgress  map[string]*LanguageProgress
	recentCompletions []CompletionInfo
	errors            []ErrorInfo
	startTime         time.Time
	jobsPerSecond     float64
}

// LanguageProgress tracks progress for a specific language
type LanguageProgress struct {
	Language   string
	Emoji      string
	Total      int
	Completed  int
	Percentage float64
}

// CompletionInfo represents a recently completed translation
type CompletionInfo struct {
	FileName string
	Language string
	Duration time.Duration
	Time     time.Time
}

// ErrorInfo represents a translation error
type ErrorInfo struct {
	Job     TranslationJob
	Error   error
	Time    time.Time
	Retries int
}

var languageEmojis = map[string]string{
	"de":    "🇩🇪",
	"es":    "🇪🇸",
	"fr":    "🇫🇷",
	"it":    "🇮🇹",
	"ja":    "🇯🇵",
	"ko":    "🇰🇷",
	"zh":    "🇨🇳",
	"zh-TW": "🇹🇼",
	"ru":    "🇷🇺",
	"pl":    "🇵🇱",
	"pt":    "🇵🇹",
	"nl":    "🇳🇱",
	"sv":    "🇸🇪",
	"id":    "🇮🇩",
	"km":    "🇰🇭",
}

var languageNames = map[string]string{
	"de":    "German",
	"es":    "Spanish",
	"fr":    "French",
	"it":    "Italian",
	"ja":    "Japanese",
	"ko":    "Korean",
	"zh":    "Chinese",
	"zh-TW": "Chinese (Traditional)",
	"ru":    "Russian",
	"pl":    "Polish",
	"pt":    "Portuguese",
	"nl":    "Dutch",
	"sv":    "Swedish",
	"id":    "Indonesian",
	"km":    "Khmer",
}

func main() {
	fmt.Println("🚀 Blue Insights Translation Engine")
	fmt.Println("===================================")

	// Scan source directory
	sourceDir := "content/en/insights"
	fmt.Printf("\n📁 Scanning %s...\n", sourceDir)

	sourceFiles, err := scanMarkdownFiles(sourceDir)
	if err != nil {
		fmt.Printf("❌ Error scanning source directory: %v\n", err)
		return
	}

	if len(sourceFiles) == 0 {
		fmt.Printf("❌ No markdown files found in %s\n", sourceDir)
		return
	}

	fmt.Printf("✅ Found %d markdown files\n", len(sourceFiles))

	// Get target languages (exclude English)
	targetLanguages := getTargetLanguages()
	if len(targetLanguages) == 0 {
		fmt.Println("❌ No target languages found")
		return
	}

	fmt.Printf("\n🌐 Target languages: %s (%d languages)\n",
		strings.Join(targetLanguages, ", "), len(targetLanguages))

	totalJobs := len(sourceFiles) * len(targetLanguages)
	fmt.Printf("📊 Total translation jobs: %d (%d files × %d languages)\n",
		totalJobs, len(sourceFiles), len(targetLanguages))

	// Initialize progress tracker
	tracker := &ProgressTracker{
		totalJobs:         int64(totalJobs),
		languageProgress:  make(map[string]*LanguageProgress),
		recentCompletions: make([]CompletionInfo, 0),
		errors:            make([]ErrorInfo, 0),
		startTime:         time.Now(),
	}

	// Initialize language progress tracking
	for _, lang := range targetLanguages {
		emoji := languageEmojis[lang]
		if emoji == "" {
			emoji = "🌐"
		}
		tracker.languageProgress[lang] = &LanguageProgress{
			Language:   lang,
			Emoji:      emoji,
			Total:      len(sourceFiles),
			Completed:  0,
			Percentage: 0,
		}
	}

	// Create rate limiter for 60 requests per minute
	rateLimiter := NewRateLimiter(60)
	defer rateLimiter.Stop()

	// Create job queue and result channels
	const maxWorkers = 10 // Reduced from 100 to ensure rate limiting is effective
	jobQueue := make(chan TranslationJob, totalJobs)
	resultChan := make(chan TranslationResult, maxWorkers)

	fmt.Printf("\n🔧 Starting %d worker goroutines with rate limit of 60 requests/minute...\n", maxWorkers)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(i+1, jobQueue, resultChan, tracker, rateLimiter, &wg)
	}

	// Start progress display
	progressDone := make(chan bool)
	go displayProgress(tracker, progressDone)

	// Queue all jobs
	go func() {
		defer close(jobQueue)

		for _, sourceFile := range sourceFiles {
			// Read source content
			content, err := os.ReadFile(sourceFile)
			if err != nil {
				fmt.Printf("❌ Error reading %s: %v\n", sourceFile, err)
				continue
			}

			// Create jobs for all target languages
			for _, lang := range targetLanguages {
				targetDir := fmt.Sprintf("content/%s/insights", lang)

				job := TranslationJob{
					SourceFile: sourceFile,
					TargetLang: lang,
					SourceDir:  sourceDir,
					TargetDir:  targetDir,
					Content:    string(content),
					Attempt:    1,
				}

				jobQueue <- job
			}
		}
	}()

	fmt.Println("⚡ Workers ready - beginning translation\n")

	// Collect results
	go func() {
		defer close(resultChan)
		wg.Wait()
	}()

	// Process results
	successCount := 0
	errorCount := 0

	for result := range resultChan {
		if result.Success {
			successCount++
			tracker.recordCompletion(result)
		} else {
			errorCount++
			tracker.recordError(result)

			// Retry failed jobs (up to 3 attempts)
			if result.Job.Attempt < 3 {
				retryJob := result.Job
				retryJob.Attempt++
				go func() {
					time.Sleep(time.Duration(retryJob.Attempt) * time.Second) // Exponential backoff
					select {
					case jobQueue <- retryJob:
					default:
						// Job queue closed, skip retry
					}
				}()
			}
		}
	}

	// Stop progress display
	progressDone <- true

	// Final statistics
	tracker.displayFinalStats(successCount, errorCount)
}

func scanMarkdownFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			files = append(files, path)
		}

		return nil
	})

	sort.Strings(files)
	return files, err
}

func getTargetLanguages() []string {
	contentDir := "content"
	entries, err := os.ReadDir(contentDir)
	if err != nil {
		return []string{}
	}

	var languages []string
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "en" {
			languages = append(languages, entry.Name())
		}
	}

	sort.Strings(languages)
	return languages
}

func worker(id int, jobs <-chan TranslationJob, results chan<- TranslationResult, tracker *ProgressTracker, rateLimiter *RateLimiter, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		atomic.AddInt64(&tracker.activeWorkers, 1)

		result := processTranslationJob(job, rateLimiter)
		results <- result

		atomic.AddInt64(&tracker.activeWorkers, -1)
		atomic.AddInt64(&tracker.completedJobs, 1)
	}
}

func processTranslationJob(job TranslationJob, rateLimiter *RateLimiter) TranslationResult {
	startTime := time.Now()

	// Ensure target directory exists
	if err := os.MkdirAll(job.TargetDir, 0755); err != nil {
		return TranslationResult{
			Job:      job,
			Success:  false,
			Error:    fmt.Errorf("failed to create directory: %v", err),
			Duration: time.Since(startTime),
		}
	}

	// Build target file path
	sourceFileName := filepath.Base(job.SourceFile)
	targetPath := filepath.Join(job.TargetDir, sourceFileName)

	// Skip if file already exists (resume capability)
	if _, err := os.Stat(targetPath); err == nil {
		return TranslationResult{
			Job:        job,
			Success:    true,
			Duration:   time.Since(startTime),
			OutputPath: targetPath,
		}
	}

	// Wait for rate limiter before making API call
	rateLimiter.Wait()

	// Call Claude for translation
	translatedContent, err := callClaude(job.Content, job.TargetLang, sourceFileName)
	if err != nil {
		return TranslationResult{
			Job:      job,
			Success:  false,
			Error:    err,
			Duration: time.Since(startTime),
		}
	}

	// Write translated content
	if err := os.WriteFile(targetPath, []byte(translatedContent), 0644); err != nil {
		return TranslationResult{
			Job:      job,
			Success:  false,
			Error:    fmt.Errorf("failed to write file: %v", err),
			Duration: time.Since(startTime),
		}
	}

	return TranslationResult{
		Job:        job,
		Success:    true,
		Duration:   time.Since(startTime),
		OutputPath: targetPath,
	}
}

func callClaude(content, targetLang, fileName string) (string, error) {
	langName := languageNames[targetLang]
	if langName == "" {
		langName = targetLang
	}

	prompt := fmt.Sprintf(`You are translating marketing content for Blue, a B2B SaaS process management platform.

CRITICAL TRANSLATION RULES - NEVER VIOLATE THESE:

1. NEVER TRANSLATE THESE ELEMENTS:
   - "Blue" - This is our product name and must NEVER be translated
   - Email addresses - Keep support@blue.cc, sales@blue.cc, etc. unchanged
   - URLs - All links must remain as-is
   - Image filenames - Keep /product/dashboard.png unchanged
   - Code blocks - Any code snippets or technical content
   - Placeholders - Keep {name}, {{count}}, %%s, etc. intact
   - Technical terms - API, GraphQL, webhook, OAuth, etc. (unless there's an established translation)
   - Person names - Emanuele Faja, etc.
   - Category in the front-matter - "category: "

2. MARKDOWN PRESERVATION:
   - Keep all markdown formatting exactly as is (headers, links, lists, etc.)
   - Preserve frontmatter structure (YAML between --- markers)
   - Keep code blocks unchanged
   - Preserve image references
   - Keep internal links (only update language prefix if needed: /en/ → /%s/)
   - Preserve callouts structure: ::callout{type="info"} (translate content, keep structure)

3. TRANSLATION CONTEXT:
   - Blue is a B2B SaaS process management platform
   - Maintain professional, business-appropriate tone
   - Prioritize clarity over literal translation
   - Use consistent terminology throughout
   - Consider screen space limitations for UI text

4. LANGUAGE-SPECIFIC GUIDELINES:
   %s

5. QUALITY REQUIREMENTS:
   - All placeholders ({name}, {{count}}) must be preserved
   - Technical terms appropriately handled
   - Consistent terminology throughout
   - Appropriate tone for B2B SaaS
   - Special characters properly encoded
   - HTML tags preserved (e.g., <br>)

6. COMMON PITFALLS TO AVOID:
   - Don't translate company names mentioned in content
   - Don't localize technical paths (/api/v1/...)
   - Don't change date/number formats in placeholders
   - Don't translate acronyms unless widely established
   - Don't modify code examples or configuration snippets
   - Don't remove HTML tags like <br> in translations

TARGET LANGUAGE: %s
FILE: %s

INSTRUCTIONS: Return ONLY the translated markdown content maintaining exact formatting. No explanations, no metadata, no wrapping.

CONTENT TO TRANSLATE:

%s`, targetLang, getLanguageSpecificGuidelines(targetLang), langName, fileName, content)

	// Call claude command
	cmd := exec.Command("claude", "-p", prompt)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("claude command failed: %v", err)
	}

	return string(output), nil
}

func getLanguageSpecificGuidelines(langCode string) string {
	guidelines := map[string]string{
		"es":    "Spanish: Use formal \"usted\" for business context. Maintain Latin American neutral Spanish. Common terms: Task → Tarea, Project → Proyecto, Workflow → Flujo de trabajo, Dashboard → Panel de control.",
		"fr":    "French: Use formal \"vous\" form. Maintain international French. Common terms: Task → Tâche, Project → Projet, Workflow → Flux de travail, Dashboard → Tableau de bord.",
		"de":    "German: Use formal \"Sie\" form. Maintain standard Hochdeutsch. Consider compound words for clarity.",
		"pt":    "Portuguese: Use Brazilian Portuguese as default. Formal business tone. Avoid regional slang.",
		"it":    "Italian: Use formal business tone. Maintain standard Italian.",
		"ja":    "Japanese: Use formal business language (keigo). Maintain professional tone appropriate for B2B context.",
		"ko":    "Korean: Use formal business language (jondaetmal). Maintain professional tone appropriate for B2B context.",
		"zh":    "Chinese (Simplified): Use formal business language. Maintain professional tone appropriate for B2B context.",
		"zh-TW": "Chinese (Traditional): Use formal business language. Maintain professional tone appropriate for B2B context.",
		"ru":    "Russian: Use formal business language. Maintain professional tone appropriate for B2B context.",
		"pl":    "Polish: Use formal business language. Maintain professional tone appropriate for B2B context.",
		"nl":    "Dutch: Use formal business language. Maintain professional tone appropriate for B2B context.",
		"sv":    "Swedish: Use formal business language. Maintain professional tone appropriate for B2B context.",
		"id":    "Indonesian: Use formal business language (bahasa formal). Maintain professional tone appropriate for B2B context.",
		"km":    "Khmer: Use formal business language. Maintain professional tone appropriate for B2B context.",
	}

	if guideline, exists := guidelines[langCode]; exists {
		return guideline
	}
	return "Use formal business language. Maintain professional tone appropriate for B2B context."
}

func (t *ProgressTracker) recordCompletion(result TranslationResult) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Update language progress
	if langProgress, exists := t.languageProgress[result.Job.TargetLang]; exists {
		langProgress.Completed++
		langProgress.Percentage = float64(langProgress.Completed) / float64(langProgress.Total) * 100
	}

	// Add to recent completions (keep last 5)
	fileName := filepath.Base(result.Job.SourceFile)
	completion := CompletionInfo{
		FileName: fileName,
		Language: result.Job.TargetLang,
		Duration: result.Duration,
		Time:     time.Now(),
	}

	t.recentCompletions = append(t.recentCompletions, completion)
	if len(t.recentCompletions) > 5 {
		t.recentCompletions = t.recentCompletions[1:]
	}

	// Update jobs per second
	elapsed := time.Since(t.startTime).Seconds()
	if elapsed > 0 {
		t.jobsPerSecond = float64(t.completedJobs) / elapsed
	}
}

func (t *ProgressTracker) recordError(result TranslationResult) {
	t.mu.Lock()
	defer t.mu.Unlock()

	errorInfo := ErrorInfo{
		Job:     result.Job,
		Error:   result.Error,
		Time:    time.Now(),
		Retries: result.Job.Attempt,
	}

	t.errors = append(t.errors, errorInfo)

	// Keep only last 10 errors
	if len(t.errors) > 10 {
		t.errors = t.errors[1:]
	}
}

func displayProgress(tracker *ProgressTracker, done <-chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			tracker.displayCurrentProgress()
		}
	}
}

func (t *ProgressTracker) displayCurrentProgress() {
	t.mu.RLock()
	defer t.mu.RUnlock()

	// Clear screen and move cursor to top
	fmt.Print("\033[2J\033[H")

	fmt.Println("🚀 Blue Insights Translation Engine")
	fmt.Println("===================================")

	// Overall progress
	completed := atomic.LoadInt64(&t.completedJobs)
	activeWorkers := atomic.LoadInt64(&t.activeWorkers)
	remaining := t.totalJobs - completed

	percentage := float64(completed) / float64(t.totalJobs) * 100
	progressBar := generateProgressBar(int(percentage), 40)

	fmt.Printf("\nProgress: [%s] %.1f%% (%d/%d) ", progressBar, percentage, completed, t.totalJobs)
	if percentage >= 100 {
		fmt.Print("✅ Complete!")
	}
	fmt.Println()

	fmt.Printf("\nActive Workers: %d/10 | Queue: %d jobs remaining | Rate: %.1f jobs/sec (limited to 60/min)\n",
		activeWorkers, remaining, t.jobsPerSecond)

	// Recent completions
	if len(t.recentCompletions) > 0 {
		fmt.Println("\nRecent Completions:")
		for i := len(t.recentCompletions) - 1; i >= 0; i-- {
			comp := t.recentCompletions[i]
			langName := languageNames[comp.Language]
			if langName == "" {
				langName = comp.Language
			}
			fmt.Printf("✅ %s → %s (%s) - %.1fs\n",
				comp.FileName, langName, comp.Language, comp.Duration.Seconds())
		}
	}

	// Language progress
	fmt.Println("\nLanguage Progress:")

	// Sort languages by completion percentage (highest first)
	var sortedLangs []string
	for lang := range t.languageProgress {
		sortedLangs = append(sortedLangs, lang)
	}
	sort.Slice(sortedLangs, func(i, j int) bool {
		return t.languageProgress[sortedLangs[i]].Percentage > t.languageProgress[sortedLangs[j]].Percentage
	})

	for _, lang := range sortedLangs {
		progress := t.languageProgress[lang]
		progressBar := generateProgressBar(int(progress.Percentage), 10)
		status := ""
		if progress.Percentage >= 100 {
			status = " ✅"
		}

		langName := languageNames[lang]
		if langName == "" {
			langName = lang
		}

		fmt.Printf("%s %-20s [%s] %3.0f%% (%d/%d)%s\n",
			progress.Emoji, langName+":", progressBar, progress.Percentage,
			progress.Completed, progress.Total, status)
	}

	// Errors
	if len(t.errors) > 0 {
		fmt.Printf("\nErrors: %d (retrying...)\n", len(t.errors))
		for i := len(t.errors) - 1; i >= 0 && i >= len(t.errors)-3; i-- {
			err := t.errors[i]
			fileName := filepath.Base(err.Job.SourceFile)
			langName := languageNames[err.Job.TargetLang]
			if langName == "" {
				langName = err.Job.TargetLang
			}
			fmt.Printf("❌ %s → %s (%s - retry %d/3)\n",
				fileName, langName, err.Error, err.Retries)
		}
	}
}

func (t *ProgressTracker) displayFinalStats(successCount, errorCount int) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📈 Final Statistics:")

	elapsed := time.Since(t.startTime)
	successRate := float64(successCount) / float64(successCount+errorCount) * 100

	fmt.Printf("⏱️  Runtime: %v\n", elapsed.Round(time.Second))
	fmt.Printf("🏃‍♂️ Average speed: %.1f translations/sec\n", t.jobsPerSecond)
	fmt.Printf("🎯 Success rate: %.1f%% (%d/%d)\n", successRate, successCount, successCount+errorCount)
	fmt.Printf("💾 Files created: %d\n", successCount)
	fmt.Printf("🔄 Total errors: %d\n", errorCount)

	if errorCount == 0 {
		fmt.Println("\n🎉 Translation complete! All insights translated to all languages.")
	} else {
		fmt.Printf("\n⚠️  Translation completed with %d errors. Check failed files manually.\n", errorCount)
	}
}

func generateProgressBar(percentage, width int) string {
	if percentage > 100 {
		percentage = 100
	}
	if percentage < 0 {
		percentage = 0
	}

	filled := (percentage * width) / 100
	bar := strings.Repeat("█", filled) + strings.Repeat("▒", width-filled)

	return bar
}
