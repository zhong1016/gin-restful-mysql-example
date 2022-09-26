package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "todoList"
)

type MySQL struct {
	DB *sql.DB
}

func GetConnection() (*sql.DB, error) {
	dsn :=
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", userName, password, ip, port, dbName,
		)
	var (
		mysql MySQL
		err   error
	)
	mysql.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// test ping mysql
	err = mysql.DB.Ping()
	if err != nil {
		return nil, err
	} else {
		fmt.Println("連線成功!")
	}

	mysql.DB.SetConnMaxLifetime(time.Minute * 3)
	mysql.DB.SetMaxOpenConns(100)
	mysql.DB.SetMaxIdleConns(10)
	return mysql.DB, nil
}
