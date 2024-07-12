package global

import (
	"fmt"

	"github.com/spf13/viper"
)

// only read setting from file
type Config struct {
	serviceName string
	environment string
}

var setting *Config

func InitSetting() error {
	setting = &Config{
		serviceName: viper.GetString("service.name"),
		environment: viper.GetString("service.environment"),
	}

	if (setting.serviceName == "") || (setting.environment == "") {
		return fmt.Errorf("Config is invalid")
	}

	return nil
}

func ServiceName() string {
	if setting == nil {
		return ""
	}

	return setting.serviceName
}

func Environment() string {
	if setting == nil {
		return ""
	}

	return setting.environment
}
