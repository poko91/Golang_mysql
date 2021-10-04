package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// define a struct for car
type Car struct {
	Car_id int    `json:"Car_id"`
	RegNum string `json:"RegNum"`
	Colour string `json:"Colour"`
}

func main() {
	fmt.Println("Go Mysql tutorial")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/car_parking")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to MYSQL database")

	results, err := db.Query("SELECT * FROM cars")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var car Car

		err = results.Scan(&car.Car_id, &car.RegNum, &car.Colour)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(car.Car_id, car.RegNum, car.Colour)
	}
}
