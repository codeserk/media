package fixture

import (
	"media/internal/email"
	mocks "media/internal/email/mock"
	"testing"
)

func Email(t *testing.T) email.Service {
	return mocks.NewService(t)
}
