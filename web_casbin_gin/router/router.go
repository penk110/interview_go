package router

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_go/web_casbin_gin/config"
)

var (
	engine   *gin.Engine
	enforcer *casbin.Enforcer
	err      error
)

func init() {
	engine = gin.Default()
	register()

	if enforcer, err = casbin.NewEnforcer(config.ModelPath, config.PolicyPath); err != nil {
		log.Fatalf("load config file failed, %v\n", err.Error())
	}
}

func register() {
	var (
		v1 gin.IRoutes
	)
	engine.GET("", indexH)
	engine.GET("/favicon.ico", faviconH)

	v1 = engine.Group("/api/v1")
	userRegister(v1)

}

func indexH(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"code":    "1000010",
		"data":    "index page",
		"message": "ok",
	})
	return
}

func faviconH(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "1000010",
		"data":    "favicon.ico",
		"message": "ok",
	})
	return
}

func Run(ls string) {
	var err error
	if err = engine.Run(ls); err != nil {
		log.Fatalf("start engine ")
		return
	}
	return
}

func CheckPerm(ctx *gin.Context, sub, obj, act string) bool {
	var (
		ok  bool
		err error
	)
	log.Printf("sub = %s obj = %s act = %s", sub, obj, act)

	if ok, err = enforcer.Enforce(sub, obj, act); err != nil {
		log.Printf("enforce failed %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 400050,
			"msg":  "内部服务器错误",
		})
		return false
	}
	if !ok {
		log.Println("权限验证不通过: ", sub, obj, act)
		ctx.JSON(http.StatusForbidden, gin.H{
			"code": 400040,
			"msg":  "权限验证不通过",
		})
		return false
	}
	log.Println("权限验证通过")
	//ctx.String(http.StatusOK, "权限验证通过")
	return true
}
