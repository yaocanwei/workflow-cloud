package header

import (
	"github.com/gin-gonic/gin"
)

// Validator 指定用户访问
func Validator(c *gin.Context) {
	userID := c.GetHeader("wf-usereid")
	if userID == "" {
		c.AbortWithStatusJSON(401, gin.H{"data": nil, "msg": "当前用户不允许访问"})
		return
	}
	c.Next()
	return
}
