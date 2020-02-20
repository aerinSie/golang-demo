package handler
import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

// UserRegisterRequest is ...
type UserRegisterRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
}

// Data is ...
type Data struct {
	Token string `json:"token" example:"abcd.abcd.abcd"`
}

// UserLoginRequest is ...
type UserLoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// UserInfoRequest is ...
type UserInfoRequest struct {
	Account string `json:"account"`
	Token   string `json:"token"`
}

// UserInfoResponse is ...
type UserInfoResponse struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
}

// UserChangePwdRequest is ...
type UserChangePwdRequest struct {
	Account     string `json:"account"`
	OldPassword string `json:"oldPassword`
	NewPassword string `json:"newPassword"`
}
// GetResponse is ...
type GetResponse struct{
	Code     int `json:"code" example:"1"`
	Message string `json:"message" example:"ok"`
	data UserInfoResponse `json:"data`
}

// BaseResponse is ...
type BaseResponse struct{
	Code     int `json:"code" example:"1"`
	Message string `json:"message" example:"ok"`
	Data Data `json:"data"`
}

// GetUserInfo is ...
func GetUserInfo(c *gin.Context) {
	userID := c.Param("userID")
	c.String(http.StatusOK, fmt.Sprintf("hello, world! user: %s", userID))
}

// PostRegister godoc
// @Summary 使用者註冊
// @Tags User
// @Description add user
// @Accept  json
// @Produce json
// @Param body body UserRegisterRequest true "必填"
// @Success 200 object BaseResponse 
// @Router /user/register [post]
func PostRegister(c *gin.Context) {
	var data Data
	data.Token = "abcd.abcd.abcd"

	var req UserRegisterRequest
	c.BindJSON(&req)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    data,
	})

}

// PostLogin godoc
// @Summary 使用者登錄
// @Tags Common
// @Description user login
// @Accept  json
// @Produce json
// @Param body body UserLoginRequest true "必填"
// @Success 200 object BaseResponse "1-ok 0-fail"
// @Router /login [post]
func PostLogin(c *gin.Context) {

	var data Data
	data.Token = "abcd.abcd.abcd"
	var req UserLoginRequest
	c.BindJSON(&req)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    data,
	})

}

// GetUserInfoByAccount godoc
// @Summary 查詢使用者資訊
// @Tags User
// @Description get user info data
// @Accept  json
// @Produce json
// @Param account query string true "search by account" 
// @Success 200 object GetResponse 
// @Router /user/info [get]
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

// PutUserChangePassword godoc
// @Summary 修改使用者密碼
// @Tags User
// @Description update user login password
// @Accept  json
// @Produce json
// @Param body body UserChangePwdRequest true "必填"
// @Success 200 object BaseResponse 
// @Router /user/change-password [put]
func PutUserChangePassword(c *gin.Context) {

	var req UserChangePwdRequest
	c.BindJSON(&req)

	// response
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
	})

}
