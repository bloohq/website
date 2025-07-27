# Development Server Manager

This Go program manages the development environment for the Blue website by running both Air (Go hot-reload) and Tailwind CSS in watch mode simultaneously.

## What it does

- Starts Tailwind CSS in watch mode (silently in the background)
- Shows a confirmation message when Tailwind starts successfully
- Runs Air with full output display
- Handles graceful shutdown of both processes when you press Ctrl+C

## How to set up the `start` command

### 1. Create the shell script (if it doesn't exist)

In the root directory of blue-new-website:

```bash
# Create the start script
cat > start << 'EOF'
#!/usr/bin/env bash

# Change to the directory where this script is located
cd "$(dirname "$0")"

# Run the Go development manager
go run cmd/dev/main.go
EOF

# Make it executable
chmod +x start
```

### 2. Add the shell alias

Add this line to your shell configuration file (`~/.zshrc` for zsh or `~/.bashrc` for bash):

```bash
# Blue website development command
alias start='cd /Users/manny/Blue/blue-new-website && ./start'
```

Then reload your shell:

```bash
source ~/.zshrc  # or source ~/.bashrc
```

### 3. Usage

From anywhere in your terminal, simply run:

```bash
start
```

This will:
1. Change to the blue-new-website directory
2. Start both Air and Tailwind CSS
3. Show Air's output in your terminal
4. Stop both processes when you press Ctrl+C

## Technical Details

The program uses Go's `os/exec` package to spawn both processes as child processes. Tailwind's output is discarded to keep the terminal clean, while Air's output is piped directly to stdout/stderr for full visibility of the Go application logs.

Signal handling ensures both processes are properly terminated when the program exits, preventing orphaned processes.