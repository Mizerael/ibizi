package getMessage

import (
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	MessagePath        string
	StegocontainerPath string
}

var conf Config

var rootCmd = &cobra.Command{
	Use:   "get-message",
	Short: "show hide message",
	Long:  `Solution second part of first lab on course infoSec SSU`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.Flags().StringVarP(&conf.StegocontainerPath, "stego", "s", "",
		"PATH to stegocontainer")
	rootCmd.Flags().StringVarP(&conf.MessagePath, "message", "m", "",
		"PATH to message")
}

func Execute() *Config {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	return &conf
}
