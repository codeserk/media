package jwt

import (
	"media/internal/config"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	service := New(config.JWT{Secret: "secret", Issuer: "issuer"})

	t.Run("should generate a valid JWT token", func(t *testing.T) {
		// Arrange
		id := "user123"
		claimType := ClaimTypeUser
		duration := time.Hour

		// Act
		token, err := service.GenerateJWT(id, claimType, duration)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("should generate a valid JWT token", func(t *testing.T) {
		// Arrange
		id := "user123"
		claimType := ClaimTypeUser
		duration := time.Hour
		token, _ := service.GenerateJWT(id, claimType, duration)

		// Act
		claims, err := service.GetClaimsFromToken(token)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, claims.Id, id)
	})

}
