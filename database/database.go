package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB


func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_ip := os.Getenv("DB_IP")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", db_username, db_password, db_ip, db_name)
	db, _ = sql.Open("postgres", connStr)
}

func GetDbConnection() *sql.DB {
	return db
}

func IsExistShortURL(s string) bool {
	table_name := os.Getenv("TABLE_NAME")
	sql := fmt.Sprintf("SELECT (url) FROM %s WHERE short_url='%s'", table_name, s)
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalln(err)
	}
	isExist := rows.Next()
	return isExist
}