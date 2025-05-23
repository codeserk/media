package auth

import (
	"media/internal/errors"
	"media/internal/test/fixture"
	"media/module/user"
	"media/module/user/features/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// setup
	conf := fixture.DashboardConfig()
	db := fixture.Mongo(t)
	cache := fixture.Cache(t)
	jwt := fixture.JWT()
	emails := fixture.Email(t)
	repository := repository.New(db)
	service := New(conf, repository, cache, jwt, emails)

	beforeEach := func() {
		fixture.ClearMongo(db)
		fixture.ClearCache(cache)
	}

	t.Run("should fail if the user does not exist", func(t *testing.T) {
		// Arrange
		beforeEach()

		// Act
		result, err := service.Login(user.LoginParams{Email: "notfound@mail.com", Password: "valid"})

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.(*errors.PressError).Error(), "user not found")
		assert.Nil(t, result)
	})

	t.Run("should fail if the user has wrong password", func(t *testing.T) {
		// Arrange
		beforeEach()
		service.Register(user.RegisterParams{
			Name:     "Name",
			Email:    "user@mail.com",
			Password: "password",
		})

		// Act
		result, err := service.Login(user.LoginParams{Email: "user@mail.com", Password: "invalid"})

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.(*errors.PressError).Error(), "invalid credentials")
		assert.Nil(t, result)
	})

	t.Run("should be able to login with good credentials", func(t *testing.T) {
		// Arrange
		beforeEach()
		service.Register(user.RegisterParams{
			Name:     "Name",
			Email:    "user@mail.com",
			Password: "password",
		})

		// Act
		result, err := service.Login(user.LoginParams{Email: "user@mail.com", Password: "password"})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Name", result.User.Name)
		assert.Equal(t, "user@mail.com", result.User.Email)
		assert.NotEmpty(t, result.Token)
	})
}
