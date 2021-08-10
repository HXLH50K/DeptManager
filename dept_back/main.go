package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	cors "github.com/rs/cors/wrapper/gin"
)

//UserInfo --> data table
type Dept struct {
	gorm.Model
	Deptno int
	Dname  string
	Loc    string
}

var (
	DB *gorm.DB
)

func initMySQL() (err error) {
	dsn := "root:123456@(127.0.0.1:3306)/dept?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		method := c.Request.Method
// 		origin := c.Request.Header.Get("Origin") //请求头部
// 		if origin != "" {
// 			//接收客户端发送的origin （重要！）
// 			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
// 			//服务器支持的所有跨域请求的方法
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
// 			//允许跨域设置可以返回其他子段，可以自定义字段
// 			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
// 			// 允许浏览器（客户端）可以解析的头部 （重要）
// 			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
// 			//设置缓存时间
// 			c.Header("Access-Control-Max-Age", "172800")
// 			//允许客户端传递校验信息比如 cookie (重要)
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 		}

// 		//允许类型校验
// 		if method == "OPTIONS" {
// 			c.JSON(http.StatusOK, "ok!")
// 		}

// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Printf("Panic info is: %v", err)
// 			}
// 		}()

// 		c.Next()
// 	}
// }
func main() {
	//connect mysql
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	DB.AutoMigrate(&Dept{})

	r := gin.Default()
	// r.Use(Cors())
	r.Use(cors.AllowAll())
	v1Group := r.Group("v1")
	{
		//Create
		v1Group.POST("/dept", func(c *gin.Context) {
			//前端页面点击待办事项，点击提交
			//从请求中拿数据
			var t Dept
			c.BindJSON(&t)
			//存入数据库
			//返回响应
			if err = DB.Create(&t).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "ok",
				})
			}
		})
		//Read all
		v1Group.GET("/dept/:page/:size", func(c *gin.Context) {
			// 查询todos表所有内容
			var dList []Dept
			pageindex, _ := strconv.Atoi(c.Param("page"))
			pagesize, _ := strconv.Atoi(c.Param("size"))

			var count int
			// if err = DB.Find(&dList).Error; err != nil {
			if err = DB.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&dList).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"content":       dList,
					"totalElements": count,
					"size":          pagesize,
				})
				// fmt.Printf("%#v", dList)
			}
		})
		// Update
		v1Group.PUT("/dept/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在",
				})
				return
			}
			var t Dept
			if err = DB.Debug().Where("id=?", id).First(&t).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.ShouldBind(&t)
			if err = DB.Debug().Save(&t).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"content": t,
					"message": "ok",
				})
			}
		})
		//Delete
		v1Group.DELETE("/dept/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在",
				})
				return
			}
			if err = DB.Debug().Delete(Dept{}, id).Error; err != nil {
				c.JSON((http.StatusOK), gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"content": id,
					"message": "ok",
				})
			}
		})

	}

	r.Run(":8081")
	// var depts []Dept
	// // create
	// dept := Dept{Deptno: 1, Dname: "yanfa", Loc: "beijing"}
	// db.NewRecord(&dept)
	// db.Debug().Create(&dept)

	// db.Find(&depts)
}
