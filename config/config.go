package config

import (
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Bind each key in the config file to a command-line flag
	for _, key := range viper.AllKeys() {
		switch viper.Get(key).(type) {
		case int:
			pflag.Int(key, viper.GetInt(key), "")
		case bool:
			pflag.Bool(key, viper.GetBool(key), "")
		case float64:
			pflag.Float64(key, viper.GetFloat64(key), "")
		case string:
			pflag.String(key, viper.GetString(key), "")
		default:
			pflag.String(key, viper.GetString(key), "")
		}
	}
	// Parse and bind command-line flags
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func GetConfig(key string) interface{} {
	return viper.Get(key)
}

func GetAllConfig() map[string]interface{} {
	return viper.AllSettings()
}
