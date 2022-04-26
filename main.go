package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	start := time.Now()

	fmt.Println("START")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://jqnqjfrd:%s@%s/jqnqjfrd",
		os.Getenv("ELEPH_PASS"),
		os.Getenv("PSQL_HOST")))
	if err != nil {
		log.Fatalln(err)
	}
	//tx := db.MustBegin()

	cats := make([]Cat, 0)
	//var cat Cat

	err = db.Ping()
	if err != nil {
		return
	}
	err = db.Select(&cats, "SELECT * FROM cats")
	if err != nil {
		return
	}
	//rows, err := db.Query("SELECT * FROM cats")
	//for rows.Next() {
	//	err := rows.Scan(&cat.Name, &cat.Age, &cat.Color)
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//	cats = append(cats, cat)
	//}

	defer db.Close()
	fmt.Println(cats)

	fmt.Println(time.Now().Sub(start))
}
