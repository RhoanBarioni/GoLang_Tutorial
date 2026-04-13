package service

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Db(){
	_, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/escola")
	if err != nil{
		log.Fatal()
	}
}