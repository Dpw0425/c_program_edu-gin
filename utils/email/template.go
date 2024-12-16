package email

import (
	"bytes"
	"text/template"
)

func RenderTemplate(text []byte, data interface{}) ([]byte, error) {
	tmpl, _ := template.New("tmpl").Parse(string(text))

	var content bytes.Buffer
	if err := tmpl.Execute(&content, data); err != nil {
		return nil, err
	}

	return content.Bytes(), nil
}
