package router

import (
	. "github.com/zknow/redisRankDemo/api/apis"

	"github.com/gin-gonic/gin"
)

//GetRouter get instance
func GetRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", Index)
	router.GET("/input", Input)
	router.GET("/user", Search)
	router.GET("/users", ShowAll)
	router.GET("/user/:name", ShowUser)
	router.POST("/user", UpdateUser)
	router.DELETE("/delete/:name", DeleteUser)
	router.DELETE("/delete", DeleteAll)
	return router
}
