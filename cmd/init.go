package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/powehihihi/jobleaf/internal/user"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "creates blank resume.yaml at your default config directory",
	Long:  "creates blank resume.yaml at your default config directory. Doesn't override existing resume.yaml",
	RunE: func(_ *cobra.Command, args []string) error {
		cfg, err := user.Init(BlankResumeYaml)
		if err != nil {
			return fmt.Errorf("user.Init: %w", err)
		}

		fmt.Println("User resume: ", cfg.UserResumePath())

		if edit {
			var arguments []string

			switch runtime.GOOS {
			case "windows":
				arguments = []string{"cmd", "/C", "notepad"}
			default:
				defaultEditor, _ := os.LookupEnv("EDITOR")
				arguments = []string{defaultEditor}
			}

			editorIndex := len(arguments) - 1

			if len(args) > 0 {
				arguments[editorIndex] = args[0]
			}

			arguments = append(arguments, cfg.UserResumePath())

			command := exec.Command(arguments[0], arguments[1:]...)
			command.Stdin = os.Stdin
			command.Stdout = os.Stdout

			if err := command.Run(); err != nil {
				fmt.Println("failed to run:")
				fmt.Println(strings.Join(arguments, " "))
			}
		}

		return nil
	},
}

var edit bool

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().
		BoolVarP(&edit, "edit", "e", false, "Use your flag to edit your resume.yaml. If editor is not specified as argument, $EDITOR (notepad on windows) will be used")
}
