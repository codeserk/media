package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func (s *service) send(to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", s.conf.FromEmail, s.conf.FromName)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	err := s.dialer.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("dial and send: %v", err)
	}

	return nil
}
