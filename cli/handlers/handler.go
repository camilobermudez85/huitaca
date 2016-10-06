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
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//"os"
	//	"bitbucket.org/camilobermudez/huitaca/cmd"
	"bitbucket.org/camilobermudez/huitaca/utils"
	"github.com/openshift/source-to-image/pkg/api"
	"log"
)

const defaultScriptsURL = "image:///usr/libexec/s2i"
const defaultAssembleUser = "1001"

type CommandContext struct {
	Service       string
	Command       *cobra.Command
	Config        map[string]interface{}
	VerboseLogger *log.Logger
	StdErrLogger  *log.Logger
	StdOutLogger  *log.Logger
}

type Handler interface {
	HandleBuild(ctx *CommandContext) bool
	Build(ctx *CommandContext) (error, int)

	HandleInspect(ctx *CommandContext) bool
	Inspect(ctx *CommandContext) (error, int)

	HandleLog(ctx *CommandContext) bool
	Log(ctx *CommandContext) (error, int)

	HandleRestart(ctx *CommandContext) bool
	Restart(ctx *CommandContext) (error, int)

	HandleRun(ctx *CommandContext) bool
	Run(ctx *CommandContext) (error, int)

	HandleStop(ctx *CommandContext) bool
	Stop(ctx *CommandContext) (error, int)
}

func buildS2iConfig(ctx *CommandContext) *api.Config {

	serviceConfig := ctx.Config[ctx.Service].(map[string]interface{})
	s2iConfig := api.Config{
		DisplayName:   getString(serviceConfig, []string{"displayName"}),
		Description:   getString(serviceConfig, []string{"description"}),
		DockerConfig:  buildDockerConfig(ctx.Config),
		DockerCfgPath: getString(ctx.Config, []string{"huitaca", "docker", "dockerCfgPath"}),
		// PullAuthentication: ...
		// IncrementalAuthentication: ...
		// DockerNetworkMode: ...
		PreserveWorkingDir: true,
		DisableRecursive:   false,
		Source:             utils.Getwd(),
		//Tag: ... Defined at the platform
		BuilderPullPolicy:       api.PullIfNotPresent,
		PreviousImagePullPolicy: api.PullIfNotPresent,
		Incremental:             true,
		RemovePreviousImage:     false,
		ScriptsURL:              defaultScriptsURL,
		AssembleUser:            defaultAssembleUser,
	}

	if gitRef, err := ctx.Command.Flags().GetString("git-ref"); err == nil {
		s2iConfig.Ref = gitRef
	}

	s2iConfig.Tag = ctx.Service + ":" + utils.GetwdHash()
	if imageTag, err := ctx.Command.Flags().GetString("image-tag"); err == nil {
		s2iConfig.Tag = imageTag
	}

	return &s2iConfig
}

func buildDockerConfig(config map[string]interface{}) *api.DockerConfig {

	return &api.DockerConfig{
		Endpoint: getString(config, []string{"huitaca", "docker", "endpoint"}),
		CertFile: getString(config, []string{"huitaca", "docker", "certFile"}),
		KeyFile:  getString(config, []string{"huitaca", "docker", "keyFile"}),
		CAFile:   getString(config, []string{"huitaca", "docker", "CAFile"}),
	}
}

func getString(m map[string]interface{}, path []string) string {
	var v interface{} = m
	var found bool
	for i, k := range path {
		v, found = v.(map[string]interface{})[k]
		if !found {
			break
		}
		switch v.(type) {
		case string:
			if i == len(path)-1 {
				return v.(string)
			}
			break
		}
	}
	return ""
}
