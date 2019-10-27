package template

import (
	"bytes"
	"text/template"
)

type Input struct {
	Name   string
	Tag    string
	Date   string
	Commit string
}

// GetName returns current name based on template
func GetName(tmpl string) (string, error) {
	inp := &Input{}
	t, err := template.New("goreo").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var tmplData bytes.Buffer
	err = t.Execute(&tmplData, inp)
	if err != nil {
		return "", err
	}

	return tmplData.String(), nil
}
