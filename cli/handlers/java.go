package handlers

import (
	//	"fmt"
	//	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
	//	"strings"
	"errors"
	"github.com/openshift/source-to-image/pkg/api"
)

type JavaHandler struct{}

var builderImages = map[string]string{
	"7": "huitaca/java:open-jdk-7",
	"8": "huitaca/java:open-jdk-8",
}

func (handler JavaHandler) HandleBuild(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Build(ctx *CommandContext) (error, int) {

	serviceConfig := ctx.Config[service].(map[string]interface{})
	s2iConfig := api.Config{
		DisplayName: getString(serviceConfig, []string{"displayName"}),
		Description: getString(serviceConfig, []string{"description"}),
	}

	javaVersion, _ := getString(serviceConfig, []string{"java"})
	if builderImage, ok := builderImages[javaVersion]; ok {
		s2iConfig.BuilderImage = builderImage
		ctx.VerboseLogger.Println("Using builder image '" + builderImage + "'")
	} else {
		return errors.New("Error: Java version '" + javaVersion + "' not found"), 1
	}

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
	if _, found := getString(config, []string{service, "java"}); found {
		return true
	}
	return false
}
