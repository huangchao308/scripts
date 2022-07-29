package cmd

import (
	"scripts/internal"

	"github.com/spf13/cobra"
)

var calCmd = &cobra.Command{
	Use:   "cal",
	Short: "计算经验值",
	Long:  "计算经验值",
	Run: func(cmd *cobra.Command, args []string) {
		err := internal.CalExp()
		if err != nil {
			panic(err)
		}
	},
}
