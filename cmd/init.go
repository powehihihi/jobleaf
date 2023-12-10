package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/powehihihi/jobleaf/internal/user"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "creates blank resume.yaml at your default config directory",
	Long:  "creates blank resume.yaml at your default config directory. Doesn't override existing resume.yaml",
	RunE: func(_ *cobra.Command, _ []string) error {
		cfg, err := user.Init(BlankResumeYaml)
		if err != nil {
			return fmt.Errorf("user.Init: %w", err)
		}

		fmt.Println("User resume: ", cfg.UserResumePath())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
