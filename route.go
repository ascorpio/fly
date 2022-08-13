package main

import (
	"go_frame/framework"
	"time"
)

// 注册路由
func registerRouter(core *framework.Core) {
	// 需求1+2:HTTP方法+静态路由匹配
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Post("/add", SubjectAddController)
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}

	// 需求4：可实现多层嵌套数据
	gp1 := core.Group("/g1")
	gp1.Get("/login1", UserLoginController)
	gp2 := gp1.Group("/g2")
	gp2.Get("/login2", UserLoginController)
}
