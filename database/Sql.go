package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var sqlDB *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	dbUser := os.Getenv("SQL_DB_USER")
	dbPass := os.Getenv("SQL_DB_PASS")
	dbHost := os.Getenv("SQL_DB_HOST")
	dbName := os.Getenv("SQL_DB_NAME")
	dbPort, err := strconv.ParseInt(os.Getenv("SQL_DB_PORT"), 10, 32)
	if err != nil {
		fmt.Println("database port convert error:", err)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		dbHost, dbUser, dbPass, dbPort, dbName)

	// Create connection pool

	conn, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println("mssql database connection error", err)
	} else {
		fmt.Println("mssql database connected")
	}

	sqlDB = conn
	// defer db.Close()

}

func GetSQLDB() *gorm.DB {
	return sqlDB
}
