package handler

import (
	"giniladmin/internal/errorcode"
	"giniladmin/internal/routers/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler func(*gin.Context) (int, string, any, error)

func fill(status int, message string, ret any, err error) (int, string, any) {
	if err != nil {
		if status == 0 {
			status = int(errorcode.ERROR_CODE_FAILED)
		}
		if message == "" {
			message = err.Error()
		}
		if ret == nil {
			ret = err.Error()
		}
	} else {
		if status == 0 {
			status = int(errorcode.ERROR_CODE_SUCCESS)
		}
		if message == "" && status == 0 {
			message = "success"
		}
		if ret == nil {
			ret = ""
		}
	}

	return status, message, ret
}

func HandleFailure(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		status, message, ret, err := h(c)
		status, message, ret = fill(status, message, ret, err)
		c.JSON(http.StatusOK, resp.Result(status, message, ret))
	}
}
