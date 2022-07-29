package internal

import (
	"errors"
	"os"
	"scripts/global"

	"github.com/spf13/viper"
)

func InitConfig(cfg *global.ScriptsConfig, name string) error {
	if name == "" {
		return errors.New("config is required")
	}
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	cfg.AddExpClient.Host = viper.GetString("add_exp_client.host")
	cfg.AddExpClient.Path = viper.GetString("add_exp_client.path")
	cfg.AddCoinClient.Host = viper.GetString("add_coin_client.host")
	cfg.AddCoinClient.Path = viper.GetString("add_coin_client.path")

	return nil
}
