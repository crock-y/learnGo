package v1

import (
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
		code = e.SUCCESS
		//
		data = models.GetArticle(id)
		//data["total"] = models.GetArticleTotal()
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

//获取多个文章
func GetArticles(c *gin.Context) {
}

//新增文章
func Create(c *gin.Context) {
}

//修改文章
func Edit(c *gin.Context) {
}

//删除文章
func Delete(c *gin.Context) {
}
