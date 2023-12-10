package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/powehihihi/jobleaf/internal/lib"
	"github.com/powehihihi/jobleaf/internal/user"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "generates resume!",
	// TODO: Long....
	Long: "",
	RunE: func(_ *cobra.Command, _ []string) error {
		if input == "" {
			cfg, err := user.Init(nil)
			if err != nil {
				return fmt.Errorf("error while initializing user config: %w", err)
			}

			input = cfg.UserResumePath()
		}

		resume, err := lib.LoadFromFile(input)
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
			if err := lib.Generate(LatexTemplate, writer, resume); err != nil {
				return fmt.Errorf("lib.Generate: %w", err)
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
