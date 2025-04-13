package main

import (
	"fmt"
	"os"
	"time"

	"Former/cmd" // INTERNAL

	"github.com/briandowns/spinner"
)

func main() {
	// Show a welcome spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = " 📄 Markdown Converter "
	s.Color("fgHiMagenta")
	s.Start()
	time.Sleep(1000 * time.Millisecond)
	s.Stop()

	// Print version and tool info
	fmt.Fprintf(os.Stderr, "GoMark Conversion Set ∙ v4.2.0\n")
	fmt.Fprintf(os.Stderr, "Parsed by Goldmark🏅\n")

	// Check for command-line flags
	// Here you can implement command-line flags for different modes or options
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		printHelp()
		return
	}

	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		printVersion()
		return
	}

	// Pass control to the Goldmark implementation
	err := cmd.RunGoldhub()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`
∙ GoMark ∙

Usage:
  gomark [options] [file]

Options:
  -h, --help     Show this help message
  -v, --version  Show version information

Arguments:
  file           Path to a markdown file. If not provided, reads from stdin.

Examples:
  gomark document.md        # Convert document.md to document.html
  cat document.md | gohub   # Convert stdin to output.html
  
Environment Variables:
  MARKDOWN_TEMPLATE_PATH   Path to custom HTML template`)
}

func printVersion() {
	fmt.Println("GoMark Markdown Conversion Set ∙ v4.2.0")
	fmt.Println("MIT ∙ Limp Cheney")
}
