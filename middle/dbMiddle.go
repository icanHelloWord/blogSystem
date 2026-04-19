package middle

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DbMiddle(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func GetDBFromContext(c *gin.Context) *gorm.DB {
	return c.Value("db").(*gorm.DB)
}
