# extension

package module ![](/static/shared/icon/content_copy_gm_grey_24dp.svg)

[Version: v0.1.0](SuperSubscript.md)

Opens a new window with list of versions in this module.

Latest

Latest

![Warning](/static/shared/icon/alert_gm_grey_24dp.svg)

This package is not in the latest version of its module.

[Go to latest](/github.com/bowman2001/goldmark-supersubscript)

Published: Feb 11, 2024 License: [MIT](/github.com/bowman2001/goldmark-supersubscript?tab=licenses)

Opens a new window with license information.

[Imports: 8](/github.com/bowman2001/goldmark-supersubscript?tab=imports)

Opens a new window with list of imports.

[Imported by: 0](/github.com/bowman2001/goldmark-supersubscript?tab=importedby)

Opens a new window with list of known importers.

Main Versions Licenses Imports Imported By

## Details

- ![checked](/static/shared/icon/check_circle_gm_grey_24dp.svg) Valid [go.mod](https://github.com/bowman2001/goldmark-supersubscript/tree/v0.1.0/go.mod) file ![](/static/shared/icon/help_gm_grey_24dp.svg)

  The Go module system was introduced in Go 1.11 and is the official dependency management solution for Go.
- ![checked](/static/shared/icon/check_circle_gm_grey_24dp.svg) Redistributable license ![](/static/shared/icon/help_gm_grey_24dp.svg)

  Redistributable licenses place minimal restrictions on how software can be used, modified, and redistributed.
- ![checked](/static/shared/icon/check_circle_gm_grey_24dp.svg) Tagged version ![](/static/shared/icon/help_gm_grey_24dp.svg)

  Modules with tagged versions give importers more predictable builds.
- ![unchecked](/static/shared/icon/cancel_gm_grey_24dp.svg) Stable version ![](/static/shared/icon/help_gm_grey_24dp.svg)

  When a project reaches major version v1 it is considered stable.
- [Learn more about best practices](/about#best-practices)

## Repository

[github.com/bowman2001/goldmark-supersubscript](https://github.com/bowman2001/goldmark-supersubscript "https://github.com/bowman2001/goldmark-supersubscript")

## Links

- [![Open Source Insights Logo](/static/shared/icon/depsdev-logo.svg) Open Source Insights](https://deps.dev/go/github.com%2Fbowman2001%2Fgoldmark-supersubscript/v0.1.0 "View this module on Open Source Insights")

Jump to ...

- [extension](#extension)
  - [Details](#details)
  - [Repository](#repository)
  - [Links](#links)
  - [ README ¶](#-readme-)
    - [goldmark-SuperSubscript](#goldmark-supersubscript)
      - [Syntax](#syntax)
  - [ Documentation ¶](#-documentation-)
    - [Index ¶](#index-)
    - [Constants ¶](#constants-)
    - [Variables ¶](#variables-)
    - [Functions ¶](#functions-)
      - [func NewSubscriptHTMLRenderer ¶](#func-newsubscripthtmlrenderer-)
      - [func NewSubscriptParser ¶](#func-newsubscriptparser-)
      - [func NewSuperscriptHTMLRenderer ¶](#func-newsuperscripthtmlrenderer-)
      - [func NewSuperscriptParser ¶](#func-newsuperscriptparser-)
    - [Types ¶](#types-)
      - [type SubscriptHTMLRenderer ¶](#type-subscripthtmlrenderer-)
      - [func (\*SubscriptHTMLRenderer) RegisterFuncs ¶](#func-subscripthtmlrenderer-registerfuncs-)
      - [type SuperscriptHTMLRenderer ¶](#type-superscripthtmlrenderer-)
      - [func (\*SuperscriptHTMLRenderer) RegisterFuncs ¶](#func-superscripthtmlrenderer-registerfuncs-)
  - [ Source Files ¶](#-source-files-)
  - [ Directories ¶](#-directories-)
  - [Jump to](#jump-to)
  - [Keyboard shortcuts](#keyboard-shortcuts)

README

## ![](/static/shared/icon/chrome_reader_mode_gm_grey_24dp.svg) README [¶](#section-readme "Go to Readme")

[![Documentation](https://pkg.go.dev/badge/github.com/bowman2001/goldmark-supersubscript.svg)](https://pkg.go.dev/github.com/bowman2001/goldmark-supersubscript) [![Test Status](https://github.com/bowman2001/goldmark-supersubscript/workflows/test/badge.svg)](https://github.com/bowman2001/goldmark-supersubscript/actions?query=workflow%3Atest) [![Coverage Status](https://coveralls.io/repos/github/bowman2001/goldmark-supersubscript/badge.svg)](https://coveralls.io/github/bowman2001/goldmark-supersubscript) [![Report Card](https://goreportcard.com/badge/github.com/bowman2001/goldmark-supersubscript)](https://goreportcard.com/report/github.com/bowman2001/goldmark-supersubscript)

### goldmark-SuperSubscript

This Go module contains two extensions for the Markdown parser [Goldmark](https://github.com/yuin/goldmark) providing super- and subscripts.

#### Syntax

Similar to [markdown-it](https://github.com/markdown-it/markdown-it) the new markup characters are:

- The circumflex `^` for superscript
- The tilde `~` for subscript

We need to place one before and one after the text segment like `H~2~O` or `x^2^`.

**No whitespace** between the two surrounding markup characters is allowed. This way the common slip using TeX syntax like `x^2 + x^5` does not lead to messed up HTML. In case we definitely want to insert space we need to place a non-breaking space—either directly or as the HTML entity `&nbsp;`.

Expand ▾ Collapse ▴

## ![](/static/shared/icon/code_gm_grey_24dp.svg) Documentation [¶](#section-documentation "Go to Documentation")

### Index [¶](#pkg-index "Go to Index")

- [extension](#extension)
  - [Details](#details)
  - [Repository](#repository)
  - [Links](#links)
  - [ README ¶](#-readme-)
    - [goldmark-SuperSubscript](#goldmark-supersubscript)
      - [Syntax](#syntax)
  - [ Documentation ¶](#-documentation-)
    - [Index ¶](#index-)
    - [Constants ¶](#constants-)
    - [Variables ¶](#variables-)
    - [Functions ¶](#functions-)
      - [func NewSubscriptHTMLRenderer ¶](#func-newsubscripthtmlrenderer-)
      - [func NewSubscriptParser ¶](#func-newsubscriptparser-)
      - [func NewSuperscriptHTMLRenderer ¶](#func-newsuperscripthtmlrenderer-)
      - [func NewSuperscriptParser ¶](#func-newsuperscriptparser-)
    - [Types ¶](#types-)
      - [type SubscriptHTMLRenderer ¶](#type-subscripthtmlrenderer-)
      - [func (\*SubscriptHTMLRenderer) RegisterFuncs ¶](#func-subscripthtmlrenderer-registerfuncs-)
      - [type SuperscriptHTMLRenderer ¶](#type-superscripthtmlrenderer-)
      - [func (\*SuperscriptHTMLRenderer) RegisterFuncs ¶](#func-superscripthtmlrenderer-registerfuncs-)
  - [ Source Files ¶](#-source-files-)
  - [ Directories ¶](#-directories-)
  - [Jump to](#jump-to)
  - [Keyboard shortcuts](#keyboard-shortcuts)

### Constants [¶](#pkg-constants "Go to Constants")

This section is empty.

### Variables [¶](#pkg-variables "Go to Variables")

[View Source](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go#L119)

```
var Subscript = &subscript{}
```

Subscript is an extension that allows you to use a subscript expression like 'x~0~'.

[View Source](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go#L97)

```
var SubscriptAttributeFilter = html.GlobalAttributeFilter
```

SubscriptAttributeFilter defines attribute names which dd elements can have.

[View Source](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go#L119)

```
var Superscript = &superscript{}
```

Superscript is an extension that allows you to use a superscript expression like 'x^2^'.

[View Source](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go#L97)

```
var SuperscriptAttributeFilter = html.GlobalAttributeFilter
```

SuperscriptAttributeFilter defines attribute names which dd elements can have.

### Functions [¶](#pkg-functions "Go to Functions")

#### func [NewSubscriptHTMLRenderer](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go#L81) [¶](#NewSubscriptHTMLRenderer "Go to NewSubscriptHTMLRenderer")

```
func NewSubscriptHTMLRenderer(opts ...html.Option) renderer.NodeRenderer
```

NewSubscriptHTMLRenderer returns a new SubscriptHTMLRenderer.

#### func [NewSubscriptParser](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go#L38) [¶](#NewSubscriptParser "Go to NewSubscriptParser")

```
func NewSubscriptParser() parser.InlineParser
```

NewSubscriptParser return a new InlineParser that parses subscript expressions.

#### func [NewSuperscriptHTMLRenderer](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go#L81) [¶](#NewSuperscriptHTMLRenderer "Go to NewSuperscriptHTMLRenderer")

```
func NewSuperscriptHTMLRenderer(opts ...html.Option) renderer.NodeRenderer
```

NewSuperscriptHTMLRenderer returns a new SuperscriptHTMLRenderer.

#### func [NewSuperscriptParser](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go#L38) [¶](#NewSuperscriptParser "Go to NewSuperscriptParser")

```
func NewSuperscriptParser() parser.InlineParser
```

NewSuperscriptParser return a new InlineParser that parses superscript expressions.

### Types [¶](#pkg-types "Go to Types")

#### type [SubscriptHTMLRenderer](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go#L76) [¶](#SubscriptHTMLRenderer "Go to SubscriptHTMLRenderer")

```
type SubscriptHTMLRenderer struct {
	html.Config
}
```

SubscriptHTMLRenderer is a renderer.NodeRenderer implementation that renders Subscript nodes.

#### func (\*SubscriptHTMLRenderer) [RegisterFuncs](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go#L92) [¶](#SubscriptHTMLRenderer.RegisterFuncs "Go to SubscriptHTMLRenderer.RegisterFuncs")

```
func (r *SubscriptHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer)
```

RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.

#### type [SuperscriptHTMLRenderer](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go#L76) [¶](#SuperscriptHTMLRenderer "Go to SuperscriptHTMLRenderer")

```
type SuperscriptHTMLRenderer struct {
	html.Config
}
```

SuperscriptHTMLRenderer is a renderer.NodeRenderer implementation that renders Superscript nodes.

#### func (\*SuperscriptHTMLRenderer) [RegisterFuncs](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go#L92) [¶](#SuperscriptHTMLRenderer.RegisterFuncs "Go to SuperscriptHTMLRenderer.RegisterFuncs")

```
func (r *SuperscriptHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer)
```

RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.

## ![](/static/shared/icon/insert_drive_file_gm_grey_24dp.svg) Source Files [¶](#section-sourcefiles "Go to Source Files")

[View all Source files](https://github.com/bowman2001/goldmark-supersubscript/tree/v0.1.0)

- [subscript.go](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/subscript.go "subscript.go")
- [superscript.go](https://github.com/bowman2001/goldmark-supersubscript/blob/v0.1.0/superscript.go "superscript.go")

## ![](/static/shared/icon/folder_gm_grey_24dp.svg) Directories [¶](#section-directories "Go to Directories")

Show internal Expand all

Path Synopsis

[ast](/github.com/bowman2001/goldmark-supersubscript@v0.1.0/ast)

Package ast defines AST nodes that represents extension's elements

Package ast defines AST nodes that represents extension's elements

Click to show internal directories.

Click to hide internal directories.

[Why Go](https://go.dev/solutions) [Use Cases](https://go.dev/solutions#use-cases) [Case Studies](https://go.dev/solutions#case-studies)

[Get Started](https://learn.go.dev/) [Playground](https://play.golang.org) [Tour](https://tour.golang.org) [Stack Overflow](https://stackoverflow.com/questions/tagged/go?tab=Newest) [Help](https://go.dev/help)

[Packages](https://pkg.go.dev) [Standard Library](/std) [Sub-repositories](/golang.org/x) [About Go Packages](https://pkg.go.dev/about)

[About](https://go.dev/project) [Download](https://go.dev/dl/) [Blog](https://go.dev/blog) [Issue Tracker](https://github.com/golang/go/issues) [Release Notes](https://go.dev/doc/devel/release.html) [Brand Guidelines](https://blog.golang.org/go-brand) [Code of Conduct](https://go.dev/conduct)

[Connect](https://www.twitter.com/golang) [Twitter](https://www.twitter.com/golang) [GitHub](https://github.com/golang) [Slack](https://invite.slack.golangbridge.org/) [r/golang](https://reddit.com/r/golang) [Meetup](https://www.meetup.com/pro/go) [Golang Weekly](https://golangweekly.com/)

![Gopher in flight goggles](/static/shared/gopher/pilot-bust-1431x901.svg)

- [Copyright](https://go.dev/copyright)
- [Terms of Service](https://go.dev/tos)
- [Privacy Policy](http://www.google.com/intl/en/policies/privacy/)
- [Report an Issue](https://go.dev/s/pkgsite-feedback)
- ![System theme](/static/shared/icon/brightness_6_gm_grey_24dp.svg) ![Dark theme](/static/shared/icon/brightness_2_gm_grey_24dp.svg) ![Light theme](/static/shared/icon/light_mode_gm_grey_24dp.svg)

  Theme Toggle
- ![](/static/shared/icon/keyboard_grey_24dp.svg)

  Shortcuts Modal

[![Google logo](/static/shared/logo/google-white.svg)](https://google.com)

## Jump to

![](/static/shared/icon/close_gm_grey_24dp.svg)

Close

## Keyboard shortcuts

![](/static/shared/icon/close_gm_grey_24dp.svg)

| **?**          | : This menu     |
|----------------|-----------------|
| **/**          | : Search site   |
| **f** or **F** | : Jump to       |
| **y** or **Y** | : Canonical URL |

Close

go.dev uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic. [Learn more.](https://policies.google.com/technologies/cookies)

Okay
