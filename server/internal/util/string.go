package util

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Sanitize(text string) string {
	text = strings.TrimSpace(text)

	return text
}

var caser = cases.Title(language.AmericanEnglish)

func Capitalize(text string) string {
	return caser.String(text)
}
