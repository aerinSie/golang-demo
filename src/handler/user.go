package userhandler
import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)


// UserRegisterRequest is ...
type UserRegisterRequest struct {
	Account  string `json:"account" gorm:"column:account;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
	Name     string `json:"name" gorm:"column:name;not null" binding:"required" validate:"min=5,max=128"`
	Birthday string `json:"birthday" gorm:"column:birthday;not null" binding:"required" validate:"min=10,max=10"`
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

// UserInfoResponse is ...
type UserInfoResponse struct {
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
	var token Token
	token.Token = "abcd.abcd.abcd"

	var req UserRegisterRequest
	c.BindJSON(&req)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    token,
	})

}

// PostLogin is ...
func PostLogin(c *gin.Context) {

	var token Token
	token.Token = "abcd.abcd.abcd"
	var req UserLoginRequest
	c.BindJSON(&req)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    token,
	})

}

// GetUserInfoByAccount is ...
func GetUserInfoByAccount(c *gin.Context) {
	// account, _ := c.GetQuery("account")
	// token, _ := c.GetQuery("token")

	var userInfoRes UserInfoResponse
	userInfoRes.Name = "Allen"
	userInfoRes.Birthday = "1990/01/01"
	userInfoRes.Gender = 1
	// response
	c.JSON(http.StatusOK, userInfoRes)
}

// PutUserChangePassword is ...
func PutUserChangePassword(c *gin.Context) {

	var req UserLoginRequest
	c.BindJSON(&req)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
	})

}
