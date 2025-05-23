package logger

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

func Middleware() func(next http.Handler) http.Handler {
	log := zerolog.New(Output()).With().
		Timestamp().
		Str("api", "stats").
		Str("host", "localhost").
		Logger()

	c := alice.New()
	c = c.Append(hlog.NewHandler(log))
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Str("ip", readUserIP(r)).
			Int("size", size).
			Dur("duration", duration).
			Msg("request")
	}))
	c = c.Append(hlog.RequestIDHandler("reqId", "Request-Id"))

	return func(next http.Handler) http.Handler {
		return c.Then(next)
	}
}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return strings.Split(r.RemoteAddr, ":")[0]
		}
		IPAddress = ip
	}

	return IPAddress
}
