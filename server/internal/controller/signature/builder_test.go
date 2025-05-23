package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	t.Run("should set secret when WithSecret is called", func(t *testing.T) {
		// Arrange
		expected := "test-secret"

		// Act
		builder := NewBuilder(expected)

		// Assert
		assert.Equal(t, expected, builder.secret)
	})

	t.Run("should append string when AddString is called", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		expected := "test-string"

		// Act
		builder.AddString(expected)

		// Assert
		assert.Equal(t, expected, builder.parts[0])
	})

	t.Run("should append object as string when AddObject is called", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		testObj := struct {
			Name string
		}{
			Name: "test",
		}
		expected := "{\n\t\"Name\": \"test\"\n}"

		// Act
		builder.AddObject(testObj)

		// Assert
		assert.Equal(t, expected, builder.parts[0])
	})

	t.Run("should add request components when WithRequest is called", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		req, _ := http.NewRequest("GET", "http://test.com/path?q=1", nil)
		body := "test-body"
		expectedURI := "/path?q=1"
		expectedBody := base64.StdEncoding.EncodeToString([]byte(body))

		// Act
		builder.WithRequest(req, body)

		// Assert
		assert.Equal(t, expectedURI, builder.parts[0])
		assert.Equal(t, expectedBody, builder.parts[1])
	})

	t.Run("should generate HMAC when Build is called", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		builder.AddString("test")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		content := builder.generateContent(&fixedTime, 0)
		h := hmac.New(sha256.New, []byte("secret"))
		h.Write([]byte(content))
		expected := base64.StdEncoding.EncodeToString(h.Sum(nil))

		// Act
		result := builder.Build(&fixedTime, 0)

		// Assert
		assert.Equal(t, expected, result)
	})

	t.Run("should build correctly with multiple strings", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		builder.AddString("test1")
		builder.AddString("test2")
		builder.AddString("test3")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedContent := base64.StdEncoding.EncodeToString([]byte("test1_test2_test3_2023-01-01T00:00:00.000Z"))

		// Act
		actualContent := builder.generateContent(&fixedTime, 0)
		result := builder.Build(&fixedTime, 0)

		// Assert
		assert.Equal(t, expectedContent, actualContent)
		h := hmac.New(sha256.New, []byte("secret"))
		h.Write([]byte(expectedContent))
		expected := base64.StdEncoding.EncodeToString(h.Sum(nil))
		assert.Equal(t, expected, result)
	})

	t.Run("should build correctly with object and string", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		testObj := struct {
			Name string
		}{
			Name: "test",
		}
		builder.AddObject(testObj)
		builder.AddString("additional")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedContent := base64.StdEncoding.EncodeToString([]byte("{\n\t\"Name\": \"test\"\n}_additional_2023-01-01T00:00:00.000Z"))

		// Act
		actualContent := builder.generateContent(&fixedTime, 0)
		result := builder.Build(&fixedTime, 0)

		// Assert
		assert.Equal(t, expectedContent, actualContent)
		h := hmac.New(sha256.New, []byte("secret"))
		h.Write([]byte(expectedContent))
		expected := base64.StdEncoding.EncodeToString(h.Sum(nil))
		assert.Equal(t, expected, result)
	})

	t.Run("should build correctly with request", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		req, _ := http.NewRequest("POST", "http://test.com/path", nil)
		builder.WithRequest(req, "request-body")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedContent := base64.StdEncoding.EncodeToString([]byte("/path_" + base64.StdEncoding.EncodeToString([]byte("request-body")) + "_2023-01-01T00:00:00.000Z"))

		// Act
		actualContent := builder.generateContent(&fixedTime, 0)
		result := builder.Build(&fixedTime, 0)

		// Assert
		assert.Equal(t, expectedContent, actualContent)
		h := hmac.New(sha256.New, []byte("secret"))
		h.Write([]byte(expectedContent))
		expected := base64.StdEncoding.EncodeToString(h.Sum(nil))
		assert.Equal(t, expected, result)
	})

	t.Run("should build correctly with request without body", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		req, _ := http.NewRequest("GET", "http://test.com/path", nil)
		builder.WithRequest(req, "")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedContent := base64.StdEncoding.EncodeToString([]byte("/path_e30=_2023-01-01T00:00:00.000Z"))

		// Act
		actualContent := builder.generateContent(&fixedTime, 0)
		result := builder.Build(&fixedTime, 0)

		// Assert
		assert.Equal(t, expectedContent, actualContent)
		h := hmac.New(sha256.New, []byte("secret"))
		h.Write([]byte(expectedContent))
		expected := base64.StdEncoding.EncodeToString(h.Sum(nil))
		assert.Equal(t, expected, result)
	})

	t.Run("should format time correctly", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expected := "2023-01-01T00:00:00.000Z"

		// Act
		result := builder.getTimeString(&fixedTime, 0)

		// Assert
		assert.Equal(t, expected, result)
	})

	t.Run("should apply offset correctly to time", func(t *testing.T) {
		// Arrange
		builder := NewBuilder("secret")
		fixedTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		expected := "2023-01-01T00:01:00.000Z"

		// Act
		result := builder.getTimeString(&fixedTime, 1)

		// Assert
		assert.Equal(t, expected, result)
	})
}
