package signature

import (
	"fmt"
	"media/internal/config"
	"net/http"
)

type Service struct {
	conf *config.APISignature
}

type SignatureApp = string

var (
	SignatureAppDashboard SignatureApp = "dashboard"
	SignatureAppConsumer  SignatureApp = "consumer"
)
var validSignatureApps = map[SignatureApp]bool{SignatureAppDashboard: true,
	SignatureAppConsumer: true,
}

var signatureCheckOffsets = []int{0, -1, 1, -2, 2, -3, -4, -5}

var SignatureAppHeaderName = "x-signature-app"
var SignatureHeaderName = "x-signature"

func NewService(conf *config.APISignature) *Service {
	return &Service{conf}
}

func (s *Service) Check(r *http.Request, body string, requiredApp SignatureApp) error {
	if s.conf.Skip {
		return nil
	}

	// check app
	app, err := s.extractSignatureApp(r)
	if err != nil {
		return err
	}
	if app == nil {
		return fmt.Errorf("failed to get app from request")
	}
	if *app != requiredApp {
		return fmt.Errorf("invalid signature app: %v expected %v", &app, requiredApp)
	}

	// check signature
	sig := s.extractSignature(r)
	builder := NewBuilder(s.extractAppSecret(requiredApp))
	builder.WithRequest(r, body)
	for _, offset := range signatureCheckOffsets {
		if sig == builder.Build(nil, offset) {
			return nil
		}
	}

	return fmt.Errorf("invalid signature")
}

func (s *Service) extractSignatureApp(r *http.Request) (*SignatureApp, error) {
	value := r.Header.Get(SignatureAppHeaderName)
	if _, exists := validSignatureApps[SignatureApp(value)]; !exists {
		return nil, fmt.Errorf("invalid signature app: %v", value)
	}

	return &value, nil
}

func (s *Service) extractSignature(r *http.Request) string {
	return r.Header.Get(SignatureHeaderName)
}

func (s *Service) extractAppSecret(app SignatureApp) string {
	switch app {
	case SignatureAppDashboard:
		return s.conf.DashboardSecret
	case SignatureAppConsumer:
		return s.conf.ConsumerSecret
	}
	return s.conf.DashboardSecret
}
