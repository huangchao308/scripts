package cmd

import (
	"scripts/global"
	"scripts/internal"

	"github.com/spf13/cobra"
)

var cf string
var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitConfig(global.Conf, cf)
	},
}

func Execute() error {
	// fmt.Printf("config file: %s, config: %+v", cf, scriptsConfig)
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(calCmd)
	rootCmd.PersistentFlags().StringVarP(&cf, "config", "c", "", "config file, required")
}
