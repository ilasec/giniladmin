package application

import (
	"context"
	"giniladmin/internal/apps/admin/global"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAppList
// @Tags      App API
// @Summary   获取列表（支持分页和搜索）
// @Produce   application/json
// @Param x-token header string true "token for authentication"
// @Param     page     query    int     false    "页码"    default(1)
// @Param     pageSize query    int     false    "每页数量"    default(10)
// @Param     keyword  query    string  false    "搜索关键词"
// @Success   200  {object} models.CommonResp{data=PageData{list=[]AppModel, total=int}} "{"message":"success","status":200, "result": {"list": [], "total": 0}}"
// @Router    /api/v1/system/app [get]
func GetAppList(c *gin.Context) (status int, message string, ret any, err error) {
	// 1. 获取查询参数
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if pageSize <= 0 {
		pageSize = global.PAGE_SIZE
	}

	keyword := c.Query("keyword")

	return DoGetAppList(context.Background(), page, pageSize, keyword)
}

// CreateApp
// @Tags      App API
// @Summary   创建
// @Accept    application/json
// @Produce   application/json
// @Param x-token header string true "token for authentication"
// @Param data body AppModel true "{}"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/app [post]
func CreateApp(c *gin.Context) (status int, message string, ret any, err error) {
	param := AppModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	return DoCreateApp(context.Background(), param)
}

// GetApp
// @Tags      App API
// @Summary   获取
// @Produce   application/json
// @Param     id  path    int     true    "App ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/app/{id} [get]
func GetApp(c *gin.Context) (status int, message string, ret any, err error) {
	AppID := c.Param("id")
	if AppID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}
	return DoGetApp(context.Background(), utils.Atoi(AppID))

}

// UpdateApp
// @Tags      App API
// @Summary   更新
// @Accept    application/json
// @Produce   application/json
// @Param     id  path    int     true    "App ID"
// @Param     App body    AppModel    true    "App info"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/app/{id} [put]  // Or PATCH for partial updates
func UpdateApp(c *gin.Context) (status int, message string, ret any, err error) {
	param := AppModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}

	AppID := c.Param("id")
	if AppID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}

	param.ID = utils.Atoi(AppID)
	return DoUpdateApp(context.Background(), param)
}

// DeleteApp
// @Tags      App API
// @Summary   删除
// @Produce   application/json
// @Param     id  path    int     true    "App ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/app/{id} [delete]
func DeleteApp(c *gin.Context) (status int, message string, ret any, err error) {
	AppID := c.Param("id")
	if AppID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}

	return DoDeleteApp(context.Background(), utils.Atoi(AppID))

}
