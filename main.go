package main

import (
	_ "embed"

	"github.com/powehihihi/jobleaf/cmd"
)

func main() {
	cmd.Execute()
}

// So... embed doesn't allow relative pathes. And i really need to pass templates and examples to cmd
// Solution: https://stackoverflow.com/a/67357103/14075443

var (
	//go:embed examples/empty.yaml
	blankResumeYaml []byte
	//go:embed templates/resume.tmpl.tex
	latexTemplate string
)

func init() {
	cmd.BlankResumeYaml = blankResumeYaml
	cmd.LatexTemplate = latexTemplate
}
