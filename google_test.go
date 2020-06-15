package dorkgen

import (
	"fmt"
	"testing"

	assertion "github.com/stretchr/testify/assert"
)

var dork *GoogleSearch

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should convert to URL correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Site("example.com").
			ToURL()

		assert.Equal(result, "https://www.google.com/search?q=site%3Aexample.com", "they should be equal")
	})

	t.Run("should convert to string correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := fmt.Sprint(dork.Site("example.com"))

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle site tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Site("example.com").
			String()

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle intext tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Intext("text").
			String()

		assert.Equal(result, "intext:\"text\"", "they should be equal")
	})

	t.Run("should handle inurl tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Inurl("index.php").
			String()

		assert.Equal(result, "inurl:\"index.php\"", "they should be equal")
	})

	t.Run("should handle filetype tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Filetype("pdf").
			String()

		assert.Equal(result, "filetype:\"pdf\"", "they should be equal")
	})

	t.Run("should handle cache tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Cache("www.google.com").
			String()

		assert.Equal(result, "cache:\"www.google.com\"", "they should be equal")
	})

	t.Run("should handle related tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Related("www.google.com").
			String()

		assert.Equal(result, "related:\"www.google.com\"", "they should be equal")
	})

	t.Run("should handle ext tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Ext("(doc | pdf | xls | txt | xml)").
			String()

		assert.Equal(result, "ext:(doc | pdf | xls | txt | xml)", "they should be equal")
	})

	t.Run("should handle exclude tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Exclude("html").
			Exclude("htm").
			Exclude("php").
			Exclude("md5sums").
			String()

		assert.Equal(result, "-html -htm -php -md5sums", "they should be equal")
	})

	t.Run("should handle or tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Site("facebook.com").
			Or().
			Site("twitter.com").
			String()

		assert.Equal(result, "site:facebook.com OR site:twitter.com", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Site("linkedin.com").
			Group((&GoogleSearch{}).Intext("1").Or().Intext("2")).
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" OR intext:\"2\")", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = &GoogleSearch{}

		result := dork.
			Site("linkedin.com").
			Group((&GoogleSearch{}).Intext("1").Or().Intext("2")).
			Intitle("jordan").
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" OR intext:\"2\") intitle:\"jordan\"", "they should be equal")
	})
}
