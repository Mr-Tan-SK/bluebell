package mysql

import (
	"bluebell_backend/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "WorryFree"

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user(user_id, username, password) values(?,?,?)"
	user.Password = encrypt(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// Login 用户登录核心操作
func Login(u *models.User) (err error) {
	opasswd := u.Password
	sqlStr := "select user_id,username,password from user where username=?"
	err = db.Get(u, sqlStr, u.Username)
	if err != nil {
		return errors.New("数据库获取数据错误")
	}
	if encrypt(opasswd) != u.Password {
		return errors.New("用户名或密码错误")
	}
	return nil
}

// CheckUserExit 查询用户是否存在
func CheckUserExit(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int8
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// 密码加密
func encrypt(opasswd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(opasswd)))
}
