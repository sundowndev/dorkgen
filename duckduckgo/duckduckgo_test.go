package duckduckgo_test

import (
	"fmt"
	"github.com/sundowndev/dorkgen/duckduckgo"
	"net/url"
	"testing"

	assertion "github.com/stretchr/testify/assert"
)

var dork *duckduckgo.DuckDuckGo

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should convert to URL correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Site("example.com").
			URL()

		assert.Equal(result, "https://www.google.com/search?q=site%3Aexample.com", "they should be equal")
	})

	t.Run("should convert to string correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := fmt.Sprint(dork.Site("example.com"))

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle site tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Site("example.com").
			String()

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle intext tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Intext("text").
			String()

		assert.Equal(result, "intext:\"text\"", "they should be equal")
	})

	t.Run("should handle inurl tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Inurl("index.php").
			String()

		assert.Equal(result, "inurl:\"index.php\"", "they should be equal")
	})

	t.Run("should handle filetype tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Filetype("pdf").
			String()

		assert.Equal(result, "filetype:\"pdf\"", "they should be equal")
	})

	t.Run("should handle AllInURL tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			AllInURL("test").
			String()

		assert.Equal(result, "allinurl:\"test\"", "they should be equal")
	})

	t.Run("should handle Location tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Location("fr").
			String()

		assert.Equal(result, "region:\"fr\"", "they should be equal")
	})

	t.Run("should handle Feed tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Feed("rss").
			String()

		assert.Equal(result, "feed:rss", "they should be equal")
	})

	t.Run("should handle HasFeed tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			HasFeed("https://sundaypapers.libsyn.com/rss").
			String()

		assert.Equal(result, "hasfeed:\"https://sundaypapers.libsyn.com/rss\"", "they should be equal")
	})

	t.Run("should handle Language tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Language("zh").
			String()

		assert.Equal(result, "language:zh", "they should be equal")
	})

	t.Run("should handle AllInTitle tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			AllInTitle("test").
			String()

		assert.Equal(result, "allintitle:\"test\"", "they should be equal")
	})

	t.Run("should handle ext tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Ext("(doc | pdf | xls | txt | xml)").
			String()

		assert.Equal(result, "ext:(doc | pdf | xls | txt | xml)", "they should be equal")
	})

	t.Run("should handle exclude tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Exclude(duckduckgo.New().Plain("html")).
			Exclude(duckduckgo.New().Plain("htm")).
			Exclude(duckduckgo.New().Plain("php")).
			Exclude(duckduckgo.New().Plain("md5sums")).
			String()

		assert.Equal(result, "-html -htm -php -md5sums", "they should be equal")
	})

	t.Run("should handle 'OR' tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Site("facebook.com").
			Or().
			Site("twitter.com").
			String()

		assert.Equal(result, "site:facebook.com | site:twitter.com", "they should be equal")
	})

	t.Run("should handle 'AND' tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Intitle("facebook").
			And().
			Intitle("twitter").
			String()

		assert.Equal(result, "intitle:\"facebook\" + intitle:\"twitter\"", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Site("linkedin.com").
			Group(duckduckgo.New().Intext("1").Or().Intext("2")).
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" | intext:\"2\")", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Site("linkedin.com").
			Group(duckduckgo.New().Intext("1").Or().Intext("2")).
			Intitle("jordan").
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" | intext:\"2\") intitle:\"jordan\"", "they should be equal")
	})

	t.Run("should return URL values", func(t *testing.T) {
		dork = duckduckgo.New()

		result := dork.
			Site("linkedin.com").
			Group(duckduckgo.New().Intext("1").Or().Intext("2")).
			Intitle("jordan").
			QueryValues()

		assert.Equal(url.Values{
			"q": []string{"site:linkedin.com (intext:\"1\" | intext:\"2\") intitle:\"jordan\""},
		}, result, "they should be equal")
	})
}
