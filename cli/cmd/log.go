// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	//	"fmt"
	"bitbucket.org/camilobermudez/huitaca/handlers"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log [service...]",
	Short: "Manages the output of the specified service(s)",
	Long: `
.`,
	Run: func(cmd *cobra.Command, args []string) {

		handleCommand(
			cmd,
			args,
			handlers.Handler.HandleLog,
			handlers.Handler.Log)
	},
}

func init() {
	RootCmd.AddCommand(logCmd)

	logCmd.Flags().BoolVarP(new(bool), "follow", "f", false,
		"Attach and follow the output of the provided service(s)")

}
