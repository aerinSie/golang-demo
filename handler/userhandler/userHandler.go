package userhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegisterRequest is ...
type UserRegisterRequest struct {
	Account  string `json:"account" gorm:"column:account;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Name     string `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=5,max=128"`
	Birthday string `json:"birthday" gorm:"column:birthday;not null" binding:"required" validate:"min=8,max=8"`
	Gender   int    `json:"gender" gorm:"column:gender;not null" binding:"required" `
}

// Token is ...
type Token struct {
	Token string `json:"token"`
}

// UserLoginRequest is ...
type UserLoginRequest struct {
	Account  string `json:"account" gorm:"column:account;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// UserInfoRequest is ...
type UserInfoRequest struct {
	Account string `json:"account" gorm:"column:account;not null" binding:"required" validate:"min=1,max=32"`
	Token   string `json:"token" gorm:"column:token;not null" binding:"required" validate:"min=5,max=128"`
}

// UserInfo is ...
type UserInfo struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
}

// UserChangePwdRequest is ...
type UserChangePwdRequest struct {
	Account     string `json:"account" gorm:"column:account;not null" binding:"required" validate:"min=1,max=32"`
	OldPassword string `json:"oldPassword" gorm:"column:oldPassword;not null" binding:"required" validate:"min=5,max=128"`
	NewPassword string `json:"newPassword" gorm:"column:newPassword;not null" binding:"required" validate:"min=5,max=128"`
}

// GetUserInfo is ...
func GetUserInfo(c *gin.Context) {
	userID := c.Param("userID")
	c.String(http.StatusOK, fmt.Sprintf("hello, world! user: %s", userID))
}

// PostRegister is ...
func PostRegister(c *gin.Context) {

	var req UserLoginRequest
	c.BindJSON(&req)

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    req.Account,
	})

}
