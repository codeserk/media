package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Builder struct {
	secret string
	parts  []string
}

func (b *Builder) AddString(input string) *Builder {
	b.parts = append(b.parts, input)

	return b
}

func (b *Builder) AddObject(object interface{}) *Builder {
	str, err := json.MarshalIndent(&object, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	b.AddString(string(str))

	return b
}

func (b *Builder) WithRequest(r *http.Request, body string) *Builder {
	b.AddString(r.URL.RequestURI())
	if body == "" {
		body = "{}"
	}
	b.AddString(base64.StdEncoding.EncodeToString([]byte(body)))

	return b
}

func (b *Builder) Build(t *time.Time, offset int) string {
	content := b.generateContent(t, offset)

	h := hmac.New(sha256.New, []byte(b.secret))
	h.Write([]byte(content))
	digest := h.Sum(nil)

	return base64.StdEncoding.EncodeToString(digest)
}

func (b *Builder) getTimeString(t *time.Time, offset int) string {
	var currentTime time.Time = time.Now()
	if t != nil {
		currentTime = *t
	}
	currentTime = currentTime.Add(time.Minute * time.Duration(offset))

	return currentTime.Truncate(time.Minute).UTC().Format("2006-01-02T15:04:05.000Z")
}

func (b *Builder) generateContent(t *time.Time, offset int) string {
	timeString := b.getTimeString(t, offset)
	content := strings.Join(append(b.parts, timeString), "_")

	return base64.StdEncoding.EncodeToString([]byte(content))
}

func NewBuilder(secret string) Builder {
	return Builder{secret: secret, parts: []string{}}
}
