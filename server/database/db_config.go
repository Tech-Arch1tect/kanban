package database

type DBConfig interface {
	GetDBType() string
	GetSQLiteConfig() SQLiteConfig
	GetMySQLConfig() MySQLConfig
}

type SQLiteConfig interface {
	GetFilePath() string
}

type MySQLConfig interface {
	GetUser() string
	GetPassword() string
	GetHost() string
	GetPort() string
	GetDatabase() string
}
