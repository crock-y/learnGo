package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"net/http"

	"github.com/astaxie/beego/validation"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id 不能小于零")

	var data interface{}
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.ERROR_NOT_EXIST_ARTICLE
		if models.ExistArticleById(id) {
			data = models.GetArticle(id)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//获取多个文章
func GetArticles(c *gin.Context) {

	maps := make(map[string]interface{})
	if title := c.Query("title"); title != "" {
		maps["title"] = title
	}

	code := e.SUCCESS
	data := models.GetArticles(util.GetPage(c), setting.PageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章
func Create(c *gin.Context) {

	//models.AddArticle("")
}

//修改文章
func Edit(c *gin.Context) {
}

//删除文章
func Delete(c *gin.Context) {
}
