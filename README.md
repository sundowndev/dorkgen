# Dorkgen

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
  <a href="https://github.com/sundowndev/dorkgen/releases">
    <img src="https://img.shields.io/github/release/SundownDEV/dorkgen.svg" alt="Latest version" />
  </a>
</div>

Dorkgen is a dork request generator for popular search engines such as Google, DuckDuckGo and Bing. [Learn more about Google Hacking](https://en.wikipedia.org/wiki/Google_hacking).

## Install

```bash
go get github.com/sundowndev/dorkgen
```

## Usage

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork := &dorkgen.GoogleSearch{}
  // dork := &dorkgen.DuckDuckGo{}
  // dork := &dorkgen.Bing{}

  dork.Site("example.com").Intext("06792489265").ToString()
  // returns: site:example.com "06792489265"

  dork.Site("example.com").Or().Intext("06792489265").ToString()
  // returns: site:example.com OR "06792489265"

  dork.Site("facebook.*").Exclude("site:facebook.com").ToURL()
  // returns: https://www.google.com/search?q=site%3A"facebook.*"+-site%3Afacebook.com
}
```

## API

```go
type EngineFactory interface {
  Site(string) *GoogleSearch
  ToString() string
  ToURL() string
  Intext(string) *GoogleSearch
  Inurl(string) *GoogleSearch
  Filetype(string) *GoogleSearch
  Cache(string) *GoogleSearch
  Related(string) *GoogleSearch
  Ext(string) *GoogleSearch
  Exclude(string) *GoogleSearch
}
```
