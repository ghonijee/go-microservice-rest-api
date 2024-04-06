package database

import (
	"database/sql"
	"fmt"
	"go-microservice/user-service/internal/common"
	"go-microservice/user-service/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	config := config.Config
	username := config.GetString("database.username")
	password := config.GetString("database.password")
	host := config.GetString("database.host")
	port := config.GetInt("database.port")
	database := config.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)

	common.PanicIfError(err)

	DB = db
}
