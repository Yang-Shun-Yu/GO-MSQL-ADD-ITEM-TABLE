package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Hh124828526@tcp(localhost:3306)/testdb")
	if err != nil {
		fmt.Println("error validating sql.open arguments")
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	insert, err := db.Query("insert into `testdb`.`students`(`firstname`,`lastname`)values('Yang','Yu');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successful connection to database")
}
