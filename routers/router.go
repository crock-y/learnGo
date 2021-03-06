package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章详情
		apiv1.GET("/article/:id", v1.GetArticle)
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//新建文章
		apiv1.POST("/articles", v1.Create)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.Edit)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.Delete)
	}
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
