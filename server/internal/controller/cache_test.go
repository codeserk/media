package controller

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheKey(t *testing.T) {
	t.Run("should create cache key from URL with single query param", func(t *testing.T) {
		// Arrange
		req, _ := http.NewRequest("GET", "/test?param=value", nil)

		// Act
		key := CacheKey(req)

		// Assert
		assert.Equal(t, "/test?param=value", key)
	})

	t.Run("should create cache key from URL with multiple query params in sorted order", func(t *testing.T) {
		// Arrange
		url, _ := url.Parse("/test?c=3&a=1&b=2")
		req := &http.Request{URL: url}

		// Act
		key := CacheKey(req)

		// Assert
		assert.Equal(t, "/test?a=1&b=2&c=3", key)
	})

	t.Run("should create cache key from URL with no query params", func(t *testing.T) {
		// Arrange
		req, _ := http.NewRequest("GET", "/test", nil)

		// Act
		key := CacheKey(req)

		// Assert
		assert.Equal(t, "/test?", key)
	})

	t.Run("should create cache key from URL with duplicate query params", func(t *testing.T) {
		// Arrange
		url, _ := url.Parse("/test?param=1&other=2&param=3")
		req := &http.Request{URL: url}

		// Act
		key := CacheKey(req)

		// Assert
		assert.Equal(t, "/test?other=2&param=1&param=3", key)
	})
}
