package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
	//	"bitbucket.org/camilobermudez/huitaca/cmd"
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
