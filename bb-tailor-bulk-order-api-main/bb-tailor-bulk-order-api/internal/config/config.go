package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"PORT" validate:"required,numeric"`
		Host string `mapstructure:"HOST" validate:"required"`
	} `mapstructure:"server"`

	BasicAuth struct {
		Username string `mapstructure:"BASIC_AUTH_USERNAME"`
		Password string `mapstructure:"BASIC_AUTH_PASSWORD" validate:"required"`
	} `mapstructure:"basicauth"`

	MySqlDb struct {
		DbPort         string `mapstructure:"DB_PORT" validate:"required,numeric"`
		DbHost         string `mapstructure:"DB_HOST" validate:"required"`
		DbName         string `mapstructure:"DB_NAME" validate:"required"`
		DbUser         string `mapstructure:"DB_USER" validate:"required"`
		DbUserPassword string `mapstructure:"DB_USER_PASSWORD" validate:"required"`
	} `mapstructure:"mysqldb"`

	JwtAuth struct {
		AccessTokenSecret     string `mapstructure:"ACCESS_TOKEN_SECRET" validate:"required"`
		RefreshTokenSecret    string `mapstructure:"REFRESH_TOKEN_SECRET" validate:"required"`
		AccessTokenExpDays    int    `mapstructure:"ACCESS_TOKEN_EXP_DAYS"`
		RefreshTokenExpMonths int    `mapstructure:"REFRESH_TOKEN_EXP_MONTHS"`
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	v := viper.New()

	//Set the defaults first
	v.SetDefault("BASIC_AUTH_USERNAME", "admin")
	v.SetDefault("ACCESS_TOKEN_EXP_DAYS", 7)
	v.SetDefault("REFRESH_TOKEN_EXP_MONTHS", 3)

	v.SetConfigFile(".env")
	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.AutomaticEnv()

	err := v.ReadInConfig()

	if err == nil {
		fmt.Println("using configuration file:", v.ConfigFileUsed())
	} else {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	//create the cfg using the loaded env file variables
	cfg.Server.Port = v.GetString("PORT")
	cfg.Server.Host = v.GetString("HOST")
	cfg.BasicAuth.Username = v.GetString("BASIC_AUTH_USERNAME")
	cfg.BasicAuth.Password = v.GetString("BASIC_AUTH_PASSWORD")
	cfg.MySqlDb.DbHost = v.GetString("DB_HOST")
	cfg.MySqlDb.DbPort = v.GetString("DB_PORT")
	cfg.MySqlDb.DbName = v.GetString("DB_NAME")
	cfg.MySqlDb.DbUser = v.GetString("DB_USER")
	cfg.MySqlDb.DbUserPassword = v.GetString("DB_USER_PASSWORD")
	cfg.JwtAuth.AccessTokenSecret = v.GetString("ACCESS_TOKEN_SECRET")
	cfg.JwtAuth.RefreshTokenSecret = v.GetString("REFRESH_TOKEN_SECRET")
	cfg.JwtAuth.AccessTokenExpDays = v.GetInt("ACCESS_TOKEN_EXP_DAYS")
	cfg.JwtAuth.RefreshTokenExpMonths = v.GetInt("REFRESH_TOKEN_EXP_MONTHS")

	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(cfg)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return fmt.Errorf("invalid Validation Error: %w", err)
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		return fmt.Errorf("validation errors: %w", err)
	}

	return nil
}
