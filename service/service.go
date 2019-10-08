package service

import (
	"database/sql"

	"github.com/ugizashinje/epoc/execution"
	"github.com/ugizashinje/epoc/failures"
)

var Db *sql.DB

func Parent(ctx *execution.Context, name string) (*string, failures.SuperError) {

	res, err := Repo(ctx, name)

	if err != nil {
		return nil, err
	}
	lastName := res + "ic"
	return &lastName, nil
}
