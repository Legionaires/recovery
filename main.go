package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "anyone@tcp([192.168.42.82]:3306)/phpnuke")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("It's alive!")

	var (
		id int
		topic int
		forum int
		poster int
		time int
		ip string
		name string
		bbcode int
		html int
		smilies int
		sig int
		edittime []byte
		count int
	)

	rows, err := db.Query("select * from nuke_bbposts")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(cols, "\n")
	}



	for rows.Next() {
		err := rows.Scan(&id, &topic, &forum, &poster, &time, &ip, &name, &bbcode, &html, &smilies, &sig, &edittime, &count)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", name)
	}

}
