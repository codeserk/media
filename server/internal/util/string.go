package util

import (
	"regexp"
	"strings"
)

func Sanitize(text string) string {
	text = strings.TrimSpace(text)
	text = removeEmojisAndSpecialChars(text)

	return text
}

func removeEmojisAndSpecialChars(text string) string {
	// Define a regular expression to match emojis and special characters
	re := regexp.MustCompile(`[^\w\s]`)

	// Replace emojis and special characters with an empty string
	text = re.ReplaceAllString(text, "")

	return text
}
