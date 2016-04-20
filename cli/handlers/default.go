package handlers

import (
	//	"fmt"
	//	"github.com/spf13/cobra"
	//	"github.com/spf13/viper"
	//	"os"
	//	"bitbucket.org/camilobermudez/huitaca/cmd"
	"errors"
)

type DefaultHandler struct{}

func (handler DefaultHandler) HandleBuild(ctx *CommandContext) bool {
	return true
}

func (handler DefaultHandler) Build(ctx *CommandContext) (error, int) {
	ctx.VerboseLogger.Println("No suitable platform could be found to handle this command")
	return errors.New("Ooops! The command could not be processed, there's probably something missing in your huitaca file, please check."), 1
}

func (handler DefaultHandler) HandleInspect(ctx *CommandContext) bool {
	return true
}

func (handler DefaultHandler) Inspect(ctx *CommandContext) (error, int) {
	ctx.VerboseLogger.Println("No suitable platform could be found to handle this command")
	return errors.New("Ooops! The command could not be processed, there's probably something missing in your huitaca file, please check."), 1
}

func (handler DefaultHandler) HandleLog(ctx *CommandContext) bool {
	return true
}

func (handler DefaultHandler) Log(ctx *CommandContext) (error, int) {
	ctx.VerboseLogger.Println("No suitable platform could be found to handle this command")
	return errors.New("Ooops! The command could not be processed, there's probably something missing in your huitaca file, please check."), 1
}

func (handler DefaultHandler) HandleRestart(ctx *CommandContext) bool {
	return true
}

func (handler DefaultHandler) Restart(ctx *CommandContext) (error, int) {
	ctx.VerboseLogger.Println("No suitable platform could be found to handle this command")
	return errors.New("Ooops! The command could not be processed, there's probably something missing in your huitaca file, please check."), 1
}

func (handler DefaultHandler) HandleRun(ctx *CommandContext) bool {
	return true
}

func (handler DefaultHandler) Run(ctx *CommandContext) (error, int) {
	ctx.VerboseLogger.Println("No suitable platform could be found to handle this command")
	return errors.New("Ooops! The command could not be processed, there's probably something missing in your huitaca file, please check."), 1
}

func (handler DefaultHandler) HandleStop(ctx *CommandContext) bool {
	return true
}

func (handler DefaultHandler) Stop(ctx *CommandContext) (error, int) {
	ctx.VerboseLogger.Println("No suitable platform could be found to handle this command")
	return errors.New("Ooops! The command could not be processed, there's probably something missing in your huitaca file, please check."), 1
}
