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
	//	"bitbucket.org/camilobermudez/huitaca/cmd"
	"errors"
	//	"github.com/davecgh/go-spew/spew"
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
