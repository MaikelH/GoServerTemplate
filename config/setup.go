package config

import (
	"errors"
	"github.com/spf13/viper"
	"log/slog"
	"reflect"
)

// SetupConfig loads the configuration from the config file and environment variables.
func SetupConfig[T any]() (*T, error) {
	viper.AddConfigPath("./config")
	// unit tests require relative path
	viper.AddConfigPath("../config")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("hcl")

	viper.AutomaticEnv()

	// Bind automatic variables defined in types.Configuration
	var config T
	err := bindVariables(config)
	if err != nil {
		return nil, err
	}

	// Set default values
	err = viper.ReadInConfig()
	if err != nil { // Handle service_error reading the config file
		var notFoundError = viper.ConfigFileNotFoundError{}
		if errors.As(err, &notFoundError) {
			slog.Warn("Config file not found, only using environment variables")
		} else {
			return nil, err
		}
	}

	// Unmarshal into pre-defined configuration struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// bindVariables automatically binds environment variables defined in the struct.
// Uses the mapstructure tag as name of the environment variable.
func bindVariables[T any](config T) error {
	configType := reflect.TypeOf(config)

	for i := 1; i < configType.NumField(); i++ {
		tag := configType.Field(i).Tag
		err := viper.BindEnv(tag.Get("config"))
		if err != nil {
			return err
		}
	}

	return nil
}
