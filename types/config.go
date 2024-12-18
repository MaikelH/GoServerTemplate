package types

type Configuration struct {
	LogLevel       int    `mapstructure:"APP_LOG_LEVEL"`
	ListenAddress  string `mapstructure:"APP_LISTEN_ADDRESS"`
	OpenAPIAddress string `mapstructure:"APP_OPENAPI_ADDRESS"`
	Auth0Audience  string `mapstructure:"AUTH0_AUDIENCE"`
	Auth0Domain    string `mapstructure:"AUTH0_DOMAIN"`
}
