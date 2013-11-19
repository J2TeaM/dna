package sqlpg

import (
	"database/sql"
	. "dna"
	"fmt"
	_ "github.com/lib/pq"
)

// Config returns config type
// Valid values for SSLMode are:
// 		* disable - No SSL
// 		* require - Always SSL (skip verification)
// 		* verify-full - Always SSL (require verification)
type Config struct {
	Username String // The user to sign in as
	Password String // The user's password
	Host     String // The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
	Post     Int    // The port to bind to. (default is 5432)
	Database String // The name of the database to connect to
	SSLMode  String
}

var DefaultConfig *Config = &Config{
	Username: "daonguyenanbinh",
	Password: "",
	Host:     "127.0.0.1",
	Post:     5432,
	Database: "daonguyenanbinh",
	SSLMode:  "disable",
}

// Connect returns database of connected server
func Connect(cf *Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cf.Username, cf.Password, cf.Host, cf.Post, cf.Database, cf.SSLMode)
	db, err := sql.Open("postgres", connectionString)
	return db, err
}
