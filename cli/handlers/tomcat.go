package handlers

import (
	//	"fmt"
	"errors"
	//	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
)

type TomcatHandler struct{}

func (handler TomcatHandler) HandleBuild(ctx *CommandContext) bool {
	return false
}

func (handler TomcatHandler) Build(ctx *CommandContext) (error, int) {
	return errors.New("Method not yet implemented"), 3
}

func (handler TomcatHandler) HandleInspect(ctx *CommandContext) bool {
	return false
}

func (handler TomcatHandler) Inspect(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler TomcatHandler) HandleLog(ctx *CommandContext) bool {
	return false
}

func (handler TomcatHandler) Log(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler TomcatHandler) HandleRestart(ctx *CommandContext) bool {
	return false
}

func (handler TomcatHandler) Restart(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler TomcatHandler) HandleRun(ctx *CommandContext) bool {
	return false
}

func (handler TomcatHandler) Run(ctx *CommandContext) (error, int) {
	return nil, 0
}

func (handler TomcatHandler) HandleStop(ctx *CommandContext) bool {
	return false
}

func (handler TomcatHandler) Stop(ctx *CommandContext) (error, int) {
	return nil, 0
}
