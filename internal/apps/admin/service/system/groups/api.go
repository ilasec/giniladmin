package groups

import (
	"context"
	"giniladmin/internal/apps/admin/global"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetGroupList
// @Tags      Group API
// @Summary   获取列表（支持分页和搜索）
// @Produce   Grouplication/json
// @Param x-token header string true "token for authentication"
// @Param     page     query    int     false    "页码"    default(1)
// @Param     pageSize query    int     false    "每页数量"    default(10)
// @Param     keyword  query    string  false    "搜索关键词"
// @Success   200  {object} models.CommonResp{data=PageData{list=[]GroupModel, total=int}} "{"message":"success","status":200, "result": {"list": [], "total": 0}}"
// @Router    /api/v1/system/group [get]
func GetGroupList(c *gin.Context) (status int, message string, ret any, err error) {
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

	return DoGetGroupList(context.Background(), page, pageSize, keyword)
}

// CreateGroup
// @Tags      Group API
// @Summary   创建
// @Accept    Grouplication/json
// @Produce   Grouplication/json
// @Param x-token header string true "token for authentication"
// @Param data body GroupModel true "{}"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/group [post]
func CreateGroup(c *gin.Context) (status int, message string, ret any, err error) {
	param := GroupModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	return DoCreateGroup(context.Background(), param)
}

// GetGroup
// @Tags      Group API
// @Summary   获取
// @Produce   Grouplication/json
// @Param     id  path    int     true    "Group ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/group/{id} [get]
func GetGroup(c *gin.Context) (status int, message string, ret any, err error) {
	GroupID := c.Param("id")
	if GroupID == "" {
		status = http.StatusBadRequest
		message = "组ID不能为空"
		return
	}
	return DoGetGroup(context.Background(), utils.Atoi(GroupID))

}

// UpdateGroup
// @Tags      Group API
// @Summary   更新
// @Accept    Grouplication/json
// @Produce   Grouplication/json
// @Param     id  path    int     true    "Group ID"
// @Param     Group body    GroupModel    true    "Group info"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/group/{id} [put]  // Or PATCH for partial updates
func UpdateGroup(c *gin.Context) (status int, message string, ret any, err error) {
	param := GroupModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}

	GroupID := c.Param("id")
	if GroupID == "" {
		status = http.StatusBadRequest
		message = "组ID不能为空"
		return
	}

	param.ID = utils.Atoi(GroupID)
	return DoUpdateGroup(context.Background(), param)
}

// DeleteGroup
// @Tags      Group API
// @Summary   删除
// @Produce   Grouplication/json
// @Param     id  path    int     true    "Group ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/group/{id} [delete]
func DeleteGroup(c *gin.Context) (status int, message string, ret any, err error) {
	GroupID := c.Param("id")
	if GroupID == "" {
		status = http.StatusBadRequest
		message = "组ID不能为空"
		return
	}

	return DoDeleteGroup(context.Background(), utils.Atoi(GroupID))

}
