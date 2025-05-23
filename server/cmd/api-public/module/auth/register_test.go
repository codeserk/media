package auth

import (
	"encoding/json"
	"media/internal/test/fixture"
	"media/module/user/features/auth"
	"media/module/user/features/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterEndpoint(t *testing.T) {
	// setup
	conf := fixture.DashboardConfig()
	db := fixture.Mongo(t)
	cache := fixture.Cache(t)
	email := fixture.Email(t)
	jwt := fixture.JWT()
	repository := repository.New(db)
	authService := auth.New(conf, repository, cache, jwt, email)
	handler := register(authService)

	beforeEach := func() {
		fixture.ClearMongo(db)
		fixture.ClearCache(cache)
	}

	t.Run("should fail if name is missing", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"email": "test@example.com", "password": "password"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, res.Code, http.StatusUnprocessableEntity)
		assert.Contains(t, res.Body.String(), "'Name' failed on the 'required' tag")
	})

	t.Run("should fail if name is missing", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"email": "test@example.com", "password": "password"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "'Name' failed on the 'required' tag")
	})

	t.Run("should fail if email is missing", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"name": "test", "password": "password"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "'Email' failed on the 'required' tag")
	})

	t.Run("should fail if email is invalid", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"name": "test", "email": "invalid-email", "password": "password"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "'Email' failed on the 'email' tag")
	})

	t.Run("should fail if password is missing", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"name": "test", "email": "test@example.com"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "'Password' failed on the 'required' tag")
	})

	t.Run("should fail if password is too short", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"name": "test", "email": "test@example.com", "password": "pass"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Contains(t, res.Body.String(), "'Password' failed on the 'min' tag")
	})

	t.Run("should succeed with valid input", func(t *testing.T) {
		// Arrange
		beforeEach()
		reqBody := `{"name": "test", "email": "test@example.com", "password": "password"}`
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(reqBody))
		res := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusOK, res.Code)
		var resJson loginResponse
		err := json.Unmarshal(res.Body.Bytes(), &resJson)
		assert.NoError(t, err)
		assert.NotEmpty(t, resJson.User.Id)
		assert.Equal(t, "test", resJson.User.Name)
		assert.Equal(t, "test@example.com", resJson.User.Email)
		assert.NotEmpty(t, resJson.Token)
	})
}
