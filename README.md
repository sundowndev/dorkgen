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
  <a href="https://codecov.io/gh/sundowndev/dorkgen">
    <img src="https://codecov.io/gh/sundowndev/dorkgen/branch/master/graph/badge.svg" />
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

[Try it in the Go playground](https://play.golang.org/p/ck_hEoX8cTK)

#### Get started

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork := &dorkgen.GoogleSearch{}
  // dork := &dorkgen.DuckDuckGo{}
  // dork := &dorkgen.Bing{}

  dork.Site("example.com").Intext("text").ToString()
  // returns: site:example.com intext:"text"
}
```

#### Operators

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork.Site("facebook.com").Or().Site("twitter.com").ToString()
  // returns: site:facebook.com OR site:twitter.com

  dork.Intext("facebook").And().Intext("twitter").ToString()
  // returns: intext:"facebook" AND intext:"twitter"
}
```

#### Exclude results

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork.
    Exclude((&dorkgen.GoogleSearch{}).
      Site("example.com").
      ToString()).
    Site("example.*").
    Or().
    Intext("text")
  // returns: -site:example.com site:example.* OR "text"
}
```

#### Group tags along with operators

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork.
    Group((&dorkgen.GoogleSearch{}).
      Site("facebook.com").
      Or().
      Site("twitter.com").
      ToString()).
    Intext("wtf").
    ToString()
  // returns: (site:facebook.com OR site:twitter.com) "wtf"
}
```

#### URL conversion

```go
package main

import "github.com/sundowndev/dorkgen"

func main() {
  dork.
    Site("facebook.*").
    Exclude((&dorkgen.GoogleSearch{}).
      Site("facebook.com").
      ToString())

  dork.ToString()
  // returns: site:facebook.* -site:facebook.com
  dork.ToURL()
  // returns: https://www.google.com/search?q=site%3Afacebook.%2A+-site%3Afacebook.com
}
```

## Support

[![](docs/jetbrains.svg)](https://www.jetbrains.com/?from=sundowndev)

Thanks to [JetBrains](https://www.jetbrains.com/?from=sundowndev) for supporting my open-source projects.
