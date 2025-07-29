package web

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// SupportedLanguages is the ONLY place to add new languages
// Just add a new language code here and everything else works automatically
var SupportedLanguages = []string{
	"en", // English
	"zh", // 简体中文 (Simplified Chinese)
	"es", // Español (Spanish)
	"fr", // Français (French)
	"de", // Deutsch (German)
	"ja", // 日本語 (Japanese)
	"pt", // Português (Portuguese)
	"ru", // Русский (Russian)
	"ko", // 한국어 (Korean)
	"it", // Italiano (Italian)
	"id", // Indonesian
	"nl", // Nederlands (Dutch)
	"pl", // Polski (Polish)
	"zh-TW", // 繁體中文 (Traditional Chinese)
	"sv", // Svenska (Swedish)
	"km", // ភាសាខ្មែរ (Khmer)
}

// DefaultLanguage is the fallback language
const DefaultLanguage = "en"

// LanguageLocales maps language codes to locale strings for og:locale
var LanguageLocales = map[string]string{
	"en":    "en_US",
	"zh":    "zh_CN",
	"es":    "es_ES",
	"fr":    "fr_FR",
	"de":    "de_DE",
	"ja":    "ja_JP",
	"pt":    "pt_BR",
	"ru":    "ru_RU",
	"ko":    "ko_KR",
	"it":    "it_IT",
	"id":    "id_ID",
	"nl":    "nl_NL",
	"pl":    "pl_PL",
	"zh-TW": "zh_TW",
	"sv":    "sv_SE",
	"km":    "km_KH",
}

// GetLocaleForLanguage returns the locale string for a given language code
func GetLocaleForLanguage(lang string) string {
	if locale, ok := LanguageLocales[lang]; ok {
		return locale
	}
	return "en_US" // Default locale
}

// detectPreferredLanguage detects user's preferred language from cookie or Accept-Language header
func detectPreferredLanguage(r *http.Request) string {
	// 1. Check cookie first (user's explicit choice)
	if cookie, err := r.Cookie("lang"); err == nil {
		// Validate the cookie value is a supported language
		for _, lang := range SupportedLanguages {
			if cookie.Value == lang {
				return lang
			}
		}
	}
	
	// 2. Check Accept-Language header
	acceptLang := r.Header.Get("Accept-Language")
	if acceptLang != "" {
		// Parse Accept-Language header (e.g., "es-ES,es;q=0.9,en;q=0.8")
		preferredLang := parseAcceptLanguage(acceptLang, SupportedLanguages)
		if preferredLang != "" {
			return preferredLang
		}
	}
	
	// 3. Default to English
	return DefaultLanguage
}

// parseAcceptLanguage parses Accept-Language header and returns the best match
func parseAcceptLanguage(header string, supported []string) string {
	type langPref struct {
		lang    string
		quality float64
	}
	
	var prefs []langPref
	
	// Parse each language preference
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		
		// Split language and quality value
		langQ := strings.Split(part, ";")
		lang := strings.TrimSpace(langQ[0])
		quality := 1.0
		
		// Parse quality value if present
		if len(langQ) > 1 {
			qPart := strings.TrimSpace(langQ[1])
			if strings.HasPrefix(qPart, "q=") {
				if q, err := strconv.ParseFloat(qPart[2:], 64); err == nil {
					quality = q
				}
			}
		}
		
		prefs = append(prefs, langPref{lang: lang, quality: quality})
	}
	
	// Sort by quality (highest first)
	sort.Slice(prefs, func(i, j int) bool {
		return prefs[i].quality > prefs[j].quality
	})
	
	// Find the best match
	for _, pref := range prefs {
		prefLower := strings.ToLower(pref.lang)
		
		// Check for exact match first
		for _, supported := range supported {
			if prefLower == strings.ToLower(supported) {
				return supported
			}
		}
		
		// Check for language-only match (e.g., "es-MX" matches "es")
		// But be careful with "zh" vs "zh-TW"
		langOnly := prefLower
		if idx := strings.IndexByte(langOnly, '-'); idx > 0 {
			langOnly = langOnly[:idx]
		}
		
		for _, supported := range supported {
			supportedLower := strings.ToLower(supported)
			// Exact language match (not substring)
			if langOnly == supportedLower {
				// Special case: don't match "zh-TW" to "zh"
				if supported == "zh" && strings.HasPrefix(prefLower, "zh-") && prefLower != "zh-cn" {
					continue
				}
				return supported
			}
		}
	}
	
	return ""
}

// isLanguagePrefix checks if a path segment is a valid language code
func isLanguagePrefix(segment string) bool {
	for _, lang := range SupportedLanguages {
		if segment == lang {
			return true
		}
	}
	return false
}

// extractLanguageFromPath extracts language code and clean path from URL
// Examples:
// "/en/about" -> "en", "/about"
// "/es/pricing" -> "es", "/pricing"
// "/about" -> "", "/about"
func extractLanguageFromPath(path string) (string, string) {
	// Remove leading slash for easier parsing
	trimmed := strings.TrimPrefix(path, "/")
	
	// Split path into segments
	segments := strings.SplitN(trimmed, "/", 2)
	
	// Check if first segment is a language code
	if len(segments) > 0 && isLanguagePrefix(segments[0]) {
		lang := segments[0]
		remainingPath := "/"
		if len(segments) > 1 {
			remainingPath = "/" + segments[1]
		}
		return lang, remainingPath
	}
	
	// No language prefix found
	return "", path
}

// setLanguageCookie sets the language preference cookie
func setLanguageCookie(w http.ResponseWriter, lang string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "lang",
		Value:    lang,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1 year
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}