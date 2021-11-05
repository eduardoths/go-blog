package input

import (
	"regexp"
	"strings"
)

func TransformSingleLine(text string) string {
	text = strings.TrimSpace(text)
	space := regexp.MustCompile(`\s+`)
	text = space.ReplaceAllString(text, " ")
	return text
} 
