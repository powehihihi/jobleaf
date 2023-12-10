package lib

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/powehihihi/jobleaf/internal/model"
)

func LoadFromFile(path string) (*model.Resume, error) {
	var resume model.Resume

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%s): %w", path, err)
	}

	if err := yaml.Unmarshal(file, &resume); err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal(%s): %w", path, err)
	}

	return &resume, nil
}

func SaveToFile(path string, resume *model.Resume) error {
	yaml, err := yaml.Marshal(resume)
	if err != nil {
		return fmt.Errorf("yaml.Marshal: %w", err)
	}

	if err := os.WriteFile(path, yaml, os.ModePerm); err != nil {
		return fmt.Errorf("os.WriteFile: %w", err)
	}

	return nil
}

func Generate(templateString string, writer io.Writer, resume *model.Resume) error {
	tmpl, err := template.New("main0").
		Funcs(template.FuncMap{"JoinWithComa": func(strs []string) string {
			return strings.Join(strs, ", ")
		}}).
		Parse(templateString)
	if err != nil {
		return fmt.Errorf("parsing template error: %w", err)
	}

	err = tmpl.Execute(writer, resume)
	if err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}
