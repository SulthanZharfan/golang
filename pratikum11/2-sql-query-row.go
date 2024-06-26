package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	// Menggunakan pengguna root, tanpa kata sandi, ke database golang pada localhost di port 3306
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		return nil, err
	}

	// Memeriksa koneksi dengan melakukan ping ke database
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "E001"
	err = db.
		QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

func main() {
	sqlQueryRow()
}
