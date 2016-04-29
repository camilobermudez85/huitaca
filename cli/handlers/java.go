package handlers

import (
	//	"fmt"
	//	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
	//	"strings"
	"errors"
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/build/strategies/sti"
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

	serviceConfig := ctx.Config[ctx.Service].(map[string]interface{})
	s2iConfig := buildS2iConfig(ctx.Service, ctx.Config)

	javaVersion := getString(serviceConfig, []string{"java"})
	if builderImage, ok := builderImages[javaVersion]; ok {
		s2iConfig.BuilderImage = builderImage
		ctx.VerboseLogger.Println("Using builder image '" + builderImage + "'")
	} else {
		return errors.New("Error: Java version '" + javaVersion + "' not found"), 1
	}

	if stiStruct, err := sti.New(s2iConfig, nil); err == null {
		if result, err := stiStruct.Build(s2iConfig); err == null {
		} else {
		}
	} else {
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
	return getString(config, []string{service, "java"}) != ""
}
