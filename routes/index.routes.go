package routes

import (
	"quality-air-golang/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/quality-air", controllers.FindQualityAir)

	return r
}
