package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Println("err:", err)
	}

	defer db.Close()

	var books []Book

	err = db.Select(&books, "select * from books where price > ?", 50)

	if err != nil {
		fmt.Println("查询失败", err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}
