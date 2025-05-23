package email

import (
	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type service struct {
	conf     *Config
	dialer   *gomail.Dialer
	template *hermes.Hermes
}
