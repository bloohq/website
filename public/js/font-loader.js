/**
 * Dynamic Font Loader
 * Loads language-specific fonts only when needed to optimize bandwidth
 */
window.FontLoader = {
    /**
     * Initialize font loader and check current language
     */
    init() {
        const currentLang = this.getCurrentLanguage();
        console.log('🔤 Font loader initialized for language:', currentLang);
        
        // Also check the data-language attribute as a fallback
        const bodyLang = document.body.getAttribute('data-language');
        const actualLang = currentLang || bodyLang;
        
        console.log('🔤 Body data-language:', bodyLang);
        console.log('🔤 Using language:', actualLang);
        
        if (actualLang === 'km') {
            this.loadKhmerFont();
        }
    },
    
    /**
     * Get current language from body attribute or URL
     * @returns {string} Current language code
     */
    getCurrentLanguage() {
        // First check body attribute
        const bodyLang = document.body.getAttribute('data-language');
        if (bodyLang) return bodyLang;
        
        // Fallback to URL parsing
        const path = window.location.pathname;
        const langMatch = path.match(/^\/([a-z]{2}|[a-z]{2}-[A-Z]{2})(\/|$)/);
        return langMatch ? langMatch[1] : 'en';
    },
    
    /**
     * Load Khmer font dynamically
     */
    loadKhmerFont() {
        // Check if already loaded
        if (document.getElementById('khmer-font-face')) {
            console.log('✅ Khmer font already loaded');
            document.body.classList.add('lang-km');
            return;
        }
        
        console.log('📥 Loading Khmer font...');
        
        // Create and inject @font-face style
        const style = document.createElement('style');
        style.id = 'khmer-font-face';
        style.textContent = `
            @font-face {
                font-family: 'Noto Sans Khmer';
                src: url('/font/NotoSansKhmer-VariableFont_wdth,wght.ttf') format('truetype');
                font-weight: 100 900;
                font-style: normal;
                font-display: swap;
            }
        `;
        document.head.appendChild(style);
        
        // Add language class to body
        document.body.classList.add('lang-km');
        
        // Preload the font to ensure it's downloaded
        const preload = document.createElement('link');
        preload.rel = 'preload';
        preload.as = 'font';
        preload.type = 'font/ttf';
        preload.href = '/font/NotoSansKhmer-VariableFont_wdth,wght.ttf';
        preload.crossOrigin = 'anonymous';
        document.head.appendChild(preload);
        
        console.log('✅ Khmer font loaded successfully');
    },
    
    /**
     * Unload Khmer font to free resources
     */
    unloadKhmerFont() {
        console.log('🗑️ Unloading Khmer font...');
        
        // Remove font face style
        const fontFace = document.getElementById('khmer-font-face');
        if (fontFace) {
            fontFace.remove();
        }
        
        // Remove preload link
        const preloadLinks = document.querySelectorAll('link[href*="NotoSansKhmer"]');
        preloadLinks.forEach(link => link.remove());
        
        // Remove language class
        document.body.classList.remove('lang-km');
        
        console.log('✅ Khmer font unloaded');
    },
    
    /**
     * Handle language switch
     * @param {string} newLang - The new language code
     */
    handleLanguageSwitch(newLang) {
        console.log('🔄 Font loader handling language switch to:', newLang);
        
        // Update body attribute for future reference
        document.body.setAttribute('data-language', newLang);
        
        if (newLang === 'km') {
            this.loadKhmerFont();
        } else {
            // Unload Khmer font if it was loaded
            if (document.getElementById('khmer-font-face')) {
                this.unloadKhmerFont();
            }
        }
    },
    
    /**
     * Check if a font is loaded
     * @param {string} fontFamily - Font family name to check
     * @returns {boolean} Whether the font is loaded
     */
    isFontLoaded(fontFamily) {
        // Use CSS Font Loading API if available
        if ('fonts' in document) {
            return document.fonts.check(`12px "${fontFamily}"`);
        }
        
        // Fallback: check if style element exists
        if (fontFamily === 'Noto Sans Khmer') {
            return !!document.getElementById('khmer-font-face');
        }
        
        return false;
    }
};