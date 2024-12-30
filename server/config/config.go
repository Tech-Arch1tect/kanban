package config

import (
	"log"
	"os"
	"strconv"
)

var (
	CFG *Config
)

type Config struct {
	DBType       string
	MySQL        MySQLConfig
	SQLite       SQLiteConfig
	CookieSecret string
	SessionName  string
	AllowOrigin  string
	AppName      string
	SmtpHost     string
	SmtpPort     string
	SmtpUser     string
	SmtpPassword string
	SmtpAuth     string
	SmtpFrom     string
	SmtpNoTLS    bool
	RateLimit    RateLimitConfig
}

type RateLimitConfig struct {
	Enabled             bool
	Limit               int
	Window              int
	LoginLimit          int
	LoginWindow         int
	PasswordResetLimit  int
	PasswordResetWindow int
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
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
	FilePath string
}

func (c SQLiteConfig) GetFilePath() string {
	return c.FilePath
}

func LoadConfig() {
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
			Limit:               getIntEnv("RATE_LIMIT_LIMIT", 100),                // default 100 requests per IP
			Window:              getIntEnv("RATE_LIMIT_WINDOW", 1),                 // default 1 minute
			LoginLimit:          getIntEnv("RATE_LIMIT_LOGIN_LIMIT", 20),           // default 20 requests per IP
			LoginWindow:         getIntEnv("RATE_LIMIT_LOGIN_WINDOW", 10),          // default 10 minute
			PasswordResetLimit:  getIntEnv("RATE_LIMIT_PASSWORD_RESET_LIMIT", 5),   // default 5 requests per IP
			PasswordResetWindow: getIntEnv("RATE_LIMIT_PASSWORD_RESET_WINDOW", 10), // default 10 minute
		},
	}

	if cfg.DBType == "" {
		cfg.DBType = "sqlite"
	}

	if cfg.DBType == "sqlite" {
		if cfg.SQLite.FilePath == "" {
			log.Println("Warning: SQLite file path is not set. Using default value: ./db.sqlite. Please set this with the environment variable SQLITE_FILE_PATH")
			cfg.SQLite.FilePath = "./db.sqlite"
		}
	}

	if cfg.CookieSecret == "" {
		log.Println("Warning: Cookie secret is not set. Using default value: secret. Please set this with the environment variable COOKIE_SECRET")
		cfg.CookieSecret = "secret"
	}

	if cfg.SessionName == "" {
		log.Println("Warning: Session name is not set. Using default value: mysession. Please set this with the environment variable SESSION_NAME")
		cfg.SessionName = "mysession"
	}

	if cfg.AllowOrigin == "" {
		log.Println("Warning: Allow origin is not set. Using default value: http://localhost:3001. Please set this with the environment variable ALLOW_ORIGIN")
		cfg.AllowOrigin = "http://localhost:3001"
	}

	if cfg.AppName == "" {
		log.Println("Warning: App name is not set. Using default value: App Name. Please set this with the environment variable APP_NAME")
		cfg.AppName = "App Name"
	}

	if cfg.SmtpHost == "" {
		log.Println("Warning: SMTP host is not set. Using default value: localhost. Please set this with the environment variable SMTP_HOST")
		cfg.SmtpHost = "localhost"
	}

	if cfg.SmtpPort == "" {
		log.Println("Warning: SMTP port is not set. Using default value: 587. Please set this with the environment variable SMTP_PORT")
		cfg.SmtpPort = "587"
	}

	if cfg.SmtpFrom == "" {
		log.Println("Warning: SMTP from is not set. Using default value: no-reply@example.com. Please set this with the environment variable SMTP_FROM")
		cfg.SmtpFrom = "no-reply@example.com"
	}

	CFG = cfg
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
