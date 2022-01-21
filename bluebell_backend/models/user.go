package models

// ParamSignUp 注册参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required" `
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required" `
	Password string `json:"password" binding:"required"`
}

// User 用户信息结构体
type User struct {
	UserID   int64  `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}
