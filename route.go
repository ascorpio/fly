package main

import "go_frame/framework"

// 注册路由
func registerRouter(core *framework.Core) {
	// 设置控制器
	core.Get("foo", FooControllerHandler)

	//// 需求1+2:HTTP方法+静态路由匹配
	//core.Post("/user/login", UserLoginController)
	//
	//// 需求3:批量通用前缀
	//subjectApi := core.Group("/subject")
	//{
	//	subjectApi.Post("/add", SubjectAddController)
	//	// 需求4:动态路由
	//	subjectApi.Delete("/:id", SubjectDelController)
	//	subjectApi.Put("/:id", SubjectUpdateController)
	//	subjectApi.Get("/:id", SubjectGetController)
	//	subjectApi.Get("/list/all", SubjectListController)
	//}
}