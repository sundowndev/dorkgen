package dorkgen

import (
	"github.com/sundowndev/dorkgen/duckduckgo"
	"github.com/sundowndev/dorkgen/googlesearch"
)

// NewGoogleSearch returns a new instance of GoogleSearch
func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}

// NewDuckDuckGo returns a new instance of DuckDuckGo
func NewDuckDuckGo() *duckduckgo.DuckDuckGo {
	return duckduckgo.New()
}
