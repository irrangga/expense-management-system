package httputil

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPathParamInt64 safely parses a path param as int64.
func GetPathParamInt64(ctx *gin.Context, key string) (int64, error) {
	valStr := ctx.Param(key)
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid path param %q", key)
	}
	return val, nil
}

// GetQueryParamInt safely parses a query param as int.
func GetQueryParamInt(ctx *gin.Context, key string) (int, error) {
	valStr := ctx.Query(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("invalid query param %q", key)
	}
	return val, nil
}
