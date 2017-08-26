package main

import (
	"fmt"

	mysql "github.com/go-sql-driver/mysql"

	"github.com/sightHearing/config"
	"github.com/sightHearing/database"
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
	sqlCred := mysql.Config{
		User:   awsCred.UserName,
		Passwd: awsCred.Password,
		Net:    "tcp",
		Addr:   awsCred.Endpoint + ":" + awsCred.Port,
		DBName: "sighthearingdb",
	}
	fmt.Println(sqlCred)
	// db, err := sql.Open("mysql", cred.FormatDSN())
	// if err != nil {
	// 	fmt.Println("sql open error:", err.Error())
	// 	panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	// }
	// defer db.Close()
	db := database.Database{}
	err := db.ConnectDb(sqlCred)
	if err != nil {
		panic("Could not connect to db " + err.Error())
	}

	// // Open doesn't open a connection. Validate DSN data:
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println("Ping error:", err.Error())
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	err = db.PingConnection()
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	//query a db for rows
	// index := 0
	// statement, err := db.Prepare("SELECT * FROM user")
	// if err != nil {
	// 	fmt.Println("prepare error:", err.Error())
	// }
	// rows, err := statement.Query()
	// // rows, err := db.Query("SELECT * FROM user", index)
	// if err != nil {
	// 	fmt.Println("Query error:", err.Error())
	// 	panic(err.Error())
	// }
	// defer rows.Close()
	statement := "SELECT * FROM user"
	rows, err := db.QueryDB(statement)
	if err != nil {
		panic("Query not working: " + err.Error())
	}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.index, &u.email, &u.password); err != nil {
			fmt.Println("Scan error", err.Error())
		}
		fmt.Printf("User name: %s password: %s\n", u.email, u.password)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Row error:", err.Error())
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
