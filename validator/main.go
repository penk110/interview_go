package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
)

const (
	exitCode = 1
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

type UserInfo struct {
	UserName   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

func SignInH(ctx *gin.Context) {
	userInfo := &UserInfo{}

	if err := ctx.ShouldBind(userInfo); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			fmt.Printf("%v\n", errs)
			ctx.JSON(200, gin.H{
				"code": -10,
				"msg":  errs.Error(),
				"data": "",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"code": -10,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code": 10,
		"msg":  "",
		"data": "success",
	})
	return
}

func main() {

	// SignIn
	router.POST("/signIn", SignInH)

	if err := router.Run(":8001"); err != nil {
		log.Fatalf("start server failed, err: %v\n", err)
		os.Exit(exitCode)
	}

}
