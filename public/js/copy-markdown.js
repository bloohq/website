/**
 * Copy Markdown to Clipboard
 * Provides functionality to copy page markdown content for LLMs
 */

(function() {
    'use strict';

    /**
     * Copy markdown content from the hidden element to clipboard
     * @returns {Promise<void>}
     */
    window.copyMarkdownToClipboard = async function() {
        try {
            // Get the markdown content from the hidden script tag
            const markdownElement = document.getElementById('page-markdown');

            if (!markdownElement) {
                console.error('Markdown content not found on this page');
                throw new Error('Markdown content not available');
            }

            const markdownContent = markdownElement.textContent;

            if (!markdownContent || markdownContent.trim() === '') {
                console.error('Markdown content is empty');
                throw new Error('No markdown content to copy');
            }

            // Use the Clipboard API to copy the text
            await navigator.clipboard.writeText(markdownContent);

            console.log('Markdown copied to clipboard successfully');
        } catch (error) {
            console.error('Failed to copy markdown:', error);

            // Fallback method for older browsers
            try {
                const markdownElement = document.getElementById('page-markdown');
                if (!markdownElement) {
                    throw new Error('Markdown content not available');
                }

                const textArea = document.createElement('textarea');
                textArea.value = markdownElement.textContent;
                textArea.style.position = 'fixed';
                textArea.style.left = '-999999px';
                textArea.style.top = '-999999px';
                document.body.appendChild(textArea);
                textArea.focus();
                textArea.select();

                const successful = document.execCommand('copy');
                document.body.removeChild(textArea);

                if (!successful) {
                    throw new Error('Fallback copy failed');
                }

                console.log('Markdown copied using fallback method');
            } catch (fallbackError) {
                console.error('Fallback copy also failed:', fallbackError);
                alert('Failed to copy markdown. Please try again or use Ctrl+C/Cmd+C manually.');
                throw fallbackError;
            }
        }
    };

    // Log that the script has loaded
    console.log('Copy markdown functionality loaded');
})();
