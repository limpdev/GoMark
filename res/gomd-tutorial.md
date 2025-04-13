![](/img/cover-61.jpg)

#    ➢ `gomarkdown`

- This article describes an advanced markdown processing in Go using [gomarkdown/markdown](https://github.com/gomarkdown/markdown) library.

All the code examples are available at [https://github.com/gomarkdown/markdown/tree/master/examples](https://github.com/gomarkdown/markdown/tree/master/examples)

# Basics

Here’s a good start for markdown =&gt; HTML conversion:

```go
func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
```

[Try it online](https://arslexis.io/goplayground/#txO7hJ-ibeU)

Basic markdown syntax is very limited and there are many extensions that provide additional feature, like tables.

HTML render is customizable as well.

Here we create markdown parser and HTML renderer with common flags plus some extensions.

Here are all available options for parser and HTML renderer:

- [https://pkg.go.dev/github.com/gomarkdown/markdown/parser#Extensions](https://pkg.go.dev/github.com/gomarkdown/markdown/parser#Extensions) : they change how parser interprets markdown
- [https://pkg.go.dev/github.com/gomarkdown/markdown/html#Flags](https://pkg.go.dev/github.com/gomarkdown/markdown/html#Flags) : they change generated HTML

# Advanced processing

Markdown to HTML processing works like this:

- `github.com/gomarkdown/markdown/parser` parses markdown and generates AST tree as defined in `github.com/gomarkdown/markdown/ast`
- `github.com/gomarkdown/markdown/html` implements HTML renderer that takes AST tree and generates HTML

There are options for even more control:

- customize HTML generator by providing `html.Renderer.RenderNodeHook`. You re-use most of the `html.Renderer` and change rendering of just some `ast.Node`.
- fork `github.com/gomarkdown/markdown/html` and make changes you want to HTML renderer
- modify ast tree after parsing but before rendering
- customize the parser, define your own `ast.Node` types, add them to the tree while parsing and customize renderer to render those nodes as you want
- pre-process markdown before sending to the parser

## `ast.Node`

You need to understand AST tree. Start by skimming [https://github.com/gomarkdown/markdown/blob/master/ast/node.go](https://github.com/gomarkdown/markdown/blob/master/ast/node.go).

`ast.Node` is an interface so you can create your own nodes as long as you implement the interface.

The are two types of nodes:

- container node has an array of children nodes e.g. a `List` contains `ListItem` nodes
- a leaf node doesn’t have children, just content

We have `ast.Leaf` and `ast.Container` to make it easy to implement custom nodes:

```go
type MyCustomLeafNode struct {
  ast.Leaf
  // .. additional data for this node
}

type MyCustomContainerNode struct {
  ast.Container
  // .. additional data for this node
}
```

To render your custom node `ast.Node` to HTML you provide a render hook function that will be structured like this:

```go
func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
  // you must render custom nodes because html.Renderer doesn't understand them
  if leafNode, ok := node.(*MyCustomLeafNode); ok {
    renderMyLeafNode(w, lefNode, entering)
    return ast.GoToNext, true
  }

  if containerNode, ok := node.(*MyCustomContainerNode); ok {
    renderMyContainerNode(w, containerNode, entering)
    return ast.GoToNext, true
  }

  // you can also over-ride rendering of some specific nodes that html.Renderer would render
  if image, ok := node.(*ast.Image); ok {
    renderImage(w, image, entering)
    return ast.GoToNext, true
  }

  // return false to tell html.Renderer to use default render
  return ast.GoToNext, false
}
```

You should always return `ast.GoToNext`.

You return `true` to indicate that you’ve rendered the node. Returning `false` tells `html.Renderer` to use default rendering.

For container nodes we typically need to render something before rendering children and after rendering children.

That’s why we need `entering` argument. The simplest rendering of `*ast.Paragraph` would be:

```go
func myRenderParagraph(w io.Writer, p *ast.Paragraph, entering bool) {
  if entering {
    io.WriteString(w, "<p>")
  } else {
    io.WriteString(w, "</p>")
  }
}
```

`html.Renderer` takes care of recursively rendering children.

For `ast.Leaf` nodes you only render on entering:

```go
func myHr(w io.Writer, p *ast.HorizontalRule, entering bool) {
  if entering {
    io.WriteString(w, "<hr/>")
  }
}
```

Skim [render.go](https://github.com/gomarkdown/markdown/blob/master/html/renderer.go) to see how different nodes are rendered to HTML.

## Customizing HTML renderer

To re-use most of `html.Renderer` but only over-ride rendering of a few nodes, you can provide a render hook.

Here’s an example of a simple hook that renders `<div>{children}</div>` instead of `<p>{children}</p>` for a `*ast.Paragraph` node.

```go
// an actual rendering of Paragraph is more complicated
func renderParagraph(w io.Writer, p *ast.Paragraph, entering bool) {
  if entering {
    io.WriteString(w, "<div>")
  } else {
    io.WriteString(w, "</div>")
  }
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
  if para, ok := node.(*ast.Paragraph); ok {
    renderParagraph(w, para, entering)
    return ast.GoToNext, true
  }
  return ast.GoToNext, false
}

func newCustomizedRender() *html.Renderer {
  opts := html.RendererOptions{
    RenderNodeHook: myRenderHook,
  }
  return html.NewRenderer(opts)
}
```

[Try it online](https://arslexis.io/goplayground/#yFRIWRiu-KL)

If a render hook needs access to more information than `io.Writer` and `ast.Node`, we can capture it in a closure:

```go
import (
  "github.com/gomarkdown/markdown/html"
)

type renderData struct {
  // ... data needed for render hook function
}

func makeRenderHook(data *renderData)  html.RenderNodeFunc {
  return myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
    // has access to data
  }
}

func newCustomizedRender() *html.Renderer {
  data := &renderData{}
  opts := html.RendererOptions{
    RenderNodeHook: makeRenderHook(data),
  }
  return html.NewRenderer(opts)
}
```

## Modify ast tree

The structure of the code would be:

```go
func modifyAst(root ast.Node) ast.Node {
  // ... tweak AST tree as needed
  return root
}

var mds = `[link](http://example.com)`

func modifyAstExample() {
	md := []byte(mds)

	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	doc = modifyAst(doc)

	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	html := markdown.Render(doc, renderer)

	fmt.Printf("-- Markdown:\n%s\n\n--- HTML:\n%s\n", md, html)
}
```

You can re-arrange the tree: add, remove, rearrange nodes or add / remove information from nodes.

Working with trees is tricky so you’ll probably use `ast.Print(io.Writer, ast.Node)` to pretty-print and understand AST tree, both before and after making changes.

I won’t cover changing the structure of the tree but here’s an example that adds `target="_blank"` attribute to link (`<a>`) nodes that go outside of my website and adds custom `blog-img` class to all image (`<img>`) nodes.

```go
func modifyAst(doc ast.Node) ast.Node {
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if img, ok := node.(*ast.Image); ok && entering {
			attr := img.Attribute
			if attr == nil {
				attr = &ast.Attribute{}
			}
			// TODO: might be duplicate
			attr.Classes = append(attr.Classes, []byte("blog-img"))
			img.Attribute = attr
		}

		if link, ok := node.(*ast.Link); ok && entering {
			isExternalURI := func(uri string) bool {
				return (strings.HasPrefix(uri, "https://") || strings.HasPrefix(uri, "http://")) && !strings.Contains(uri, "blog.kowalczyk.info")
			}
			if isExternalURI(string(link.Destination)) {
				link.AdditionalAttributes = append(link.AdditionalAttributes, `target="_blank"`)
			}
		}

		return ast.GoToNext
	})
	return doc
}
```

[Try it online](https://arslexis.io/goplayground/#2yV5-HDKBUV)

Important parts:

- `ast.WalkFunc` is a helper function that recursively calls a callback function on every node in AST
- we do modifications only once e.g. when `entering` is true
- all `Container` nodes have (optional) `*Attribute` which contains HTML id attribute, class names and attributes, which we can modify
- some nodes can have additional data we can modify e.g. `*ast.Link` has `AdditionalAttributes` which is array of attributes.

```go
type Attribute struct {
	ID      []byte
	Classes [][]byte
	Attrs   map[string][]byte
}
```

## Custom markdown parser, custom `ast.Node`

You can also extend parser to recognize additional syntax, not present in markdown.

Here’s an example of a parser extension that recognizes the following:

```go
:gallery
/img/image-1.png
/img/image-2.png
/img/image-3.png

Rest of document.
```

A gallery is a list of image urls. In generated HTML this will be a show as an image gallery.

First define a custom leaf node type:

```go
type Gallery struct {
	ast.Leaf
	ImageURLS []string
}
```

Then customize parser with a block parsing hook function which gets first dibs at parsing block text.

Block means: it only gets called at the beginning of each text block. Text blocks are separated by newlines.

It’s not possible to do custom inline text parsing.

If parser hook function recognizes its custom syntax, it returns `ast.Node` it generated (`Gallery` in this case), number of bytes consumed and inner content to be recursively parsed and inserted as children of `ast.Container` node.

Number of bytes consumed allows parser to skip the part parsed by your custom hook.

```go
func parserHook(data []byte) (ast.Node, []byte, int) {
	if node, d, n := parseGallery(data); node != nil {
		return node, d, n
	}
	return nil, nil, 0
}

func newMarkdownParser() *parser.Parser {
	extensions := parser.CommonExtensions
	p := parser.NewWithExtensions(extensions)
	p.Opts.ParserHook = parserHook
	return p
}
```

Here’s the parser for `:gallery` syntax.

If current text starts with `:gallery\n` we expect a list of urls followed by an empty line.

```go
var gallery = []byte(":gallery\n")

func parseGallery(data []byte) (ast.Node, []byte, int) {
	if !bytes.HasPrefix(data, gallery) {
		return nil, nil, 0
	}
	i := len(gallery)
  // find empty line that ends the block
  // TODO: should also consider end of document
	end := bytes.Index(data[i:], []byte("\n\n"))
	if end < 0 {
		return nil, data, 0
	}
	end = end + i
	// split into lines, each line is an image URL
	lines := string(data[i:end])
	parts := strings.Split(lines, "\n")
	res := &Gallery{
		ImageURLS: parts,
	}
	return res, nil, end
}
```

[Try it online](https://arslexis.io/goplayground/#9fqKwRbuJ04)

I will not cover how it gets rendered as HTML because it’s a very custom solution with lots of HTML obscuring the big picture.

## Syntax highlighting

Code blocks are much better with syntax highlighting.

We can use [github.com/alecthomas/chroma](https://github.com/alecthomas/chroma) library to generate HTML with syntax highlighting for many languages.

We hook it up to HTML renderer with render hook function like described above.

```go
import (
	"fmt"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
)

func init() {
	htmlFormatter = html.New(html.WithClasses(true), html.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}
	styleName := "monokailight"
	highlightStyle = styles.Get(styleName)
	if highlightStyle == nil {
		panic(fmt.Sprintf("didn't find style '%s'", styleName))
	}
}

// based on https://github.com/alecthomas/chroma/blob/master/quick/quick.go
func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return htmlFormatter.Format(w, highlightStyle, it)
}

// an actual rendering of Paragraph is more complicated
func renderCode(w io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	defaultLang := ""
	lang := string(codeBlock.Info)
	htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}
```

[Try it online](https://arslexis.io/goplayground/#Bk0zTvrzUDR)

You’ll have to include chroma CSS as HTML generation marks up nodes with chroma CSS classes.

## Pre-process markdown before parsing

Imagine you want to add ability to include markdown files.

You can build a parser extension that e.g. recognizes this syntax:

```go
@include "foo/bar.md"
```

It’ll get complicated if you want to add more advanced functionality, like loops, variables etc.

But [`template/text`](https://pkg.go.dev/text/template) library already implement such features.

You can pre-process markdown with one of the many templating Go libraries before sending it to the parser.

[go](/tag/go) [programming](/tag/programming)

Mar 11 2023

try my software

[Edna](https://edna.arslexis.io/?utm_source=blog)

note taking app for developers and power users

optimized for speed

cross between Obsidian and Notational Velocity

[SumatraPDF](https://www.sumatrapdfreader.org/free-pdf-reader?utm_source=blog "SumatraPDF reader for Windows")

small, fast, free PDF / ePub / comic book reader for Windows
