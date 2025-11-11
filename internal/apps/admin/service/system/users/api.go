package users

import (
	"context"
	"giniladmin/internal/apps/admin/global"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList
// @Tags      User API
// @Summary   获取用户列表（支持分页和搜索）
// @Produce   application/json
// @Param x-token header string true "token for authentication"
// @Param     page     query    int     false    "页码"    default(1)
// @Param     pageSize query    int     false    "每页数量"    default(10)
// @Param     keyword  query    string  false    "搜索关键词"
// @Success   200  {object} models.CommonResp{data=PageData{list=[]UserModel, total=int}} "{"message":"success","status":200, "result": {"list": [], "total": 0}}"
// @Router    /api/v1/system/user [get]
func GetUserList(c *gin.Context) (status int, message string, ret any, err error) {
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

	return DoGetUserList(context.Background(), page, pageSize, keyword)
}

// CreateUser
// @Tags      User API
// @Summary   创建账号
// @Accept    application/json
// @Produce   application/json
// @Param x-token header string true "token for authentication"
// @Param data body User true "{}"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/user [post]
func CreateUser(c *gin.Context) (status int, message string, ret any, err error) {
	param := User{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}
	// 只保留必要字段
	newUser := UserModel{
		Username: param.Username,
		Password: param.Password, // 存储时需要加密
	}

	return DoCreateUser(context.Background(), newUser)
}

// GetUser
// @Tags      User API
// @Summary   获取指定ID用户
// @Produce   application/json
// @Param     id  path    int     true    "User ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/user/{id} [get]
func GetUser(c *gin.Context) (status int, message string, ret any, err error) {
	userID := c.Param("id")
	if userID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}
	return DoGetUser(context.Background(), utils.Atoi(userID))

}

// UpdateUser
// @Tags      User API
// @Summary   更新用户信息
// @Accept    application/json
// @Produce   application/json
// @Param     id  path    int     true    "User ID"
// @Param     user body    UserModel    true    "user info"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/user/{id} [put]  // Or PATCH for partial updates
func UpdateUser(c *gin.Context) (status int, message string, ret any, err error) {
	param := UserModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}

	userID := c.Param("id")
	if userID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}

	param.ID = utils.Atoi(userID)
	return DoUpdateUser(context.Background(), param)
}

// DeleteUser
// @Tags      User API
// @Summary   删除用户
// @Produce   application/json
// @Param     id  path    int     true    "User ID"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/user/{id} [delete]
func DeleteUser(c *gin.Context) (status int, message string, ret any, err error) {
	userID := c.Param("id")
	if userID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}

	return DoDeleteUser(context.Background(), utils.Atoi(userID))

}

// ChangePassword
// @Tags      User API
// @Summary   修改密码
// @Accept    application/json
// @Produce   application/json
// @Param     id  path    int     true    "User ID" // Or consider a separate endpoint like /api/v1/system/user/{id}/password
// @Param     password body      map[string]string  true    "新密码"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  ""
// @Router    /api/v1/system/user/{id}/password [put] // Or PATCH if only password is being updated
func ChangePassword(c *gin.Context) (status int, message string, ret any, err error) {
	// 获取 URL 参数中的用户 ID
	userID := c.Param("id")
	if userID == "" {
		status = http.StatusBadRequest
		message = "用户ID不能为空"
		return
	}

	// 解析请求 body 中的 password
	var req struct {
		Password string `json:"password"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		status = http.StatusBadRequest
		message = "请求参数错误"
		return
	}

	return DoChangePassword(context.Background(), utils.Atoi(userID), req.Password)
}
