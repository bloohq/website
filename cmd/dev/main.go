package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Create a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// WaitGroup to ensure all goroutines finish
	var wg sync.WaitGroup

	// Start Tailwind CSS in watch mode (silent)
	tailwindCmd := exec.Command("./tailwindcss", "-i", "public/css/input.css", "-o", "public/css/style.css", "--watch", "--minify")
	tailwindCmd.Stdout = io.Discard
	tailwindCmd.Stderr = io.Discard

	if err := tailwindCmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start Tailwind CSS: %v\n", err)
		os.Exit(1)
	}

	// Give Tailwind a moment to start
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\033[32m✅ Tailwind CSS started in watch mode\033[0m")

	// Start Air (show output)
	airCmd := exec.Command("air")
	airCmd.Stdout = os.Stdout
	airCmd.Stderr = os.Stderr
	airCmd.Stdin = os.Stdin

	if err := airCmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start Air: %v\n", err)
		// Kill Tailwind before exiting
		tailwindCmd.Process.Kill()
		os.Exit(1)
	}

	// Goroutine to handle cleanup on interrupt
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan
		
		fmt.Println("\n\033[33mShutting down...\033[0m")
		
		// Kill both processes
		if airCmd.Process != nil {
			airCmd.Process.Kill()
		}
		if tailwindCmd.Process != nil {
			tailwindCmd.Process.Kill()
		}
	}()

	// Wait for Air to finish (it's the main process we're watching)
	airErr := airCmd.Wait()
	
	// Clean up Tailwind
	if tailwindCmd.Process != nil {
		tailwindCmd.Process.Kill()
	}

	// Close the signal channel to trigger cleanup goroutine to finish
	close(sigChan)
	wg.Wait()

	// Exit with Air's exit code
	if airErr != nil {
		if exitErr, ok := airErr.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}