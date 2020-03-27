// usermgmt is 處理 user路徑的業務邏輯

package usermgmt

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
}

var users = []User{
	{
		Account:  "Jerry",
		Password: "",
		Name:     "Jerry",
		Birthday: "1990/1/1",
		Gender:   1,
	},
}

var userinfos = []User{
	{
		Account:  "Jerry",
		Name:     "Jerry",
		Birthday: "1990/1/1",
		Gender:   1,
	},
}

//userType is可以被呼叫回傳的物件
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"account": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"birthday": &graphql.Field{
				Type: graphql.String,
			},
			"gender": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//"info": &graphql.Field{},
			/**
			  * get userlist is
				http://localhost:9205/user/v2/userlist?query={list{account,name,birthday,gender}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "Get User list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return userinfos, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			/**
			  * create user is
				http://localhost:9205/user/v2/userlist?mutation+_{create(account:"jerry",password:"password",name:"jerry",birthday:"1990/1/1",gender:1)}
			*/
			"create": &graphql.Field{
				Type:        userType,
				Description: "Create new User",
				Args: graphql.FieldConfigArgument{
					"account": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"birthday": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"gender": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					user := User{
						Account:  params.Args["account"].(string),
						Password: params.Args["password"].(string),
						Name:     params.Args["name"].(string),
						Birthday: params.Args["birthday"].(string),
						Gender:   params.Args["gender"].(int),
					}
					users = append(users, user)
					return user, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
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
func GerUserList(c *gin.Context) {

	result := executeQuery(c.Param("query"), schema)

	// response
	c.JSON(http.StatusOK, result)
}

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
type GetResponse struct {
	Data UserInfoResponse `json:"data`

	Code    int    `json:"code" example:"1"`
	Message string `json:"message" example:"ok"`
}

// BaseResponse is ...
type BaseResponse struct {
	Data Data `json:"data"`

	Code    int    `json:"code" example:"1"`
	Message string `json:"message" example:"ok"`
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
