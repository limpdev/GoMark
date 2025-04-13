package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// RunGoldhub is the main entry point for the Goldmark markdown converter
func RunGoldhub() error {
	var input io.ReadCloser // Input source: file or stdin
	var inputName string    // For messages
	var mdFilePath string   // Store the original file path if provided

	// --- Determine Input Source ---
	switch len(os.Args) {
	case 1:
		// Read from stdin
		fmt.Fprintln(os.Stderr, "No filename provided. Reading Markdown from stdin...")
		input = os.Stdin
		inputName = "stdin"
		// No mdFilePath when reading from stdin
	case 2:
		// Read from file
		mdFilePath = os.Args[1] // Store the filename
		inputName = mdFilePath
		file, err := os.Open(mdFilePath) // Use the stored path
		if err != nil {
			return fmt.Errorf("error opening file '%s': %v", mdFilePath, err)
		}
		input = file
	default:
		// Usage error
		return fmt.Errorf("usage: %s [mdFile]", os.Args[0])
	}

	// Ensure the input source is closed, important for files
	defer func() {
		if err := input.Close(); err != nil {
			if inputName != "stdin" {
				fmt.Fprintf(os.Stderr, "Warn: error closing input source %s: %v\n", inputName, err)
			}
		}
	}()

	// --- Read Markdown Content ---
	fmt.Fprintf(os.Stderr, "Processing Markdown from %s...\n", inputName)
	// Read all content from the determined input (file or stdin)
	markdownBytes, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("error reading from %s: %v", inputName, err)
	}

	// Apply the transformations defined above
	transformedMarkdown := transformSpecialCharacters(markdownBytes)

	// --- Determine Output Filename and Title ---
	var mdNude string // Base name for title and output file
	var outFile string
	if mdFilePath != "" {
		// Input was from a file, use its name
		mdFileName := filepath.Base(mdFilePath)
		extension := filepath.Ext(mdFileName)
		if extension == "" {
			mdNude = mdFileName // Handle files without extension
		} else {
			mdNude = mdFileName[:len(mdFileName)-len(extension)]
		}
	} else {
		// Input was from stdin, use a default name
		mdNude = "output" // Default base name
	}
	outFile = mdNude + ".html" // Construct output filename

	// --- CONVERT MARKDOWN TO HTML using Goldmark ---
	md := setupGoldmark()
	var buf bytes.Buffer
	if err := md.Convert(transformedMarkdown, &buf); err != nil {
		return fmt.Errorf("error converting markdown: %v", err)
	}
	htmlContent := buf.Bytes()

	// --- Find Template Path ---
	var templatePath string
	// First check environment variable
	if envPath := os.Getenv("MARKDOWN_TEMPLATE_PATH"); envPath != "" {
		templatePath = envPath
	} else {
		// Then try relative to current directory
		templatePath = "{ESTABLISH YOUR TEMPLATE'S LOCAL PATH HERE}"
	}

	// --- Load HTML Template ---
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read template at %s: %v\n", templatePath, err)
		fmt.Fprintln(os.Stderr, "Using simple built-in template as fallback")
		// Fallback to a built-in template
		templateContent = []byte(`<!DOCTYPE html>
<html>
<head>
	<title>{{TITLE}}</title>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body { font-family: Arial, sans-serif; line-height: 1.6; max-width: 800px; margin: 0 auto; padding: 1em; }
		pre { background: #f4f4f4; padding: 1em; overflow: auto; }
		code { background: #f4f4f4; padding: 0.2em 0.4em; }
	</style>
</head>
<markdown-body>
  <article>
	{{CONTENT}}
  </article>
</markdown-body>
</html>`)
	} else {
		fmt.Fprintf(os.Stderr, "Using template from: %s\n", templatePath)
	}

	// --- Inject Content into Template ---
	htmlPre := strings.Replace(string(templateContent), "{{TITLE}}", mdNude, 1)
	htmlDoc := strings.Replace(htmlPre, "{{CONTENT}}", string(htmlContent), 1)

	// --- Write Output HTML File ---
	err = os.WriteFile(outFile, []byte(htmlDoc), 0644)
	if err != nil {
		return fmt.Errorf("error writing the output HTML file %s: %v", outFile, err)
	}

	fmt.Fprintf(os.Stderr, "Successfully generated %s from %s.\n", outFile, inputName)
	return nil
}

// setupGoldmark configures and returns a Goldmark instance with extensions
func setupGoldmark() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,            // GitHub Flavored Markdown
			extension.DefinitionList, // Definition lists
			extension.Footnote,       // Footnotes
			extension.Typographer,    // Smart typography (quotes, dashes)
			extension.Strikethrough,  // Strikethrough
			extension.Linkify,        // Auto-link URLs
			extension.Table,          // Tables
			extension.TaskList,       // Task/checkbox lists
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(), // Auto-generate heading IDs
			parser.WithAttribute(),     // Custom attributes for elements
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(), // Preserve line breaks
			html.WithXHTML(),     // XHTML output
			html.WithUnsafe(),    // Allow raw HTML
		),
	)
}

// transformSpecialCharacters replaces text patterns with special characters
func transformSpecialCharacters(input []byte) []byte {
	text := string(input)
	// Create a map of transformations
	transformations := map[string]string{
		"->":   "⇾",
		"<-":   "⇽",
		"<->":  "⇿",
		"=>":   "≥",
		"<=":   "≤",
		"<=>":  "↔",
		"--":   "—", // em dash
		"...":  "…", // ellipsis
		"(c)":  "©", // copyright
		"(tm)": "™", // trademark
		"(r)":  "®", // registered trademark
		":)":   "🙂",
		":(":   "🙁",
		":D":   "😀",
		";)":   "😉",
		"<3":   "❤️",
		"+-":   "±",
		"!=":   "≠",
		"^2":   "²",
		"^3":   "³",
		"1/2":  "½",
		"1/3":  "⅓",
		"2/3":  "⅔",
		"1/4":  "¼",
		"3/4":  "¾",
		"~~":   "≈",
		"==":   "≡",
		"<<":   "«",
		">>":   "»",
		"-A-":  "⩜",
		// Add more transformations as needed
	}
	// Apply each transformation
	for from, to := range transformations {
		text = strings.ReplaceAll(text, from, to)
	}
	return []byte(text)
}
