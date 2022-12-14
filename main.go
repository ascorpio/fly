// Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/ascorpio/fly/app/console"
	"github.com/ascorpio/fly/app/http"
	"github.com/ascorpio/fly/framework"
	"github.com/ascorpio/fly/framework/provider/app"
	"github.com/ascorpio/fly/framework/provider/config"
	"github.com/ascorpio/fly/framework/provider/distributed"
	"github.com/ascorpio/fly/framework/provider/env"
	"github.com/ascorpio/fly/framework/provider/id"
	"github.com/ascorpio/fly/framework/provider/kernel"
	"github.com/ascorpio/fly/framework/provider/log"
	"github.com/ascorpio/fly/framework/provider/orm"
	"github.com/ascorpio/fly/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewFlyContainer()
	// 绑定App服务提供者
	container.Bind(&app.FlyAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.FlyEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.FlyConfigProvider{})
	container.Bind(&id.FlyIDProvider{})
	container.Bind(&trace.FlyTraceProvider{})
	container.Bind(&log.FlyLogServiceProvider{})
	container.Bind(&orm.GormProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.FlyKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
