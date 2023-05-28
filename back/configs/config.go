package configs

import "time"

type Config struct {
	DB            *DBConfig
	JWTSecretKey  string
	TokenDuration time.Duration
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
	SSLMode  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "frontend",
			DBName:   "postgres",
			Password: "password",
			SSLMode:  "disable",
		},
		JWTSecretKey:  "your-secret-key",
		TokenDuration: time.Hour * 24, // токен действителен в течение 24 часов

	}
}
