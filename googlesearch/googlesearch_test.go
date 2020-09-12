package googlesearch_test

import (
	"fmt"
	"github.com/sundowndev/dorkgen/googlesearch"
	"net/url"
	"testing"

	assertion "github.com/stretchr/testify/assert"
)

var dork *googlesearch.GoogleSearch

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should convert to URL correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("example.com").
			URL()

		assert.Equal(result, "https://www.google.com/search?q=site%3Aexample.com", "they should be equal")
	})

	t.Run("should convert to string correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := fmt.Sprint(dork.Site("example.com"))

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle site tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("example.com").
			String()

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle intext tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InText("text").
			String()

		assert.Equal(result, "intext:\"text\"", "they should be equal")
	})

	t.Run("should handle inurl tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InURL("index.php").
			String()

		assert.Equal(result, "inurl:\"index.php\"", "they should be equal")
	})

	t.Run("should handle filetype tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			FileType("pdf").
			String()

		assert.Equal(result, "filetype:\"pdf\"", "they should be equal")
	})

	t.Run("should handle cache tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Cache("www.google.com").
			String()

		assert.Equal(result, "cache:\"www.google.com\"", "they should be equal")
	})

	t.Run("should handle related tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Related("www.google.com").
			String()

		assert.Equal(result, "related:\"www.google.com\"", "they should be equal")
	})

	t.Run("should handle ext tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Ext("(doc | pdf | xls | txt | xml)").
			String()

		assert.Equal(result, "ext:(doc | pdf | xls | txt | xml)", "they should be equal")
	})

	t.Run("should handle exclude tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Exclude(googlesearch.New().Plain("html")).
			Exclude(googlesearch.New().Plain("htm")).
			Exclude(googlesearch.New().Plain("php")).
			Exclude(googlesearch.New().Plain("md5sums")).
			String()

		assert.Equal(result, "-html -htm -php -md5sums", "they should be equal")
	})

	t.Run("should handle 'OR' tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("facebook.com").
			Or().
			Site("twitter.com").
			String()

		assert.Equal(result, "site:facebook.com | site:twitter.com", "they should be equal")
	})

	t.Run("should handle 'AND' tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InTitle("facebook").
			And().
			InTitle("twitter").
			String()

		assert.Equal(result, "intitle:\"facebook\" + intitle:\"twitter\"", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("linkedin.com").
			Group(googlesearch.New().InText("1").Or().InText("2")).
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" | intext:\"2\")", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("linkedin.com").
			Group(googlesearch.New().InText("1").Or().InText("2")).
			InTitle("jordan").
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" | intext:\"2\") intitle:\"jordan\"", "they should be equal")
	})

	t.Run("should return URL values", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("linkedin.com").
			Group(googlesearch.New().InText("1").Or().InText("2")).
			InTitle("jordan").
			QueryValues()

		assert.Equal(url.Values{
			"q": []string{"site:linkedin.com (intext:\"1\" | intext:\"2\") intitle:\"jordan\""},
		}, result, "they should be equal")
	})

	t.Run("should use book tag", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Book("test").
			String()

		assert.Equal("book:\"test\"", result, "they should be equal")
	})

	t.Run("should use maps tag", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Maps("france").
			String()

		assert.Equal("maps:france", result, "they should be equal")
	})

	t.Run("should use allintext tag", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			AllInText("test").
			String()

		assert.Equal("allintext:\"test\"", result, "they should be equal")
	})

	t.Run("should use info tag", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Info("https://google.com/").
			String()

		assert.Equal("info:\"https://google.com/\"", result, "they should be equal")
	})

	t.Run("should use inanchor tag", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InAnchor("test").
			String()

		assert.Equal("inanchor:\"test\"", result, "they should be equal")
	})
}
