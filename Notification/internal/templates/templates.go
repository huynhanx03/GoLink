package templates

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

//go:embed *.html
var templateFS embed.FS

var tmpl *template.Template

func init() {
	var err error
	tmpl, err = template.ParseFS(templateFS, "*.html")
	if err != nil {
		panic(fmt.Sprintf("failed to parse notification templates: %v", err))
	}
}

// Render renders a named template with the provided data and returns the HTML string.
func Render(templateName string, data map[string]any) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, templateName, data); err != nil {
		return "", fmt.Errorf("render template %s: %w", templateName, err)
	}
	return buf.String(), nil
}
