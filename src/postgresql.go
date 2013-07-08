package main

import _ "github.com/bmizerany/pq"
import "database/sql"
import "time"
import "fmt"
import "os"

func errHandler(err error) {
	if err != nil {
		fmt.Printf("error:", err)
		os.Exit(1)
	}
}

func main() {
	conf := "dbname=gobyexample sslmode=disable"
	db, openErr := sql.Open("postgres", conf)
	errHandler(openErr)

	defer db.Close()

	fmt.Println(db)

	dropRep, dropErr := db.Exec(
		`DROP TABLE IF EXISTS items`)
	errHandler(dropErr)
	fmt.Println(dropRep)

	createRep, createErr := db.Exec(
		`CREATE TABLE items
		 (a int, b float, c boolean,
		d text, e timestamp with time zone)`)
	errHandler(createErr)
	fmt.Println(createRep)

	insertRep, insertErr := db.Exec(
		`INSERT INTO items VALUES
	(1, 2.0, false,
		'string', '2000-01-01T01:02:03Z')`)
	errHandler(insertErr)
	fmt.Println(insertRep)

	timeFmt := time.RFC3339
	t1, _ := time.Parse(timeFmt, "2000-04-08T03:02:01Z")
	t2, _ := time.Parse(timeFmt, "2007-03-02T10:15:45Z")
	minsertRep, minsertErr := db.Exec(
		`INSERT INTO items (a, b, c, d, e)
		 VALUES ($1, $2, $3, $4, $5),
		($6, $7, $8, $9, $10)`,
		3, 7.0, true, "more", t1,
		5, 1.0, false, "less", t2)
	errHandler(minsertErr)
  num, _ := minsertRep.RowsAffected()
	fmt.Println(minsertRep)
	fmt.Println(num)
}
