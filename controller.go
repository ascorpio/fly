package main

import "go_frame/framework"

func FooControllerHandler(ctx *framework.Context) error {
	return ctx.Json(200, map[string]any{
		"code": 0,
	})
}
