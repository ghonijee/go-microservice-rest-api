package database

import (
	"context"
	"database/sql"
	"fmt"
	"go-microservice/user-service/internal/common"
	"go-microservice/user-service/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	Ctx context.Context
)

func Connect() {
	config := config.Config
	username := config.GetString("database.user")
	password := config.GetString("database.password")
	host := config.GetString("database.host")
	port := config.GetInt("database.port")
	database := config.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)

	common.PanicIfError(err)

	DB = db

	Ctx = context.Background()
}
