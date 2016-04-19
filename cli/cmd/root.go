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
	"bitbucket.org/camilobermudez/huitaca/types"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

//var cfgFile string
var verbose bool

var VerboseLogger *log.Logger

var projectConfig = viper.New()
var effectiveConfig = viper.New()
var wd string

var HandlerChain = []types.Handler{
	handlers.TomcatHandler{},
	handlers.JavaHandler{},
	handlers.DefaultHandler{}}

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "huitaca",
	Short: "PaaS all the way through from the development environment to the cloud.",
	Long: `
Huitaca is a simple, easy to configure, 12-factor-app oriented PaaS that
seamlessly streamlines the delivery pipeline starting at the development
environment.`,
	// Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hello world") },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Make the operation more talkative")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	// Init loggers
	if verbose {
		VerboseLogger = log.New(os.Stderr, "* ", 0)
	} else {
		VerboseLogger = log.New(ioutil.Discard, "", 0)
	}

	// Init working dir and project config
	wd, err := os.Getwd()
	if err != nil {
		VerboseLogger.Panicln("Error resolving current directory.", err)
	} else {
		VerboseLogger.Println("Working directory: ", wd)
	}

	projectConfig.SetConfigFile(wd + "/huitaca")
	projectConfig.SetConfigType("toml")

	if err := projectConfig.ReadInConfig(); err == nil {
		VerboseLogger.Println("Huitaca file: ", projectConfig.ConfigFileUsed())
	} else {
		VerboseLogger.Panicln("Error parsing huitaca file: ", err)
	}

	fmt.Println(GetEffectiveConfig().AllKeys())

}

func GetEffectiveConfig() *viper.Viper {
	return projectConfig
}
