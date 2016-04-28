package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
	//	"bitbucket.org/camilobermudez/huitaca/cmd"
	"github.com/openshift/source-to-image/pkg/api"
	"log"
)

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

func buildS2iConfig(service string, config map[string]interface{}) *api.Config {

	serviceConfig := config[service].(map[string]interface{})
	s2iConfig := api.Config{
		DisplayName:   getString(serviceConfig, []string{"displayName"}),
		Description:   getString(serviceConfig, []string{"description"}),
		DockerConfig:  buildDockerConfig(config),
		DockerCfgPath: getString(config, []string{"huitaca", "docker", "dockerCfgPath"}),
		// PullAuthentication: ...
		// IncrementalAuthentication: ...
		// DockerNetworkMode: ...
		PreserveWorkingDir: true,
		DisableRecursive:   false,
		Source:             os.Getwd(),
		//Ref: ... Defined from a flag at the platform
		//Tag: ... Defined at the platform
		BuilderPullPolicy: api.PullIfNotPresent,
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
