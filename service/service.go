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
