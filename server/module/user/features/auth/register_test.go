package auth

import (
	"media/internal/errors"
	"media/internal/test/fixture"
	fixtures "media/internal/test/fixture"
	"media/module/user"
	"media/module/user/features/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// setup
	conf := fixture.DashboardConfig()
	db := fixtures.Mongo(t)
	cache := fixtures.Cache(t)
	jwt := fixtures.JWT()
	emails := fixture.Email(t)
	repository := repository.New(db)
	service := New(conf, repository, cache, jwt, emails)

	beforeEach := func() {
		fixtures.ClearMongo(db)
		fixtures.ClearCache(cache)
	}

	t.Run("should register a new user", func(t *testing.T) {
		// Arrange
		beforeEach()

		// Act
		result, err := service.Register(user.RegisterParams{
			Name:     "Name",
			Email:    "user@mail.com",
			Password: "password",
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Name", result.User.Name)
		assert.Equal(t, "user@mail.com", result.User.Email)
		assert.NotEmpty(t, result.Token)
	})

	t.Run("should fail if another user with the same email is registered", func(t *testing.T) {
		// Arrange
		beforeEach()
		service.Register(user.RegisterParams{
			Name:     "Name",
			Email:    "user@mail.com",
			Password: "password",
		})

		// Act
		result, err := service.Register(user.RegisterParams{
			Name:     "Name",
			Email:    "user@mail.com",
			Password: "password",
		})

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.(*errors.PressError).Public(), "is already used")
		assert.Nil(t, result)
	})
}
