package services

import (
	"os"

	"github.com/ascorpio/fly/framework"
	"github.com/ascorpio/fly/framework/contract"
)

// FlyConsoleLog 代表控制台输出
type FlyConsoleLog struct {
	FlyLog
}

// NewFlyConsoleLog 实例化FlyConsoleLog
func NewFlyConsoleLog(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	log := &FlyConsoleLog{}

	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)
	log.c = c
	return log, nil
}
