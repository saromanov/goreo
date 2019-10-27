package template

import (
	"os"
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
	t, err := template.New("goreao").Parse(tmpl)
	if err != nil {
		return "", err
	}
	err = t.ExecuteTemplate(os.Stdout, "T", inp)
	if err != nil {
		return "", err
	}

	return "", nil
}
