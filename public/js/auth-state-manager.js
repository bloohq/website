/**
 * Auth State Manager
 * Centralized manager for auth state across the marketing website
 * Polls the auth cookie once and notifies all subscribers of changes
 */
window.AuthStateManager = {
    state: null,
    subscribers: [],
    intervalId: null,
    POLL_INTERVAL: 3000, // 3 seconds
    
    /**
     * Initialize the auth state manager
     */
    init() {
        // Initial read
        this.updateState();
        
        // Start polling
        this.intervalId = setInterval(() => this.updateState(), this.POLL_INTERVAL);
        
        // Clean up on page unload
        window.addEventListener('beforeunload', () => this.destroy());
        
        // Also clean up on visibility change (when tab is hidden)
        document.addEventListener('visibilitychange', () => {
            if (document.hidden) {
                this.pause();
            } else {
                this.resume();
            }
        });
    },
    
    /**
     * Update state from cookie
     */
    updateState() {
        if (!window.AuthCookie) return;
        
        const newState = window.AuthCookie.read();
        
        // Only notify if state actually changed
        if (JSON.stringify(newState) !== JSON.stringify(this.state)) {
            this.state = newState;
            this.notifySubscribers();
        }
    },
    
    /**
     * Subscribe to auth state changes
     * @param {Function} callback - Function to call with auth state
     * @returns {Function} Unsubscribe function
     */
    subscribe(callback) {
        this.subscribers.push(callback);
        
        // Immediately call with current state
        if (callback) {
            callback(this.state);
        }
        
        // Return unsubscribe function
        return () => {
            this.subscribers = this.subscribers.filter(cb => cb !== callback);
        };
    },
    
    /**
     * Notify all subscribers of state change
     */
    notifySubscribers() {
        this.subscribers.forEach(callback => {
            if (callback) {
                callback(this.state);
            }
        });
    },
    
    /**
     * Get current auth state
     * @returns {Object|null} Current auth state
     */
    getState() {
        return this.state;
    },
    
    /**
     * Pause polling (when tab is hidden)
     */
    pause() {
        if (this.intervalId) {
            clearInterval(this.intervalId);
            this.intervalId = null;
        }
    },
    
    /**
     * Resume polling (when tab is visible)
     */
    resume() {
        if (!this.intervalId) {
            // Update immediately on resume
            this.updateState();
            // Restart polling
            this.intervalId = setInterval(() => this.updateState(), this.POLL_INTERVAL);
        }
    },
    
    /**
     * Clean up and destroy the manager
     */
    destroy() {
        this.pause();
        this.subscribers = [];
        this.state = null;
    }
};