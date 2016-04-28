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

package cmd

import (
	"bitbucket.org/camilobermudez/huitaca/handlers"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [service...]",
	Short: "Builds the indicated service(s) image(s)",
	Long: `
Builds executable container image(s) for one or multiple services
on the current project. If no service is provided all of the existing 
services on the current project will be built.`,
	Run: func(cmd *cobra.Command, args []string) {
		handleCommand(
			cmd,
			args,
			handlers.Handler.HandleBuild,
			handlers.Handler.Build)
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)
	logCmd.Flags().StringVar(new(string), "git-ref", false,
		"Optionally provide a git ref to use to build the image")
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
