package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	INFURA struct {
		URL        string `mapstructure:"url"`
		NETWORKING int64  `mapstructure:"networking"`
		API_KEY    string `mapstructure:"api_key"`
	} `mapstructure:"infura"`
	WALLET struct {
		PRIVATE_KEY string `mapstructure:"private_key"`
	} `mapstructure:"wallet"`
}

func LoadConfig() (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigFile("app.yaml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
