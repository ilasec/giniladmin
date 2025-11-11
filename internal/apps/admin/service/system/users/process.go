package users

import (
	"context"
	"errors"
	"giniladmin/pkg/utils"
	"giniladmin/pkg/utils/structutils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func Login(user *UserModel) (userInter UserModel, err error) {
	model := UserModelDB{}

	// 1. 校验用户是否已存在
	u, err := model.FindByUsername(context.Background(), repo.Db, user.Username)
	if err != nil || u == nil {
		err = errors.New("账号不存在")
		return
	}

	// 2. 校验用户类型
	if u.IsSystem != 1 {
		err = errors.New("账号未授权")
		return
	}

	// 3. 校验密码
	if ok := utils.PasswordCheckByBcrypt(user.Password, u.Password); !ok {
		err = errors.New("密码错误")
		return
	}

	userInter.Username = u.Username
	userInter.ID = u.ID
	userInter.Enable = u.Enable
	return
}

func DoCreateUser(ctx context.Context, user UserModel) (status int, message string, ret any, err error) {
	model := UserModelDB{}

	// 1. 校验用户是否已存在
	u, err := model.FindByUsername(ctx, repo.Db, user.Username)
	if err != nil {
		return
	}

	if u != nil {
		status = http.StatusConflict
		message = "用户已存在"
		return
	}

	// 2. 加密密码
	p, err := utils.PasswordByBcrypt(user.Password)
	if err != nil {
		status = http.StatusInternalServerError
		message = "密码加密失败"
		return
	}
	user.Password = p

	// 3. 填充其他字段
	user.UUID = uuid.New()
	user.Enable = 1

	// 4. 存库
	err = model.Insert(ctx, repo.Db, user)
	return
}

func DoGetUserList(ctx context.Context, page int, pageSize int, keyword string) (status int, message string, ret any, err error) {
	model := UserModelDB{}

	users, total, err := model.FindUsers(ctx, repo.Db, page, pageSize, keyword)
	if err != nil {
		ret = map[string]any{}
		return
	}

	//封装返回数据
	pageData := PageData{
		List:  users,
		Total: int(total),
	}
	//
	ret = pageData
	return
}

func DoChangePassword(ctx context.Context, id int, password string) (status int, message string, ret any, err error) {
	model := UserModelDB{}

	// 1. 校验用户是否已存在
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("账号不存在")
		return
	}

	// 2. 校验密码 管理员不校验密码
	//if ok := utils.PasswordCheckByBcrypt(oldpassword, u.Password); !ok {
	//	err = errors.New("密码错误")
	//	return
	//}

	// 3. 更新密码
	p, err := utils.PasswordByBcrypt(password)
	if err != nil {
		status = http.StatusInternalServerError
		message = "密码加密失败"
		return
	}
	u.Password = p
	_, err = model.Update(ctx, repo.Db, u)
	if err != nil {
		status = http.StatusInternalServerError
		message = "更新失败"
		return
	}
	message = "success"
	return
}

func DoGetUser(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := UserModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("账号不存在")
		return
	}
	ret = u
	return
}

func DoUpdateUser(ctx context.Context, user UserModel) (status int, message string, ret any, err error) {
	model := UserModelDB{}
	u, err := model.FindById(ctx, repo.Db, user.ID)
	if err != nil || u == nil {
		err = errors.New("账号不存在")
		return
	}

	u.UpdatedAt = time.Time{}
	ignored := []string{"GModel"}                // 忽略 CreatedAt 和 UpdatedAt 字段
	zeroChecks := []string{"Enable", "IsSystem"} // 忽略 CreatedAt 和 UpdatedAt 字段

	updates, err := structutils.CompareStructs(u, &user, ignored, zeroChecks)
	if err != nil {
		err = errors.New("未知错误")
		return
	}
	if len(updates) > 0 {
		// updates 包含了不同的字段和值
		// 可以使用 updates 更新数据库或其他操作
		_, err = model.Update(ctx, repo.Db, u)

	}

	return
}

func DoDeleteUser(ctx context.Context, id int) (status int, message string, ret any, err error) {
	model := UserModelDB{}
	u, err := model.FindById(ctx, repo.Db, id)
	if err != nil || u == nil {
		err = errors.New("账号不存在")
		return
	}
	err = model.Remove(ctx, repo.Db, *u)
	return
}
