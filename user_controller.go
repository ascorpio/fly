package main

import (
	"fmt"
	"go_frame/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	go func() {
		// 每个独立开启的协程需要单独进行 recover 的捕捉
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("1234")
	}()
	time.Sleep(2 * time.Second)
	c.Json(200, "ok, UserLoginController")
	return nil
}
