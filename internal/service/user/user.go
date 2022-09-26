package user

import (
	"database/sql"
	"fmt"
	"todolist/configs/mysql"

	R "todolist/api/restful"
)

const (
	userList     string = "select USER_ID, USER_NAME, USER_AGE from user"
	userListById string = "select USER_ID, USER_NAME, USER_AGE from user where USER_ID = ?"
	userInsert   string = "INSERT INTO user (`USER_NAME`, `USER_AGE`) VALUES (?, ?)"
	userDelete   string = "DELETE FROM user WHERE USER_ID = ?"
	userUpdate   string = "UPDATE user SET `USER_AGE` = ? WHERE (`USER_ID` = ?)"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetAll() *R.Result {

	var (
		user User
		list []User
	)

	db, err := mysql.GetConnection()

	if err != nil {
		fmt.Printf("連線失敗,err:%v\n", err)
		return R.Fail(err.Error())

	}

	res, err := db.Query(userList)

	if err != nil {
		return R.Fail(err.Error())
	} else {
		for res.Next() {
			res.Scan(&user.Id, &user.Name, &user.Age)
			list = append(list, user)
		}
	}

	defer db.Close()

	return R.Success(list)
}

func GetById(id string) *R.Result {

	var (
		db   *sql.DB
		err  error
		user User
		list []User
	)

	db, err = mysql.GetConnection()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())
	}

	err = db.QueryRow(userListById, id).Scan(&user.Id, &user.Name, &user.Age)

	if err != nil {
		return R.Fail(err.Error())
	} else {
		list = append(list, user)
	}

	defer db.Close()

	return R.Success(list)
}

func Insert(user User) *R.Result {

	var (
		db  *sql.DB
		err error
	)

	db, err = mysql.GetConnection()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())
	}

	tx, err := db.Begin()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())
	}

	stmt, err := tx.Prepare(userInsert)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())
	}

	res, err := stmt.Exec(user.Name, user.Age)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())
	}

	tx.Commit()
	i, err := res.LastInsertId()

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())

	}

	defer db.Close()
	return R.Success(i != 0)
}

func Delete(id string) *R.Result {

	var (
		db  *sql.DB
		err error
	)

	db, err = mysql.GetConnection()
	if err != nil {
		return R.Fail(err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		return R.Fail(err.Error())
	}

	stmt, err := tx.Prepare(userDelete)
	if err != nil {
		return R.Fail(err.Error())
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return R.Fail(err.Error())
	}

	tx.Commit()

	i, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())

	}

	defer db.Close()
	return R.Success(i)
}

func Update(user User) *R.Result {

	var (
		db  *sql.DB
		err error
	)

	db, err = mysql.GetConnection()
	if err != nil {
		return R.Fail(err.Error())
	}

	tx, err := db.Begin()
	if err != nil {
		return R.Fail(err.Error())
	}

	stmt, err := tx.Prepare(userUpdate)
	if err != nil {
		return R.Fail(err.Error())
	}

	res, err := stmt.Exec(user.Age, user.Id)
	if err != nil {
		return R.Fail(err.Error())
	}

	tx.Commit()

	i, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return R.Fail(err.Error())

	}

	defer db.Close()
	return R.Success(i != 0)
}
