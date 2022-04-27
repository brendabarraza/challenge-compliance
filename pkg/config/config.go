package config

import (
	custom_error "meli-challenge-compliance/pkg/errors"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort               string `mapstructure:"HTTP_PORT"`
	TokenFileLocation      string `mapstructure:"TOKEN_FILE_LOCATION"`
	CredentialFileLocation string `mapstructure:"CREDENTIAL_FILE_LOCATION"`
}

// LoadConfig lee la configuracion de archivo y la retorna
func LoadConfig(path string) (*Config, error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, custom_error.ErrBadConfiguration
	}

	//	Viper parsea lo que se encuentra en el archivo dev.env
	config := &Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, custom_error.ErrBadConfiguration
	}
	return config, nil
}
