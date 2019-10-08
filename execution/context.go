package execution

import (
	"database/sql"

	ginPkg "github.com/gin-gonic/gin"
)

type Context struct {
	Gin *ginPkg.Context
	Tx  *sql.Tx
}
