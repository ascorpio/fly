package middleware

import "go_frame/framework"

// Recovery 机制，将协程中的函数异常进行捕获
func Recovery() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				c.Json(500, err)
			}
		}()

		c.Next()

		return nil
	}
}
