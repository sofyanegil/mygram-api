package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env        string `mapstructure:"ENV"`
	DBUrl      string `mapstructure:"DB_URL"`
	DBHost     string `mapstructure:"PGHOST"`
	DBUsername string `mapstructure:"PGUSER"`
	DBPassword string `mapstructure:"PGPASSWORD"`
	DBName     string `mapstructure:"PGDATABASE"`
	DBPort     string `mapstructure:"PGPORT"`

	ServerPort string `mapstructure:"PORT"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
