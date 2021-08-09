package cmd

import (
	"github.com/dupotato/gocommon/config"

	"github.com/spf13/cobra"
)

var defaultConfigCmd = &cobra.Command{
	Use:   "default",
	Short: "print default config",
	Run: func(*cobra.Command, []string) {
		config.PrintConfig()
	},
}

func init() {
	//fmt.Print("root cmd begin\n")
	rootCmd.AddCommand(defaultConfigCmd)
}
