package config

import (
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
)

var (
	CFG      *Config
	validate = validator.New()
)

type Config struct {
	DBType       string       `validate:"required,oneof=mysql sqlite"`
	MySQL        MySQLConfig  `validate:"omitempty"`
	SQLite       SQLiteConfig `validate:"required_if=DBType sqlite"`
	CookieSecret string       `validate:"required"`
	SessionName  string       `validate:"required"`
	AllowOrigin  string       `validate:"required,url"`
	AppName      string       `validate:"required"`
	SmtpHost     string       `validate:"required"`
	SmtpPort     string       `validate:"required"`
	SmtpUser     string       `validate:"omitempty"`
	SmtpPassword string       `validate:"omitempty"`
	SmtpAuth     string
	SmtpFrom     string `validate:"required,email"`
	SmtpNoTLS    bool
	RateLimit    RateLimitConfig
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

func (c MySQLConfig) GetUser() string {
	return c.User
}

func (c MySQLConfig) GetPassword() string {
	return c.Password
}

func (c MySQLConfig) GetHost() string {
	return c.Host
}

func (c MySQLConfig) GetPort() string {
	return c.Port
}

func (c MySQLConfig) GetDatabase() string {
	return c.Database
}

type SQLiteConfig struct {
	FilePath string `validate:"required"`
}

func (c SQLiteConfig) GetFilePath() string {
	return c.FilePath
}

func LoadConfig() error {
	cfg := &Config{
		DBType: os.Getenv("DB_TYPE"),
		MySQL: MySQLConfig{
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
			Database: os.Getenv("MYSQL_DATABASE"),
		},
		SQLite:       SQLiteConfig{FilePath: os.Getenv("SQLITE_FILE_PATH")},
		CookieSecret: os.Getenv("COOKIE_SECRET"),
		SessionName:  os.Getenv("SESSION_NAME"),
		AllowOrigin:  os.Getenv("ALLOW_ORIGIN"),
		AppName:      os.Getenv("APP_NAME"),
		SmtpHost:     os.Getenv("SMTP_HOST"),
		SmtpPort:     os.Getenv("SMTP_PORT"),
		SmtpUser:     os.Getenv("SMTP_USER"),
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),
		SmtpAuth:     os.Getenv("SMTP_AUTH"),
		SmtpFrom:     os.Getenv("SMTP_FROM"),
		SmtpNoTLS:    os.Getenv("SMTP_NO_TLS") == "true",
		RateLimit: RateLimitConfig{
			Enabled:             os.Getenv("RATE_LIMIT_ENABLED") != "false",
			Limit:               getIntEnv("RATE_LIMIT_LIMIT", 100),
			Window:              getIntEnv("RATE_LIMIT_WINDOW", 1),
			LoginLimit:          getIntEnv("RATE_LIMIT_LOGIN_LIMIT", 20),
			LoginWindow:         getIntEnv("RATE_LIMIT_LOGIN_WINDOW", 10),
			PasswordResetLimit:  getIntEnv("RATE_LIMIT_PASSWORD_RESET_LIMIT", 5),
			PasswordResetWindow: getIntEnv("RATE_LIMIT_PASSWORD_RESET_WINDOW", 10),
		},
	}

	setDefaults(cfg)

	if err := validate.Struct(cfg); err != nil {
		return err
	}

	CFG = cfg
	return nil
}

func setDefaults(cfg *Config) {
	if cfg.DBType == "" {
		cfg.DBType = "sqlite"
		log.Printf("Using default DB_TYPE: %s", cfg.DBType)
	}
	if cfg.DBType == "sqlite" && cfg.SQLite.FilePath == "" {
		cfg.SQLite.FilePath = "./db.sqlite"
		log.Printf("Using default SQLITE_FILE_PATH: %s", cfg.SQLite.FilePath)
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
	if cfg.SmtpHost == "" {
		cfg.SmtpHost = "localhost"
		log.Printf("Using default SMTP_HOST: %s", cfg.SmtpHost)
	}
	if cfg.SmtpPort == "" {
		cfg.SmtpPort = "587"
		log.Printf("Using default SMTP_PORT: %s", cfg.SmtpPort)
	}
	if cfg.SmtpFrom == "" {
		cfg.SmtpFrom = "no-reply@example.com"
		log.Printf("Using default SMTP_FROM: %s", cfg.SmtpFrom)
	}
}

func getIntEnv(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
