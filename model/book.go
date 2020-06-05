package model

import (
	"awesomeProject/base"
	"fmt"
)

func GetBook() {
	//age := 1
	rows, err := db.DB().Query("SELECT id,name FROM public.book")
	println(err)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Println("uid | name | department | created ")
		fmt.Printf("%3v | %8v \n", id, name)
	}

	defer rows.Close()
}

func GetBook2() {
	//age := 1
	rows, err := db.NewDB().Query("SELECT id,name FROM public.book")
	println(err)
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		fmt.Println("uid | name | department | created ")
		fmt.Printf("%3v | %8v \n", id, name)
	}

	rows.Close()
}