package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@/lil")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("It's alive!\n")
	rows, err := db.Query("select count(*) from classic_forums")
	

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		fmt.Printf("Count=%d\n",count)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}
