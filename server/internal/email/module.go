package email

import (
	"fmt"
	"time"

	"github.com/matcornic/hermes/v2"
	"gopkg.in/gomail.v2"
)

type Service interface {
	SendVerifyUserEmail(email string, name string, token string) error
}

func New(conf *Config) Service {
	d := gomail.NewDialer(conf.SMTP.Host, conf.SMTP.Port, conf.SMTP.User, conf.SMTP.Password)
	template := hermes.Hermes{
		Theme: &hermes.Flat{},
		Product: hermes.Product{
			Name:      conf.Product.Name,
			Link:      conf.Product.Link,
			Logo:      conf.Product.Logo,
			Copyright: fmt.Sprintf("%d codeserk.es. All rights reserved.", time.Now().Year()),
		},
	}

	return &service{
		conf:     conf,
		dialer:   d,
		template: &template,
	}
}
