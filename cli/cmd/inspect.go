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
	"bitbucket.org/camilobermudez/huitaca/handlers"
	//	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect service",
	Short: "Inspect properties for a specific service",
	Long: `
on.`,
	Run: func(cmd *cobra.Command, args []string) {
		handleCommand(
			cmd,
			args,
			handlers.Handler.HandleInspect,
			handlers.Handler.Inspect)
	},
}
var tags []string

func init() {

	RootCmd.AddCommand(inspectCmd)

	inspectCmd.Flags().StringSliceVarP(
		&tags,
		"tag",
		"t",
		[]string{"id", "ip", "port"},
		"The property tag(s) or key(s) to be inspected")

}
