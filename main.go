package main

import (
	"4hoeschele/go_dnd_project.git438/cmd/api"
	"4hoeschele/go_dnd_project.git438/db"
	"fmt"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
		err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Server is successfully running on port:", os.Getenv("PORT"))
}

func main() {
	dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbAddr := os.Getenv("DB_ADDR")
    dbName := os.Getenv("DB_NAME")

    db, err := db.NewMySQLStorage(mysql.Config{
        User:                 dbUser,
        Passwd:               dbPass,
        Addr:                 dbAddr,
        DBName:               dbName,
        Net:                  "tcp",
        AllowNativePasswords: true,
        ParseTime:            true,
    })
	if err != nil {
		fmt.Println(err)
		fmt.Println(db)
	}
    fmt.Println("Database Username: ", dbUser)
	port := os.Getenv("PORT")
	router := api.InitRoutes()
	// creates a new server instans
	server := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	server.ListenAndServe()
}
