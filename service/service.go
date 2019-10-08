package service

import (
	"database/sql"
	"fmt"

	"github.com/ugizashinje/epoc/execution"
	"github.com/ugizashinje/epoc/failures"
)

var Db *sql.DB

func Parent(ctx *execution.Context, name string) (*string, failures.SuperError) {

	fmt.Println("Parrent called")
	res, err := Repo(ctx, name)

	if err != nil {
		return nil, err
	}
	lastName := res + "ic"
	return &lastName, nil
}

func Repo(ctx *execution.Context, name string) (string, failures.SuperError) {
	fmt.Println("Child called")
	var answer string
	fmt.Println("Name is ", name)

	fmt.Println("Child done ")

	rows, err := ctx.Tx.Query(`select 'hello world' from dual`)
	if err != nil {
		return "", failures.SUBSCRIBER_DOES_NOT_EXSIS().Notify()
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&answer)
		if err != nil {
			return "", failures.SUBSCRIBER_DO_NOTIFY()
		}
	}

	return answer, nil
}
