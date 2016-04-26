package handlers

import (
//	"fmt"
//	"github.com/spf13/cobra"
//	"github.com/spf13/viper"
//	"os"
//	"github.com/openshift/source-to-image/pkg/api"
)

type JavaHandler struct{}

func (handler JavaHandler) HandleBuild(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Build(ctx *CommandContext) (error, int) {

	//	s2iConfig := api.Config{
	//		DisplayName: "abc",
	//	}

	return nil, 0
}

func (handler JavaHandler) HandleInspect(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Inspect(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler JavaHandler) HandleLog(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Log(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler JavaHandler) HandleRestart(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Restart(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler JavaHandler) HandleRun(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Run(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler JavaHandler) HandleStop(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Stop(ctx *CommandContext) (error, int) {
	return nil, 0
}

func isAJavaService(service string, config map[string]interface{}) bool {
	if serviceMap, rightType := config[service].(map[string]string); rightType {
		if _, found := serviceMap["java"]; found {
			return true
		}
	}
	return false
}
