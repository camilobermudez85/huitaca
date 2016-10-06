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
	//	"github.com/davecgh/go-spew/spew"
	//	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"os"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect [service]",
	Short: "Inspects properties for a specific service",
	Long: `
Inspects a service properties, only one service at a time.
If no service is specified inspects the properties of the project.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 1 {
			StdErrLogger.Println("Error: No more than one service can be inspected at once.")
			os.Exit(1)
		}

		handleCommand(
			cmd,
			args,
			handlers.Handler.HandleInspect,
			handlers.Handler.Inspect)
	},
}

var tags []string
var follow bool
var list bool

func init() {

	RootCmd.AddCommand(inspectCmd)

	inspectCmd.Flags().StringSliceVarP(&tags, "tag", "t", []string{"id", "ip", "port"},
		"The property key(s) to be inspected")

	inspectCmd.Flags().BoolVarP(&follow, "follow", "f", false,
		"Output inspected value(s) at regular intervals")

	inspectCmd.Flags().BoolVarP(&list, "list", "l", false,
		"List the inspectable keys for the provided service")
}
