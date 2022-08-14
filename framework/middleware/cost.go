package middleware

import (
	"fly/framework"
	"fmt"
	"time"
)

// Cost 记录请求耗时时间
func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {

		// 获取当前时间
		t0 := time.Now()
		c.Next()
		// 获取经过了多少时间
		t1 := time.Now()
		fmt.Println("run time：", t1.Sub(t0))

		return nil
	}
}
