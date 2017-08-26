package database

import (
	"database/sql"
	"fmt"

	mysql "github.com/go-sql-driver/mysql"
)

type Database struct {
	Db *sql.DB
}

func (d *Database) ConnectDb(cred mysql.Config) (err error) {
	dbInstance, err := sql.Open("mysql", cred.FormatDSN())
	d.Db = dbInstance
	if err != nil {
		fmt.Println("sql open error:", err.Error())
		return
	}
	defer dbInstance.Close()

	return
}

func (d *Database) PingConnection() (err error) {
	// Open doesn't open a connection. Validate DSN data:
	err = d.Db.Ping()
	if err != nil {
		fmt.Println("Ping error:", err.Error())
	}
	return
}

func (d *Database) QueryDB(statement string) (rows *sql.Rows, err error) {
	prep, err := d.Db.Prepare(statement)
	if err != nil {
		fmt.Println("prepare error:", err.Error())
	}
	rows, err = prep.Query()
	// rows, err := db.Query("SELECT * FROM user", index)
	if err != nil {
		fmt.Println("Query error:", err.Error())
		panic(err.Error())
	}
	defer rows.Close()
	return
}
