package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddr               string
	MongoURI                 string
	ActiveEnvProfile         string `mapstructure:"ACTIVE_ENV_PROFILE"`
	ServerHost               string `mapstructure:"SERVER_HOST"`
	ServerPort               string `mapstructure:"SERVER_PORT"`
	MongoHost                string `mapstructure:"MONGO_HOST"`
	MongoPort                string `mapstructure:"MONGO_PORT"`
	MongoUser                string `mapstructure:"MONGO_USER"`
	MongoPass                string `mapstructure:"MONGO_PASS"`
	MongoDBName              string `mapstructure:"MONGO_DBNAME"`
	SlackAuthToken           string `mapstructure:"SLACK_AUTH_TOKEN"`
	SlackChannelId           string `mapstructure:"SLACK_CHANNEL_ID"`
	SlackDebugEnabled        bool   `mapstructure:"SLACK_DEBUG_ENABLED"`
	SlackNotificationEnabled bool   `mapstructure:"SLACK_NOTIFICATION_ENABLED"`
	ConfigFilePath           string `mapstructure:"CONFIG_FILE_PATH"`
	ScriptFilePath           string `mapstructure:"SCRIPT_FILE_PATH"`
}

func LoadConfig(filename string) (config *Config, err error) {

	// load from file
	viper.AddConfigPath(".")
	viper.SetConfigName(filename)
	viper.SetConfigType("env")

	// set defaults
	viper.SetDefault("ACTIVE_ENV_PROFILE", "local")
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("MONGO_HOST", "localhost")
	viper.SetDefault("MONGO_PORT", "27017")
	viper.SetDefault("MONGO_USER", "root")
	viper.SetDefault("MONGO_PASS", "password")
	viper.SetDefault("MONGO_DBNAME", "secrets-operator")
	viper.SetDefault("SLACK_DEBUG_ENABLED", false)
	viper.SetDefault("SLACK_NOTIFICATION_ENABLED", false)
	viper.SetDefault("CONFIG_FILE_PATH", "config/config.toml")
	viper.SetDefault("SCRIPT_FILE_PATH", "config/pipelineScript.sh")

	// load from env and override defaults and values loaded from config file
	// first one in row takes precedence:
	// environment variables -> config file -> default values
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found.", err)
		} else {
			log.Fatalln("Something happened while reading config file.", err)
			return
		}
	}

	err = viper.Unmarshal(&config)

	// set compound configuration variables
	config.ServerAddr = fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	config.MongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-256",
		config.MongoUser, config.MongoPass, config.MongoHost, config.MongoPort, config.MongoDBName)

	return config, nil
}
