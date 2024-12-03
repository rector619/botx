package config

import "os"

type Configuration struct {
	Server        ServerConfiguration
	Databases     Database
	TestDatabases Database
	App           App
	IPStack       IPStack
	Microservices Microservices
}

type BaseConfig struct {
	SERVER_PORT                      string `mapstructure:"SERVER_PORT"`
	SERVER_SECRET                    string `mapstructure:"SERVER_SECRET"`
	SERVER_ACCESSTOKENEXPIREDURATION int    `mapstructure:"SERVER_ACCESSTOKENEXPIREDURATION"`

	APP_NAME string `mapstructure:"APP_NAME"`
	APP_KEY  string `mapstructure:"APP_KEY"`
	APP_URL  string `mapstructure:"APP_URL"`

	CONNECTION_STRING string `mapstructure:"CONNECTION_STRING"`
	DB_NAME           string `mapstructure:"DB_NAME"`
	MIGRATE           bool   `mapstructure:"MIGRATE"`

	TEST_CONNECTION_STRING string `mapstructure:"TEST_CONNECTION_STRING"`
	TEST_DB_NAME           string `mapstructure:"TEST_DB_NAME"`
	TEST_MIGRATE           bool   `mapstructure:"TEST_MIGRATE"`

	IPSTACK_KEY      string `mapstructure:"IPSTACK_KEY"`
	IPSTACK_BASE_URL string `mapstructure:"IPSTACK_BASE_URL"`

	AUTH_MS         string `mapstructure:"AUTH_MS"`
	NOTIFICATION_MS string `mapstructure:"NOTIFICATION_MS"`
}

func (config *BaseConfig) SetupConfigurationn() *Configuration {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.SERVER_PORT
	}
	return &Configuration{
		Server: ServerConfiguration{
			Port:                          port,
			Secret:                        config.SERVER_SECRET,
			AccessTokenExpirationDuration: config.SERVER_ACCESSTOKENEXPIREDURATION,
		},
		Databases: Database{
			ConnectionString: config.CONNECTION_STRING,
			DBName:           config.DB_NAME,
			Migrate:          config.MIGRATE,
		},
		TestDatabases: Database{
			ConnectionString: config.TEST_CONNECTION_STRING,
			DBName:           config.TEST_DB_NAME,
			Migrate:          config.TEST_MIGRATE,
		},
		App: App{
			Name: config.APP_NAME,
			Key:  config.APP_KEY,
			Url:  config.APP_URL,
		},
		IPStack: IPStack{
			Key:     config.IPSTACK_KEY,
			BaseUrl: config.IPSTACK_BASE_URL,
		},
		Microservices: Microservices{
			Auth:         config.AUTH_MS,
			Notification: config.NOTIFICATION_MS,
		},
	}
}
