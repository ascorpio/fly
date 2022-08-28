package http

import (
	"github.com/ascorpio/fly/app/http/module/demo"
	"github.com/ascorpio/fly/framework/gin"
	"github.com/ascorpio/fly/framework/middleware/static"
)

func Routes(r *gin.Engine) {

	r.Use(static.Serve("/", static.LocalFile("./vue_web/dist", false)))

	demo.Register(r)
}
