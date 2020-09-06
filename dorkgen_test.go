package dorkgen

import (
	assertion "github.com/stretchr/testify/assert"
	"github.com/sundowndev/dorkgen/duckduckgo"
	"github.com/sundowndev/dorkgen/googlesearch"
	"testing"
)

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should create a GoogleSearch instance", func(t *testing.T) {
		dork := NewGoogleSearch()

		assert.IsType(&googlesearch.GoogleSearch{}, dork, "they should be equal")
	})

	t.Run("should create a DuckDuckGo instance", func(t *testing.T) {
		dork := NewDuckDuckGo()

		assert.IsType(&duckduckgo.DuckDuckGo{}, dork, "they should be equal")
	})
}
