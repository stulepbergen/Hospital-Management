package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root:cf5sysinektg,thuty@tcp(127.0.0.1:3306)/Hospitalmanagement")

	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(`USE HospitalManagement`)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Success!")

	return db
}
