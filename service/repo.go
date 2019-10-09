package service

import (
	"database/sql"

	"github.com/ugizashinje/epoc/execution"
	"github.com/ugizashinje/epoc/failures"
)

func handlePostgresError(err error) failures.SuperError {
	if err == sql.ErrNoRows {
		return failures.SUBSCRIBER_DOES_NOT_EXSIS()
	}
	return failures.GENERAL()
}
func Repo(ctx *execution.Context, name string) (string, failures.SuperError) {
	var answer string
	// rows, err := ctx.Tx.Query(`select version()`)
	rows, err := ctx.Tx.Query(`select * from dual`)
	if err != nil {
		return "", handlePostgresError(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&answer)
		if err != nil {
			return "", failures.INVALID_TYPE().Notify()
		}
	}
	return answer, nil
}
