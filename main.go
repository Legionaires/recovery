package main

import (
	"fmt"
	"os"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
)

func main() {
	db := mysql.New("tcp", "", "192.168.42.82:3306", "anyone", "", "phpnuke")
	err := db.Connect()
	if err != nil {
		panic(err)
	}

rows, _, err := db.Query("select * from nuke_bbposts")
if err != nil {
	panic(err)
}

for _, row := range rows {
	for _, col := range row {
		if col == nil {

		} else {
			colval := col.([]byte)
			os.Stdout.Write(colval)
			fmt.Printf("\n")
		}
	}
}



	fmt.Printf("It's alive!")
}
