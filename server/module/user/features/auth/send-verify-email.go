package auth

import (
	"fmt"
	"media/module/user"
	"net/url"
)

func (s *service) SendVerifyEmail(u *user.Entity) error {
	updatedUser, err := s.repository.UpdateEmailVerify(u)
	if err != nil {
		return fmt.Errorf("update email verify: %v", err)
	}

	link := fmt.Sprintf("%s/auth/verify?e=%s&t=%s", s.dashboardConfig.BaseURL, url.QueryEscape(updatedUser.Email), url.QueryEscape(updatedUser.EmailVerify.Token))

	return s.emails.SendVerifyUserEmail(u.Email, u.Name, link)
}
