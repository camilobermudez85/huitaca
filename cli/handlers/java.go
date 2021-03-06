// Copyright © 2016 Camilo Bermúdez <camilobermudez85@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

import (
	//	"fmt"
	//	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
	//	"strings"
	"errors"
	//"github.com/openshift/source-to-image/pkg/api"
	"github.com/davecgh/go-spew/spew"
	"github.com/openshift/source-to-image/pkg/build"
	"github.com/openshift/source-to-image/pkg/build/strategies/sti"
)

type JavaHandler struct{}

var builderImages = map[string]string{
	"7": "huitaca/java:openjdk-7",
	"8": "huitaca/java:openjdk-8",
}

func (handler JavaHandler) HandleBuild(ctx *CommandContext) bool {
	return isAJavaService(ctx.Service, ctx.Config)
}

func (handler JavaHandler) Build(ctx *CommandContext) (error, int) {

	serviceConfig := ctx.Config[ctx.Service].(map[string]interface{})
	s2iConfig := buildS2iConfig(ctx)

	javaVersion := getString(serviceConfig, []string{"java"})
	if builderImage, ok := builderImages[javaVersion]; ok {
		s2iConfig.BuilderImage = builderImage
		ctx.VerboseLogger.Println("Using builder image '" + builderImage + "'")
	} else {
		return errors.New("Error: Java version '" + javaVersion + "' not found"), 1
	}

	ctx.VerboseLogger.Println(spew.Sdump(s2iConfig))
	if stiStruct, err := sti.New(s2iConfig, build.Overrides{}); err == nil {
		if result, err := stiStruct.Build(s2iConfig); err != nil {
			panic(err)
		} else {
			ctx.VerboseLogger.Println(spew.Sdump(result))
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
