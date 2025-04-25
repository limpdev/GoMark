# Converting Your Markdown Tool to WebAssembly

Yes, it's definitely possible to convert your Go-based Markdown to HTML tool into a WebAssembly (WASM) application. Here's how you could approach this:

## Key Considerations

1. **WASM Compatibility**: Your code uses standard Go libraries and goldmark which are WASM-compatible.

2. **Browser Environment Differences**:
   - No direct file system access
   - No command-line arguments
   - Need to handle I/O via JavaScript interop

## Implementation Approach

### 1. Modify the Core Functionality

You'll want to create a WASM-exported function that takes markdown input and returns HTML:

```go
// main.go
package main

import (
	"bytes"
	"syscall/js"
	
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func main() {
	// Register the function to make it available in JavaScript
	js.Global().Set("markdownToHTML", js.FuncOf(markdownToHTML))
	
	// Keep the program running
	select {}
}

func markdownToHTML(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Error: Expected exactly 1 argument (markdown string)"
	}
	
	markdownInput := args[0].String()
	
	// Convert markdown to HTML
	md := setupGoldmark()
	var buf bytes.Buffer
	if err := md.Convert([]byte(markdownInput), &buf); err != nil {
		return "Error: " + err.Error()
	}
	
	// Return the HTML content
	return buf.String()
}

// setupGoldmark remains the same as in your original code
```

### 2. Build for WASM

Compile with:
```
GOOS=js GOARCH=wasm go build -o markdown.wasm
```

### 3. Create a JavaScript Wrapper

You'll need an HTML/JS frontend to interact with your WASM module:

```html
<head>
    <title>Markdown to HTML Converter</title>
    <script src="wasm_exec.js"></script>
    <style>
        #editor { width: 100%; height: 300px; }
        #preview { border: 1px solid #ccc; padding: 10px; min-height: 300px; }
    </style>
</head>
<body>
    <h1>Markdown to HTML Converter</h1>
    <textarea id="editor" placeholder="Enter markdown here..."></textarea>
    <button id="convert">Convert</button>
    <div id="preview"></div>

    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("markdown.wasm"), go.importObject)
            .then(result => {
                go.run(result.instance);
                
                document.getElementById('convert').addEventListener('click', () => {
                    const markdown = document.getElementById('editor').value;
                    const html = markdownToHTML(markdown);
                    document.getElementById('preview').innerHTML = html;
                });
            });
    </script>
</body>
```

## Enhancements for Your Specific Code

1. **Template Handling**: Since you won't have filesystem access in the browser:
   - Either hardcode templates in Go
   - Or pass template content from JavaScript

2. **Special Character Transformations**: Your `transformSpecialCharacters` function will work as-is in WASM.

3. **Output Options**:
   - Instead of writing to a file, return HTML to JavaScript
   - Let the browser handle downloading if needed

## Deployment Steps

1. Compile your Go code to WASM
2. Include the Go WASM runtime (`wasm_exec.js`)
3. Create a simple HTML interface
4. Serve all files via a web server

## Potential Challenges

1. **Performance**: WASM can be fast, but large documents might cause UI freezes. Consider:
   - Web Workers for background processing
   - Chunked processing for very large documents

2. **Size**: The WASM binary might be large. You can:
   - Use TinyGo for smaller builds
   - Implement compression

3. **Browser Compatibility**: Most modern browsers support WASM, but:
   - Test across different browsers
   - Provide a fallback for older browsers

---

### Original Golang Converter Script

> Needs revisions to make web-compatible...

```go
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
		templatePath = "C:/Users/drewg/bin/exe/GoMark/templates/index.html"
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
```