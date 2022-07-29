package global

type HttpClient struct {
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

type ScriptsConfig struct {
	AddExpClient  HttpClient `yaml:"add_exp_client"`
	AddCoinClient HttpClient `yaml:"add_coin_client"`
}

var Conf = &ScriptsConfig{}
