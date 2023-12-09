package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "creates blank resume.yaml at your default config directory",
	Long:  "creates blank resume.yaml at your default config directory. Doesn't override existing resume.yaml",
	RunE: func(_ *cobra.Command, _ []string) error {
		configDir, err := GetJobleafDirPath()
		if err != nil {
			return fmt.Errorf("GetResumeYamlPath: %w", err)
		}

		resumePath := configDir + "/" + resumeyaml

		_, err = os.Open(resumePath)
		if err != nil {
			if os.IsNotExist(err) {
				if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
					return fmt.Errorf("os.MkdirAll(%s/%s: %w", configDir, serviceName, err)
				}

				err = os.WriteFile(resumePath, BlankResumeYaml, os.ModePerm)
				if err != nil {
					return fmt.Errorf("os.WriteFile error: %w", err)
				}

				fmt.Printf("Blank resume.yaml created at: %s\n", resumePath)

				return nil

			}

			return fmt.Errorf("os.Open(%s) unexpected error: %w", resumePath, err)
		}

		fmt.Printf("resume.yaml already exists at: %s\n", resumePath)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
