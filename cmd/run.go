package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/powehihihi/jobleaf/internal/resume"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "generates resume!",
	// TODO: Long....
	Long: "",
	RunE: func(_ *cobra.Command, _ []string) error {
		var err error

		if input == "" {
			input, err = GetUserResumeYamlPath()
			if err != nil {
				return fmt.Errorf("GetUserResumeYamlPath: %w", err)
			}
		}

		resume, err := LoadResume(input)
		if err != nil {
			return fmt.Errorf("LoadResume: %w", err)
		}

		var writer io.WriteCloser

		if output == "" {
			writer = os.Stdout
		} else {
			writer, err = os.Create(output)
			if err != nil {
				return fmt.Errorf("os.Create: %w", err)
			}
		}

		defer writer.Close()

		if genLatex {
			tmpl, err := template.New("resume.tmpl.tex").
				Funcs(template.FuncMap{"JoinWithComa": func(strs []string) string {
					return strings.Join(strs, ", ")
				}}).
				Parse(LatexTemplate)
			if err != nil {
				return fmt.Errorf("parsing template error: %w", err)
			}

			err = tmpl.Execute(writer, resume)
			if err != nil {
				return fmt.Errorf("executing template: %w", err)
			}
		}

		return nil
	},
}

var (
	input    string
	output   string
	genLatex bool
)

func init() {
	rootCmd.AddCommand(runCmd)

	flags := runCmd.Flags()

	flags.StringVarP(&input, "input", "i", "", "path to your resume.yaml")
	flags.StringVarP(&output, "output", "o", "", "where to put generated resume")
	flags.BoolVar(&genLatex, "latex", true, "where to put generated resume")
}

func LoadResume(path string) (*resume.Resume, error) {
	var resume resume.Resume

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%s): %w", path, err)
	}

	if err := yaml.Unmarshal(file, &resume); err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal(%s, {}): %w", file, err)
	}

	return &resume, nil
}

func GetJobleafDirPath() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("os.UserConfigDir: %w", err)
	}

	return configPath + "/jobleaf", nil
}

func GetUserResumeYamlPath() (string, error) {
	dir, err := GetJobleafDirPath()
	if err != nil {
		return "", fmt.Errorf("GetJobleafDirPath: %w", err)
	}

	return dir + "/resume.yaml", nil
}
