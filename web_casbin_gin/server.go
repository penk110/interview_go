package main

import (
	"github.com/penk110/interview_go/web_casbin_gin/router"
)

/*
参考：
https://casbin.org/docs/zh-CN/overview
https://casbin.org/docs/zh-CN/admin-portal


验证：
curl http://127.0.0.1:8000/api/v1/user?username=user_A
curl -X POST http://127.0.0.1:8000/api/v1/user?username=user_A

model.conf访问模型
policy.csv用户权限配置


*/

func main() {

	router.Run(":8000")
}
