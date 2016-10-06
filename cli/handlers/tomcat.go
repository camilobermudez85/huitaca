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
