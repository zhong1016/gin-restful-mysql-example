package main

import (
	"database/sql"
	"fmt"

	"todolist/configs/mysql"
)

type USER struct {
	Id   int
	Name string
	Age  int
}

func GetById(db *sql.DB, id int) {

	s := "select * from user where USER_ID = ?"

	var u USER

	err := db.QueryRow(s, id).Scan(&u.Id, &u.Name, &u.Age)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("user=: %+v\n", u)
	}
}

func GetAll(db *sql.DB) {

	s := "select * from user"
	r, err := db.Query(s)
	var u USER
	// defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for r.Next() {
			r.Scan(&u.Id, &u.Name, &u.Age)
			fmt.Printf("u: %+v\n", u)
		}
	}
}
func Insert(db *sql.DB) {
	s := "insert into user (USER_ID, USER_NAME, USER_AGE) values(?,?,?)"
	r, err := db.Exec(s, 4, "錦木千束", 17)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	db, err := mysql.GetConnection()
	if err != nil {
		fmt.Printf("連線失敗,err:%v\n", err)
		return
	} else {
		fmt.Println("連線成功!")
	}

	// GetById(db, 1)
	// GetAll(db)
	// Insert(db)

	defer db.Close()
}
