package dorkgen

import (
	"github.com/sundowndev/dorkgen/duckduckgo"
	"github.com/sundowndev/dorkgen/googlesearch"
)

func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}

func NewDuckDuckGo() *duckduckgo.DuckDuckGo {
	return duckduckgo.New()
}
