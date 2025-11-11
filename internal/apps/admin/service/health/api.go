package health

import (
	"context"
	"github.com/gin-gonic/gin"
)

// Healthz
// @Tags Monitor API
// @Summary health check
// @Description health check
// @Produce json
// @Param Authorization header   string true "Bearer xxx"
// @Success 200 {object} models.CommonResp "{"message":"success","status":200}"
// @Router /api/v1/healthCheck/healthz [get]
func Healthz(c *gin.Context) (status int, message string, ret any, err error) {
	DoHealth(context.Background())
	return
}
