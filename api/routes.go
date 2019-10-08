package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ugizashinje/epoc/execution"
	"github.com/ugizashinje/epoc/failures"
	"github.com/ugizashinje/epoc/service"
)

func InitRoutes() error {
	r := gin.Default()

	r.GET("/parent/:name", wrap(parentHandler))
	r.Run(":8080")

	return nil
}

func wrap(handler func(c *execution.Context) (interface{}, failures.SuperError)) func(gc *gin.Context) {
	return func(c *gin.Context) {
		Tx, _ := service.Db.Begin()

		ctx := &execution.Context{
			Gin: c,
			Tx:  Tx,
		}
		result, err := handler(ctx)

		if err != nil {
			superError := err.(failures.SuperError)
			ctx.Tx.Rollback()
			ctx.Gin.JSON(superError.Status(), superError)
			ctx.Gin.Abort()
			return
		}

		if err := ctx.Tx.Commit(); err != nil {
			ctx.Gin.JSON(http.StatusInternalServerError, err)
			ctx.Gin.Abort()
			return
		}
		ctx.Gin.JSON(http.StatusOK, result)

	}
}
