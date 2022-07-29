package cmd

import (
	"fmt"
	"scripts/global"
	"scripts/internal"

	"github.com/spf13/cobra"
)

var addAction string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "增加用户属性",
	Long:  "增加用户属性, 这里只能增加经验值、金币",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("action: ", addAction)
		fmt.Println("cf: ", cf)
		err := internal.InitConfig(global.Conf, cf)
		if err != nil {
			panic(err)
		}
		fmt.Printf("global.Conf: %+v\n", global.Conf)
		switch addAction {
		case "exp":
			err = internal.HandleExp()
		case "coin":
			err = internal.HandleCoin()
		}
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&addAction, "add.action", "a", "exp", "action, default is exp")
}
