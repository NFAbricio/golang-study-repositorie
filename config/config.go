package config

import (
	"fmt"

	"github.com/go-chi/jwtauth"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDriver   string `mapstructure:"DATABASE_DRIVER"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	webServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	jwtSecret        string `mapstructure:"JWT_SECRET"`
	jwtExpiresIn     string `mapstructure:"JWT_EXPIRES_IN"`
	jwtAuth          *jwtauth.JWTAuth
}

func LoadConf(path string) (*Config, error) {
	var enviroment *Config

	viper.SetConfigName(".env") // name of file who viper will search
	viper.SetConfigType("env")  // set the type of return of configuration. e.g "json"
	viper.AddConfigPath(path)   // add path to viper look and search
	//viper.SetConfigFile(".env")       //show to viper the exactly path
	viper.AutomaticEnv() // will check the part of "DATABASE_DRIVER" etc...

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&enviroment); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	} // will add to strcuture

	enviroment.jwtAuth = jwtauth.New("HS256", []byte(enviroment.jwtSecret), nil)

	return enviroment, nil
}
