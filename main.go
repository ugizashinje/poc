package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/ugizashinje/epoc/failures"
)

var connStr string
var db *sql.DB

func Parent(ctx context.Context, name string) (string, failures.SuperError) {

	fmt.Println("Parrent called")
	res, err := Repo(ctx, name)

	if err != nil {
		return "", err
	}
	fmt.Println("Parent done")
	return res + "ic", nil
}

func Repo(ctx context.Context, name string) (string, failures.SuperError) {
	fmt.Println("Child called")
	var answer string
	fmt.Println("Name is ", name)

	fmt.Println("Child done ")

	rows, err := db.Query(`select 'hello world' from dual`)
	if err != nil {
		return "", failures.SUBSCRIBER_DOES_NOT_EXSIS().Notify()
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&answer)
		if err != nil {
			return "", failures.SUBSCRIBER_DO_NOT_NOTIFY()
		}
	}

	return answer, nil

}

func main() {
	fmt.Println("Start")
	var err error
	connStr = "user=postgres password=postgres dbname=postgres host=localhost sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		// log.Fatal(err)
	}

	res, err := Parent(context.Background(), "Nikol")

	if err != nil {
		fmt.Println("error returned ", err.Error())
	} else {
		fmt.Println("END ", res)
	}
}
