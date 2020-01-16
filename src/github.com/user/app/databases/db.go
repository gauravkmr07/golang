package databases

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DB exporting
var DB *sql.DB

// DBConnect -
func DBConnect() string {

	dbUser := os.Getenv("DB_USER_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

// Connect :: will be used to connect to the datbase
func Connect() {
	dbDriver := os.Getenv("DB_DRIVER")

	var err error
	connection := DBConnect()

	DB, err = sql.Open(dbDriver, connection)

	if err != nil {
		panic("Error occured while connecting to database : " + err.Error())
	}

	err = PingDB()

	if err != nil {
		panic("Error while connecting with database :" + err.Error())
	}

	log.Println("####### Database connected successfully ########")

}

// PingDB -
func PingDB() error {

	err := DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
