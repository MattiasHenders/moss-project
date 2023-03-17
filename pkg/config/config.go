package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string, filenames []string) (*viper.Viper, error) {
	viper.AddConfigPath(path)

	for _, filename := range filenames {
		viper.SetConfigName(filename)
		if err := viper.MergeInConfig(); err != nil {
			return nil, err
		}
	}

	return viper.GetViper(), nil
}
