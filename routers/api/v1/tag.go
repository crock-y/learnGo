package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/unknwon/com"
    "net/http"

    "github.com/EDDYCJY/go-gin-example/pkg/e"
    "github.com/EDDYCJY/go-gin-example/pkg/util"
    "github.com/EDDYCJY/go-gin-example/pkg/setting"
    "github.com/EDDYCJY/go-gin-example/models"

)

//获取多个文章标签
func GetTags(c *gin.Context) {
    name := c.Query("name")

    maps := make(map[string]interface{})
    data := make(map[string]interface{})

    if name != "" {
        maps["name"] = name
    }

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        maps["state"] = state
    }

    code := e.SUCCESS

    data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
    data["total"] = models.GetTagTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

//新增文章标签
func AddTag(c *gin.Context) {
    name := c.Query("name")
    createdBy := c.Query("createdBy")
    var state int = -1


    state = com.StrTo(c.Query("state")).MustInt()


    if models.ExistTagByName(name) == true {
        return
    }
    if models.AddTag(name, state, createdBy) {
      code:=e.SUCCESS
      c.JSON(http.StatusOK, gin.H{
          "code" : code,
          "msg" : e.GetMsg(code),
      })
    }



}

//修改文章标签
func EditTag(c *gin.Context) {
}

func DeleteTag(c *gin.Context) {
}
