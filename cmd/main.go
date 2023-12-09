package main

import (
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/powehihihi/jobleaf/resume"
)

func main() {
	var resume resume.Resume

	file, err := os.ReadFile("examples/resume.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &resume)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("resume.tmpl.tex").
		Funcs(template.FuncMap{"JoinWithComa": JoinWithComa}).
		ParseFiles(
			"template/resume.tmpl.tex",
			/*		"template/tex/head.gotex",
					"template/tex/authors.gotex",
					"template/tex/contact.gotex",
					"template/tex/summary.gotex",
					"template/tex/experience.gotex",
					"template/tex/projects.gotex",
					"template/tex/education.gotex",
					"template/tex/skills.gotex", */
		)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, resume)
	if err != nil {
		panic(err)
	}
}

func JoinWithComa(strs []string) string {
	return strings.Join(strs, ", ")
}
