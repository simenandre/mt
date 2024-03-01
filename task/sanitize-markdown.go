package task

import (
	"bytes"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func markdownToText(input string) string {
	// Convert Markdown to HTML
	html := blackfriday.Run([]byte(input))

	// Use bluemonday to strip HTML tags
	p := bluemonday.StrictPolicy()
	text := p.SanitizeBytes(html)

	// remove newlines
	text = bytes.ReplaceAll(text, []byte("\n"), []byte(""))

	return string(text)
}
