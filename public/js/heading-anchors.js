/**
 * Heading Anchor Links Utilities
 * Adds clickable anchor links to H2 headings in markdown content
 */
window.HeadingAnchors = {
    /**
     * Initialize heading anchors on all H2 elements
     */
    init() {
        this.addAnchors();
    },

    /**
     * Add anchor links to all H2 elements that don't already have them
     */
    addAnchors() {
        const headings = document.querySelectorAll('.prose h2[id]:not([data-anchor-added])');
        
        headings.forEach(heading => {
            heading.setAttribute('data-anchor-added', 'true');
            
            // Create anchor element
            const anchor = document.createElement('a');
            anchor.className = 'heading-anchor';
            anchor.href = '#' + heading.id;
            anchor.innerHTML = '#';
            anchor.title = 'Copy link to this section';
            
            // Prevent default link behavior and copy to clipboard instead
            anchor.addEventListener('click', async (e) => {
                e.preventDefault();
                
                const headingId = heading.id;
                const fullUrl = window.location.origin + window.location.pathname + '#' + headingId;
                
                try {
                    await navigator.clipboard.writeText(fullUrl);
                    
                    // Show success feedback
                    anchor.classList.add('copied');
                    anchor.innerHTML = '✓';
                    anchor.title = 'Copied!';
                    
                    setTimeout(() => {
                        anchor.classList.remove('copied');
                        anchor.innerHTML = '#';
                        anchor.title = 'Copy link to this section';
                    }, 2000);
                } catch (err) {
                    console.error('Failed to copy heading link:', err);
                }
            });
            
            // Insert anchor as first child of heading
            heading.insertBefore(anchor, heading.firstChild);
        });
    }
};