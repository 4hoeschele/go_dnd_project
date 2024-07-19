package db

import (
	"fmt"

	// Import for side-effects only
	_ "github.com/go-sql-driver/mysql"

    "github.com/4hoeschele/go_dnd_project/ent"
)

// GetConnectionString returns a string representing the connection to a MySQL database,
// constructed using the fields of the Config struct.
func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", "root", "example", "localhost", 3306, "dnd_generator")
}

// SetupEntDatabaseConnection connects the ent ORM to a MySQL database using the given ConfigInterface.
// It returns an ent.Client object and an error.
func SetupEntDatabaseConnection() (*ent.Client, error) {
	return ent.Open("mysql", GetConnectionString())
}



