package dao

import (
	"os"

	"github.com/jinzhu/gorm"
	// mysql library
	_ "github.com/go-sql-driver/mysql"
)

// DB utility tool
type DB struct {
	Adapter, Database, Host, User, Password string
}

// Connect function
func (db DB) Connect() gorm.DB {
	con, _ := gorm.Open(db.Adapter, db.String())
	return con
}

// ConfigPath returns database.yml path
func (db DB) ConfigPath() string {
	wd, _ := os.Getwd()

	return wd + "/config/database.yml"
}

// Drive the connection parameters
// from the Db properties
func (db DB) String() string {
	// This is the connection String
	// meant for connecting to the database
	return db.User + ":" + db.Password + "@/" + db.Database
}
