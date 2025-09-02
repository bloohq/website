/**
 * Partner Earnings Calculator Component
 * Interactive tool for calculating partner earnings potential
 */

// Define the function immediately to ensure it's available
window.partnerROICalculator = function() {
    return {
        // Input variables
        numClients: 10,
        avgDealSize: 50000,
        clientGrowthRate: 25,
        retentionRate: 90,
        activeScenario: 'moderate', // Track which scenario is active
        
        // Scenario presets
        setScenario(type) {
            this.activeScenario = type; // Update active scenario
            switch(type) {
                case 'conservative':
                    this.numClients = 5;
                    this.avgDealSize = 25000;
                    this.clientGrowthRate = 10;
                    this.retentionRate = 85;
                    break;
                case 'moderate':
                    this.numClients = 10;
                    this.avgDealSize = 50000;
                    this.clientGrowthRate = 25;
                    this.retentionRate = 90;
                    break;
                case 'aggressive':
                    this.numClients = 20;
                    this.avgDealSize = 100000;
                    this.clientGrowthRate = 50;
                    this.retentionRate = 95;
                    break;
            }
        },
        
        // Helper function to format numbers
        formatNumber(num) {
            return Math.round(num).toLocaleString();
        },
        
        // Calculate yearly earnings
        get yearlyEarnings() {
            let earnings = [];
            let activeClients = this.numClients;
            const growthRate = this.clientGrowthRate / 100;
            const retention = this.retentionRate / 100;
            
            for (let year = 1; year <= 10; year++) {
                // Apply retention to existing clients
                if (year > 1) {
                    activeClients = activeClients * retention;
                }
                
                // Add new clients based on growth
                if (year > 1) {
                    activeClients += this.numClients * Math.pow(1 + growthRate, year - 1) - this.numClients * Math.pow(1 + growthRate, year - 2);
                }
                
                // Calculate commission based on year
                const commissionRate = year <= 2 ? 0.5 : 0.25;
                const yearEarnings = activeClients * this.avgDealSize * commissionRate;
                
                earnings.push(yearEarnings);
            }
            
            return earnings;
        },
        
        get year1Earnings() {
            return this.yearlyEarnings[0] || 0;
        },
        
        get year2Earnings() {
            return this.yearlyEarnings[1] || 0;
        },
        
        get years3to10Earnings() {
            return this.yearlyEarnings.slice(2).reduce((sum, val) => sum + val, 0);
        },
        
        get totalEarnings() {
            return this.yearlyEarnings.reduce((sum, val) => sum + val, 0);
        }
    }
};