package service

import (
	"github.com/ugizashinje/epoc/execution"
	"github.com/ugizashinje/epoc/failures"
)

func Repo(ctx *execution.Context, name string) (string, failures.SuperError) {
	var answer string
	// rows, err := ctx.Tx.Query(`select version()`)
	rows, err := ctx.Tx.Query(`select * from dual`)
	if err != nil {
		return "", failures.SUBSCRIBER_DOES_NOT_EXSIS()
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&answer)
		if err != nil {
			return "", failures.SUBSCRIPTION_ABOUT_TO_EXPIRE().Notify()
		}
	}

	return answer, nil
}
