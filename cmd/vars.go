package cmd

import (
	_ "embed"
)

const (
	serviceName = "jobleaf"
	resumeyaml  = "resume.yaml"
)

// exported for using by golang embed
var (
	BlankResumeYaml []byte
	LatexTemplate   string
)
