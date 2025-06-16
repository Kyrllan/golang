package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8parseTime=True&Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(db)
}
