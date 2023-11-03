package putMessage

import (
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	ContainerPath      string
	MessagePath        string
	StegocontainerPath string
}

var conf *Config

var rootCmd = &cobra.Command{
	Use:   "put-message",
	Short: "hide message in text",
	Long:  `Solution first part of first lab on course infoSec SSU`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.Flags().StringVarP(&conf.StegocontainerPath, "stego", "s", "",
		"PATH to stegocontainer")
	rootCmd.Flags().StringVarP(&conf.MessagePath, "message", "m", "",
		"PATH to message")
	rootCmd.Flags().StringVarP(&conf.StegocontainerPath, "container", "c", "",
		"PATH to container")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
