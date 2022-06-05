package template

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/saromanov/goreo/internal/utils"
)

type Input struct {
	Name      string
	Os        string
	Tag       string
	Date      string
	Commit    string
	Timestamp string
	Platform  string
}

// GetName returns current name based on template
func GetName(tmpl, osName, platform string) (string, error) {
	projectName, err := utils.GetProjectName()
	if err != nil {
		return "", err
	}
	inp := &Input{
		Name:      projectName,
		Os:        osName,
		Date:      time.Now().UTC().Format(time.RFC3339),
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
		Platform:  platform,
	}
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
