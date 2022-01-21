package logic

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/jwt"
	"bluebell_backend/pkg/snowflake"
	"errors"
	"fmt"
)

func Signup(u *models.ParamSignUp) error {
	// 判断用户存不存在
	if err := mysql.CheckUserExit(u.Username); err != nil {
		return err
	}
	// 生成 ID
	userID, err := snowflake.GenID()
	if err != nil {
		fmt.Printf("snowflake算法分配ID失败，err: %v\n", err.Error())
		return err
	}
	// 构造一个 user实例
	user := models.User{
		UserID:   userID,
		Username: u.Username,
		Password: u.Password,
	}
	// 保存进数据库
	err = mysql.InsertUser(&user)
	if err != nil {
		return errors.New("注册用户失败," + err.Error())
	}
	return nil
}

func Login(u *models.ParamLogin) (aToken string, err error) {
	err = mysql.CheckUserExit(u.Username)
	if err == nil {
		return "", errors.New("用户不存在")
	}
	user := &models.User{
		Username: u.Username,
		Password: u.Password,
	}
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	// 数据库判断无误后,开始生成JWT的token
	aToken, _, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return
}
