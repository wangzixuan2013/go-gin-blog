package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/upload"
	v1 "gin-blog/routers/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"gin-blog/routers/api"
)

func InitRouter() *gin.Engine {

	gin.DisableConsoleColor()//保存到文件不需要颜色
	//handleFile,_ := os.OpenFile(logging.AccessLog(),os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	file, _ := os.Create(logging.AccessLog())
	gin.DefaultWriter = file

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	//token 验证中间件

	//file访问服务

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))


	r.GET("/auth", api.GetAuth)

	r.POST("/upload", api.UploadImage)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
