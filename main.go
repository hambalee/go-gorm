package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Cover struct {
	Id   int
	Name string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlserver", "sqlserver://sa:yourStrong(!)Password@localhost:1433?database=gorm&connection+timeout=30")
	if err != nil {
		panic(err)
	}
	covers, err := GetCovers()
	if err != nil {
		fmt.Println((err))
		return
	}

	for _, cover := range covers {
		fmt.Println(cover)
	}
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}
	// if err = db.Ping(); err != nil {
	// 	panic(err)
	// }

	query := "select id, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}
		covers = append(covers, cover)
	}

	return covers, nil
}
