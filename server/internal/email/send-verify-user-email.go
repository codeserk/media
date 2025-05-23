package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
)

func (s *service) SendVerifyUserEmail(email string, name string, url string) error {
	content := hermes.Email{
		Body: hermes.Body{
			Name: name,
			Intros: []string{
				"Please verify your email.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your email",
						Link:  url,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	body, err := s.template.GenerateHTML(content)
	if err != nil {
		return fmt.Errorf("hermes generate: %v", err)
	}

	return s.send(email, "Verify your email", body)
}
