package sqlpg

import (
	. "dna"
	"dna/cfg"
)

// SQLConfig contains relevant fields to connect to database.
// It returns config type
// Valid values for SSLMode are:
// 		* disable - No SSL
// 		* require - Always SSL (skip verification)
// 		* verify-full - Always SSL (require verification)
type SQLConfig struct {
	Username String // The user to sign in as
	Password String // The user's password
	Host     String // The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
	Post     Int    // The port to bind to. (default is 5432)
	Database String // The name of the database to connect to
	SSLMode  String
}

// NewSQLConfig loads sql config from a file and returns new SQLConfig.
func NewSQLConfig(path String) *SQLConfig {
	cfg, err := cfg.LoadConfigFile(path)
	PanicError(err)
	db, err := cfg.GetSection("database")
	PanicError(err)
	return &SQLConfig{
		Username: db["username"],
		Password: db["password"],
		Host:     db["host"],
		Post:     db["port"].ToInt(),
		Database: db["db"],
		SSLMode:  db["sslmode"],
	}
}
