package association

import (
	"context"
	"giniladmin/internal/apps/admin/global"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAssociationList
// @Tags      Association API
// @Summary   获取列表（支持分页和搜索）
// @Produce   Associationlication/json
// @Param x-token header string true "token for authentication"
// @Param     page     query    int     false    "页码"    default(1)
// @Param     pageSize query    int     false    "每页数量"    default(10)
// @Param     keyword  query    string  false    "搜索关键词"
// @Success   200  {object} models.CommonResp{data=PageData{list=[]AssociationModel, total=int}} "{"message":"success","status":200, "result": {"list": [], "total": 0}}"
// @Router    /api/v1/system/association [get]
func GetAssociationList(c *gin.Context) (status int, message string, ret any, err error) {
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

	return DoGetAssociationList(context.Background(), page, pageSize, keyword)
}

// CreateAssociation
// @Tags      Association API
// @Summary   创建
// @Accept    Associationlication/json
// @Produce   Associationlication/json
// @Param x-token header string true "token for authentication"
// @Param data body AssociationModel true "{}"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/association [post]
func CreateAssociation(c *gin.Context) (status int, message string, ret any, err error) {
	param := AssociationModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	return DoCreateAssociation(context.Background(), param)
}

// GetAssociation
// @Tags      Association API
// @Summary   获取
// @Produce   Associationlication/json
// @Param     id  path    int     true    "Association ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/association/{id} [get]
func GetAssociation(c *gin.Context) (status int, message string, ret any, err error) {
	AssociationID := c.Param("id")
	if AssociationID == "" {
		status = http.StatusBadRequest
		message = "ID不能为空"
		return
	}
	return DoGetAssociation(context.Background(), utils.Atoi(AssociationID))

}

// UpdateAssociation
// @Tags      Association API
// @Summary   更新
// @Accept    Associationlication/json
// @Produce   Associationlication/json
// @Param     id  path    int     true    "Association ID"
// @Param     Association body    AssociationModel    true    "Association info"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/association/{id} [put]  // Or PATCH for partial updates
func UpdateAssociation(c *gin.Context) (status int, message string, ret any, err error) {
	param := AssociationModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}

	AssociationID := c.Param("id")
	if AssociationID == "" {
		status = http.StatusBadRequest
		message = "ID不能为空"
		return
	}

	param.ID = utils.Atoi(AssociationID)
	return DoUpdateAssociation(context.Background(), param)
}

// DeleteAssociation
// @Tags      Association API
// @Summary   删除
// @Produce   Associationlication/json
// @Param     id  path    int     true    "Association ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/association/{id} [delete]
func DeleteAssociation(c *gin.Context) (status int, message string, ret any, err error) {
	AssociationID := c.Param("id")
	if AssociationID == "" {
		status = http.StatusBadRequest
		message = "ID不能为空"
		return
	}

	return DoDeleteAssociation(context.Background(), utils.Atoi(AssociationID))

}
