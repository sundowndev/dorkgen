package dorkgen

import "github.com/sundowndev/dorkgen/googlesearch"

func NewGoogleSearch() *googlesearch.GoogleSearch {
	return googlesearch.New()
}
