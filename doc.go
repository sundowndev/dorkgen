/*
Package dorkgen is a Go package to generate dork requests for popular search engines such as Google, DuckDuckGo and Bing.
It allows you to define requests programmatically and convert them into string.
You can use it as following:

package main

import "github.com/sundowndev/dorkgen"

func main() {
	dork := dorkgen.NewGoogleSearch()
	// dork := dorkgen.NewDuckDuckGo()
	// dork := dorkgen.NewBing()

	dork.Site("example.com").Intext("text").String()
	// returns: site:example.com "text"
}

// You can also isolate tags between parentheses
func main() {
	dork := dorkgen.NewGoogleSearch()

	dork.Group(
		dorkgen.NewGoogleSearch().
			Site("facebook.com")).
	Or().
	Group(
		dorkgen.NewGoogleSearch().
			Site("twitter.com")).
	Intext("text")
	String()
	// returns: (site:facebook.com) OR (site:twitter.com) "text"
}
*/
package dorkgen
