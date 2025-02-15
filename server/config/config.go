package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type DatabaseConfig struct {
	Type   string       `validate:"required,oneof=mysql sqlite"`
	MySQL  MySQLConfig  `validate:"omitempty"`
	SQLite SQLiteConfig `validate:"required_if=Type sqlite"`
}

type SMTPConfig struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"omitempty"`
	Password string `validate:"omitempty"`
	Auth     string
	From     string `validate:"required,email"`
	NoTLS    bool
}

type Config struct {
	Database       DatabaseConfig
	CookieSecret   string `validate:"required"`
	CookieMaxAge   int    `validate:"min=0"`
	CookieHttpOnly bool
	CookieSecure   bool
	SessionName    string `validate:"required"`
	AllowOrigin    string `validate:"required,url"`
	AppName        string `validate:"required"`
	AppUrl         string `validate:"required,url"`
	SMTP           SMTPConfig
	RateLimit      RateLimitConfig
	DataDir        string `validate:"required"`
}

type RateLimitConfig struct {
	Enabled             bool
	Limit               int `validate:"min=1"`
	Window              int `validate:"min=1"`
	LoginLimit          int `validate:"min=1"`
	LoginWindow         int `validate:"min=1"`
	PasswordResetLimit  int `validate:"min=1"`
	PasswordResetWindow int `validate:"min=1"`
}

type MySQLConfig struct {
	User     string `validate:"required"`
	Password string `validate:"required"`
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	Database string `validate:"required"`
}

type SQLiteConfig struct {
	FilePath string `validate:"required"`
}

var validate = validator.New()

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Database: DatabaseConfig{
			Type: os.Getenv("DB_TYPE"),
			MySQL: MySQLConfig{
				User:     os.Getenv("MYSQL_USER"),
				Password: os.Getenv("MYSQL_PASSWORD"),
				Host:     os.Getenv("MYSQL_HOST"),
				Port:     os.Getenv("MYSQL_PORT"),
				Database: os.Getenv("MYSQL_DATABASE"),
			},
			SQLite: SQLiteConfig{FilePath: os.Getenv("SQLITE_FILE_PATH")},
		},
		CookieSecret:   os.Getenv("COOKIE_SECRET"),
		CookieMaxAge:   getIntEnv("COOKIE_MAX_AGE", 2592000),
		CookieHttpOnly: os.Getenv("COOKIE_HTTP_ONLY") != "false",
		CookieSecure:   os.Getenv("COOKIE_SECURE") != "false",
		SessionName:    os.Getenv("SESSION_NAME"),
		AllowOrigin:    os.Getenv("ALLOW_ORIGIN"),
		AppName:        os.Getenv("APP_NAME"),
		AppUrl:         os.Getenv("APP_URL"),
		SMTP: SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     os.Getenv("SMTP_PORT"),
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Auth:     os.Getenv("SMTP_AUTH"),
			From:     os.Getenv("SMTP_FROM"),
			NoTLS:    os.Getenv("SMTP_NO_TLS") == "true",
		},
		RateLimit: RateLimitConfig{
			Enabled:             os.Getenv("RATE_LIMIT_ENABLED") != "false",
			Limit:               getIntEnv("RATE_LIMIT_LIMIT", 100),
			Window:              getIntEnv("RATE_LIMIT_WINDOW", 1),
			LoginLimit:          getIntEnv("RATE_LIMIT_LOGIN_LIMIT", 20),
			LoginWindow:         getIntEnv("RATE_LIMIT_LOGIN_WINDOW", 10),
			PasswordResetLimit:  getIntEnv("RATE_LIMIT_PASSWORD_RESET_LIMIT", 5),
			PasswordResetWindow: getIntEnv("RATE_LIMIT_PASSWORD_RESET_WINDOW", 10),
		},
		DataDir: os.Getenv("DATA_DIR"),
	}

	setDefaults(cfg)

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	return cfg, nil
}

func setDefaults(cfg *Config) {
	if cfg.Database.Type == "" {
		cfg.Database.Type = "sqlite"
		log.Printf("Using default DB_TYPE: %s", cfg.Database.Type)
	}
	if cfg.Database.Type == "sqlite" && cfg.Database.SQLite.FilePath == "" {
		cfg.Database.SQLite.FilePath = "./db.sqlite"
		log.Printf("Using default SQLITE_FILE_PATH: %s", cfg.Database.SQLite.FilePath)
	}
	if cfg.CookieSecret == "" {
		cfg.CookieSecret = "secret"
		log.Printf("Using default COOKIE_SECRET: %s", cfg.CookieSecret)
	}
	if cfg.SessionName == "" {
		cfg.SessionName = "mysession"
		log.Printf("Using default SESSION_NAME: %s", cfg.SessionName)
	}
	if cfg.AllowOrigin == "" {
		cfg.AllowOrigin = "http://localhost:3001"
		log.Printf("Using default ALLOW_ORIGIN: %s", cfg.AllowOrigin)
	}
	if cfg.AppName == "" {
		cfg.AppName = "App Name"
		log.Printf("Using default APP_NAME: %s", cfg.AppName)
	}
	if cfg.AppUrl == "" {
		cfg.AppUrl = "http://localhost:3001"
		log.Printf("Using default APP_URL: %s", cfg.AppUrl)
	}
	if cfg.SMTP.Host == "" {
		cfg.SMTP.Host = "localhost"
		log.Printf("Using default SMTP_HOST: %s", cfg.SMTP.Host)
	}
	if cfg.SMTP.Port == "" {
		cfg.SMTP.Port = "587"
		log.Printf("Using default SMTP_PORT: %s", cfg.SMTP.Port)
	}
	if cfg.SMTP.From == "" {
		cfg.SMTP.From = "no-reply@example.com"
		log.Printf("Using default SMTP_FROM: %s", cfg.SMTP.From)
	}
	if cfg.DataDir == "" {
		cfg.DataDir = "./data"
		log.Printf("Using default DATA_DIR: %s", cfg.DataDir)
	}
}

func getIntEnv(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid value for %s: %s. Using default: %d", key, value, defaultValue)
		return defaultValue
	}
	return intValue
}
