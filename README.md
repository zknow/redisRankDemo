# 使用redis + gin實作CRUD功能

### 使用套件
```
github.com/gin-gonic/gin
github.com/go-redis/redis
```

### Routes
```
router.GET("/", Index)
router.GET("/input", Input)
router.GET("/user", Serch)
router.GET("/users", ShowAll)
router.GET("/user/:name", ShowUser)
router.POST("/user", UpdateUser)
router.DELETE("/delete/:name", DeleteUser)
router.DELETE("/delete", DeleteAll)
```