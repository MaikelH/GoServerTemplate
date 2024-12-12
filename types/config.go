package types

type Configuration struct {
	LogLevel       string `mapstructure:"APP_LOG_LEVEL"`
	ListenAddress  string `mapstructure:"APP_LISTEN_ADDRESS"`
	OpenAPIAddress string `mapstructure:"APP_OPENAPI_ADDRESS"`
}
