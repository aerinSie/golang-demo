package userhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo is ...
func GetUserInfo(c *gin.Context) {
	userID := c.Param("userID")
	c.String(http.StatusOK, fmt.Sprintf("hello, world! user: %s", userID))
}
