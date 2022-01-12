package main

import (
	"github.com/go-playground/validator"
	"net/http"

	"github.com/penk110/interview_go/validator_gin/trans"

	"github.com/gin-gonic/gin"
)

// UserInfo is user info struct
type UserInfo struct {
	Name       string `json:"name" binding:"required"`
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func main() {
	var err error

	// init transletor
	err = trans.Init("zh")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.POST("/signup", func(c *gin.Context) {
		var u UserInfo
		if err := c.ShouldBind(&u); err != nil {
			// 获取validator.ValidationErrors类型的errors
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// 非validator.ValidationErrors类型错误直接返回
				c.JSON(http.StatusOK, gin.H{
					"msg": err.Error(),
				})
				return
			}
			// validator.ValidationErrors类型错误则进行翻译
			c.JSON(http.StatusOK, gin.H{
				"msg": trans.RemoveTopStruct(errs.Translate(trans.Trans)),
			})
			return
		}
		// 业务逻辑代码...

		c.JSON(http.StatusOK, "success")
	})

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 40,
			"data": "",
			"msg":  "page not found",
		})
		return
	})
	err = engine.Run(":8000")
	if err != nil {
		panic("start engine failed, err: %v" + err.Error())
	}
}
