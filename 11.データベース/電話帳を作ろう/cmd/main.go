package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sMakeTable = `CREATE TABLE IF NOT EXISTS phonebook (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		number TEXT NOT NULL
		)`
)

// PhoneBook is pb
type PhoneBook struct {
	ID     int64
	Name   string
	Number string
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	datafile := dir + "/sqlite/mydata.db"
	dbconnection, err := sql.Open("sqlite3", datafile)
	if err != nil {
		panic(err.Error())
	}
	defer dbconnection.Close()
	if err := dbconnection.Ping(); err != nil { // sql.Openでは接続しないので事前疎通確認。
		panic(err.Error())
	}
	if _, err := dbconnection.Exec(sMakeTable); err != nil {
		panic(err.Error())
	}
	// ここまで起動準備
	for {
		rows, err := dbconnection.Query("SELECT * FROM phonebook")
		if err != nil {
			panic(err.Error())
		}
		for rows.Next() {
			var pb PhoneBook
			if err := rows.Scan(&pb.ID, &pb.Name, &pb.Number); err != nil {
				panic(err.Error())
			}
			fmt.Printf("%v\n", pb)
		}
		if err := rows.Err(); err != nil {
			panic(err.Error())
		}

		fmt.Printf("-----------------------------------------\n")
		if ans := ReadStdin("1.電話帳に登録する\n2.電話帳を閉じる"); ans == "1" {
			name := ReadStdin("名前：")
			telNumber := ReadStdin("電話番号：") // バリデーション手抜き
			result, err := dbconnection.Exec("INSERT INTO phonebook ( name, number ) VALUES (?,?)", name, telNumber)
			if err != nil {
				panic(err.Error())
			}
			id, err := result.LastInsertId()
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("電話帳に登録しました. ID:%d, Name:%s, Number:%s\n", id, name, telNumber)
			fmt.Printf("-----------------------------------------\n")
		} else {
			break
		}

	}
}

// ReadStdin read from stdin
func ReadStdin(msg string) string {
	fmt.Printf("%s\n> ", msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
