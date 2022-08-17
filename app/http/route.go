package http

import (
	"github.com/ascorpio/fly/app/http/module/demo"
	"github.com/ascorpio/fly/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
