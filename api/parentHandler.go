package api

import (
	"github.com/ugizashinje/epoc/execution"
	"github.com/ugizashinje/epoc/failures"
	"github.com/ugizashinje/epoc/service"
)

func parentHandler(ctx *execution.Context) (interface{}, failures.SuperError) {
	request := ctx.Gin.Param("name")
	response, err := service.Parent(ctx, request)
	return &response, err
}
