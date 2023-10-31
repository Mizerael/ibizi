package putMessage

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	containerPath      string
	messagePath        string
	stegocontainerPath string
}

var rootCmd = &cobra.Command{
	Use:   "put-message",
	Short: "hide message in text",
	Long:  `Solution first part of first lab on course infoSec SSU`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
