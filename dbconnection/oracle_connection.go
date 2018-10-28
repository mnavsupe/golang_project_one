package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("GO Mysql Connection")

	//db, err := sql.Open("mysql", "root:root@tcp(10.97.108.211:3306)/adb")
	db, err := sql.Open("goracle", "gq1/g1@10.102.136.50@oracle2")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Connected Successfully")

}
