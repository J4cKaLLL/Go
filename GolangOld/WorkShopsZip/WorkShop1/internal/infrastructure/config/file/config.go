package file

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	// Inicialización con Singleton
	onceInit sync.Once
)

// Config ...
type Config struct {
	Repository ClientConfig
}

// ClientConfig ...
type ClientConfig struct {
	Client []ServiceConfig
}

// ServiceConfig ...
type ServiceConfig struct {
	Service string
	Country string
	IP      string
	Port    string
}

// NewFileConfig ...
func NewFileConfig() (*Config, error) {
	var err error
	config := &Config{}

	onceInit.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("./config")

		if err = viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
		err = viper.Unmarshal(&config)
		if err != nil {
			fmt.Printf("unable to decode into struct, %v", err)
		}
	})
	return config, err
}
