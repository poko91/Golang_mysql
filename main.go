package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql tutorial")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/car_parking")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to MYSQL database")

	insert, err := db.Query("INSERT INTO cars(RegNum,Colour) VALUES ('MH-14-3456', 'Red')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted data into cars table")
}
