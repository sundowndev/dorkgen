<div align="left">
  <a href="https://godoc.org/github.com/sundowndev/dorkgen">
    <img src="https://godoc.org/github.com/sundowndev/dorkgen?status.svg" alt="GoDoc">
  </a>
  <a href="https://github.com/sundowndev/dorkgen/actions">
    <img src="https://img.shields.io/endpoint.svg?url=https://actions-badge.atrox.dev/sundowndev/dorkgen/badge?ref=master" alt="build status" />
  </a>
  <a href="https://goreportcard.com/report/github.com/sundowndev/dorkgen">
    <img src="https://goreportcard.com/badge/github.com/sundowndev/dorkgen" alt="go report" />
  </a>
  <a href="https://codeclimate.com/github/sundowndev/dorkgen/maintainability">
    <img src="https://api.codeclimate.com/v1/badges/e827d7cc994c6519d319/maintainability" />
  </a>
  <a href="https://codecov.io/gh/sundowndev/dorkgen">
    <img src="https://codecov.io/gh/sundowndev/dorkgen/branch/master/graph/badge.svg" />
  </a>
  <a href="https://github.com/sundowndev/dorkgen/releases">
    <img src="https://img.shields.io/github/release/SundownDEV/dorkgen.svg" alt="Latest version" />
  </a>
</div>

# Dorkgen

Dorkgen is a dork request generator for popular search engines such as Google Search, DuckDuckGo, Yahoo and Bing. [Learn more about Google Hacking](https://en.wikipedia.org/wiki/Google_hacking). The goal of this package is to provide simple interfaces to creates valid dork queries for various search engines.

## Current status

Version 1 is ongoing, but the API is still unstable. **See below table for supported search engines**.

| Search engine        | Implementation status |
|---------------|:-----------------------:|
| [Google Search](https://godoc.org/github.com/sundowndev/dorkgen/googlesearch) | Partial      |
| [DuckDuckGo](https://godoc.org/github.com/sundowndev/dorkgen/duckduckgo)    | Complete              |
| Yahoo Search  | WIP                   |
| Bing Search   | WIP                   |

## Install

Fetch the module :

```bash
go get github.com/sundowndev/dorkgen
```

## Usage

**[Try it in the Go playground](https://play.golang.org/p/H_sCs_CB10A)**

#### Get started

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork := dorkgen.NewGoogleSearch()
  // dork := dorkgen.NewDuckDuckGo()
  // dork := dorkgen.NewBingSearch()
  // dork := dorkgen.NewYahooSearch()

  dork.Site("example.com").InText("text").String()
  // returns: site:example.com intext:"text"
}
```

#### Operators

```go
func main() {
  dork.Site("facebook.com").Or().Site("twitter.com").String()
  // returns: site:facebook.com | site:twitter.com

  dork.InText("facebook").And().InText("twitter").String()
  // returns: intext:"facebook" + intext:"twitter"
}
```

#### Exclude results

```go
func main() {
  dork.
    Exclude((dorkgen.NewGoogleSearch()).
      Site("example.com").
      String()).
    Site("example.*").
    Or().
    InText("text")
  // returns: -site:example.com site:example.* | "text"
}
```

#### Group tags along with operators

```go
func main() {
  dork.
    Group((dorkgen.NewGoogleSearch()).
      Site("facebook.com").
      Or().
      Site("twitter.com")).
    InText("wtf").
    String()
  // returns: (site:facebook.com | site:twitter.com) "wtf"
}
```

#### URL conversion

```go
func main() {
  dork.
    Site("facebook.*").
    Exclude((dorkgen.NewGoogleSearch()).
      Site("facebook.com").
      String())

  dork.String()
  // returns: site:facebook.* -site:facebook.com
  dork.ToURL()
  // returns: https://www.google.com/search?q=site%3Afacebook.%2A+-site%3Afacebook.com
}
```

## Support

[![](docs/jetbrains.svg)](https://www.jetbrains.com/?from=sundowndev)

Thanks to [JetBrains](https://www.jetbrains.com/?from=sundowndev) for supporting my open-source projects.
