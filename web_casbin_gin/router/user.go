package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _userR = &userR{}

type userR struct {
}

func userRegister(iRouter gin.IRoutes) {
	iRouter.GET("/user", _userR.Get)
	iRouter.POST("/user", _userR.Post)
	iRouter.DELETE("/user", _userR.Delete)
	iRouter.PUT("/user/:user_id", _userR.Put)
}

func (ur *userR) Get(ctx *gin.Context) {
	var (
		sub string
		obj string
		act string
	)
	sub = ctx.Query("username")
	obj = ctx.Request.URL.Path
	act = ctx.Request.Method
	if !CheckPerm(ctx, sub, obj, act) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "1000010",
		"data":    "get user page",
		"message": "ok",
	})
	return
}

func (ur *userR) Post(ctx *gin.Context) {
	var (
		sub string
		obj string
		act string
	)
	sub = ctx.Query("username")
	obj = ctx.Request.URL.Path
	act = ctx.Request.Method
	if !CheckPerm(ctx, sub, obj, act) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "1000010",
		"data":    "post user page",
		"message": "ok",
	})
}

func (ur *userR) Delete(ctx *gin.Context) {
	var (
		sub string
		obj string
		act string
	)
	sub = ctx.Query("username")
	obj = ctx.Request.URL.Path
	act = ctx.Request.Method
	if !CheckPerm(ctx, sub, obj, act) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "1000010",
		"data":    "delete user page",
		"message": "ok",
	})
}

func (ur *userR) Put(ctx *gin.Context) {
	var (
		sub string
		obj string
		act string
	)
	sub = ctx.Query("username")
	obj = ctx.Request.URL.Path
	act = ctx.Request.Method
	log.Println(ctx.Request.URL.Path, ctx.Request.URL.RawPath, ctx.Request.URL.RequestURI())
	if !CheckPerm(ctx, sub, obj, act) {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "1000010",
		"data":    "put user page",
		"message": "ok",
	})
}
