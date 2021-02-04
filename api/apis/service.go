package controllers

import (
	"net/http"
	"strconv"

	"github.com/zknow/redisRankDemo/api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

//Index 顯示所有Rank資訊
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//Search 新增rk的頁面
func Search(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", gin.H{"title": "Main website"})
}

//Input 新增rk的頁面
func Input(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", gin.H{"title": "Main website"})
}

//ShowAll 顯示所有Rank資訊
func ShowAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.RanksForTop100())
}

//ShowUser 顯示使用者排名跟分數
func ShowUser(c *gin.Context) {
	name := c.Params.ByName("name")

	if models.IsUserExist(name) == redis.Nil {
		c.JSON(http.StatusOK, map[string]interface{}{"Error": "找不到User"})
	} else {
		rank, score := models.GetUser(name)
		c.JSON(http.StatusOK, map[string]interface{}{"Rank": rank + 1, "Score": score})
	}
}

//DeleteUser 刪除使用者
func DeleteUser(c *gin.Context) {
	name := c.Params.ByName("name")
	if models.RemoveUser(name) {
		c.JSON(http.StatusOK, map[string]interface{}{"Success": "Remove" + name})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"Error": "找不到User"})
	}
}

//DeleteAll 刪除使用者
func DeleteAll(c *gin.Context) {
	if models.RemoveAll() {
		c.JSON(http.StatusOK, map[string]interface{}{"Success": "Remove all success"})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"Error": "找不到User"})
	}
}

//UpdateUser update使用者
func UpdateUser(c *gin.Context) {
	name := c.Request.FormValue("name")
	score, err := strconv.ParseFloat(c.Request.FormValue("score"), 64)
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		if models.AddMember(name, score) != nil {
			c.JSON(http.StatusOK, "更新排名失敗,請重新嘗試")
		}
		c.JSON(http.StatusOK, "更新成功")
	}
}
