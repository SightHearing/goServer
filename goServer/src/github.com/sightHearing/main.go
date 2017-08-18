package main

import (
	"database/sql"
	"fmt"

	mysql "github.com/go-sql-driver/mysql"

	"github.com/sightHearing/config"
)

type user struct {
	index    int
	email    string
	password string
}

func main() {
	config := config.Config{}
	config.InitConfig()
	awsCred := config.AwsSql
	//id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	// cred := awsCred.UserName + ":" + awsCred.Password + "@tcp(" + awsCred.Endpoint + ":" + awsCred.Port + ")/" + "sighthearingdb"
	cred := mysql.Config{
		User:   awsCred.UserName,
		Passwd: awsCred.Password,
		Net:    "tcp",
		Addr:   awsCred.Endpoint + ":" + awsCred.Port,
		DBName: "sighthearingdb",
	}
	fmt.Println(cred)
	db, err := sql.Open("mysql", cred.FormatDSN())
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping error:", err.Error())
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// db, err := sql.Open("mysql", cred)
	// if err != nil {
	// 	println("error:", err.Error())
	// }
	//
	// rows, err := db.Query("SELECT * FROM user")
	// for rows.Next() {
	// 	var u user
	// 	err = rows.Scan(&u)
	// 	fmt.Println(u)
	// }
}
