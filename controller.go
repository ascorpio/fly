package main

import (
	"context"
	"fly/framework"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Second)
	defer cancel()

	finish := make(chan struct{}, 1)
	panicChan := make(chan any, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		c.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "panic")
	case <-finish:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(200, "ok")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "timeout")
		c.SetHasTimeout()
	}

	return nil
}
