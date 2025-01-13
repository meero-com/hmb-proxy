package config

import (
    "log"

	"github.com/spf13/pflag"
    "github.com/spf13/viper"
)

func InitConfig() {
	// Define command-line flags
    pflag.String("env", "development", "Application environment")
    pflag.Int("server.port", 8080, "Server port")

    // Parse command-line flags
    pflag.Parse()

    // Bind flags to Viper
    viper.BindPFlags(pflag.CommandLine)

    viper.SetConfigName("config")
    viper.AddConfigPath("config/")
    viper.SetConfigType("yaml")

    // Read the configuration file
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }
}

func GetConfig(key string) interface{} {
    return viper.Get(key)
}