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
	"bitbucket.org/camilobermudez/huitaca/utils"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

const huitacaFile string = "huitaca"

var verbose bool
var debug bool

var VerboseLogger *log.Logger
var DebugLogger *log.Logger
var StdErrLogger *log.Logger
var StdOutLogger *log.Logger

var Config = map[string]interface{}{}
var wd string

var HandlerChain = []handlers.Handler{
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
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "x", false, "Enable debugging, much more verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {

	// Init loggers
	spew.Config.Indent = "    "
	StdErrLogger = log.New(os.Stderr, "", 0)
	StdOutLogger = log.New(os.Stdout, "", 0)
	if verbose || debug {
		VerboseLogger = log.New(os.Stderr, "* ", 0)
	} else {
		VerboseLogger = log.New(ioutil.Discard, "", 0)
	}
	if debug {
		DebugLogger = log.New(os.Stderr, "*** ", 0)
	} else {
		DebugLogger = log.New(ioutil.Discard, "", 0)
	}

	// Init working directory
	wd, err := os.Getwd()
	if err != nil {
		VerboseLogger.Println("Error resolving current directory: " + err.Error())
		DebugLogger.Println(err)
		os.Exit(1)
	} else {
		VerboseLogger.Println("Working directory: ", wd)
	}

	// Load the default configuration
	deepMerge(Config, utils.ConfigDefaults)

	// Merge in the global configuration
	globalConfig := viper.New()
	globalConfig.SetConfigFile(utils.GlobalHuitacaFile)
	globalConfig.SetConfigType("toml")
	if err := globalConfig.ReadInConfig(); err == nil {
		VerboseLogger.Println("Huitaca global file: ", globalConfig.ConfigFileUsed())
		DebugLogger.Println("global config:\n" + spew.Sdump(globalConfig.AllSettings()))
		deepMerge(Config, (globalConfig.AllSettings()))
	} else {
		VerboseLogger.Println("Error parsing global huitaca file: " + err.Error())
		DebugLogger.Println(err)
		os.Exit(1)
	}

	// Now load and merge in the project configuration
	projectConfig := viper.New()
	projectConfig.SetConfigFile(wd + string(os.PathSeparator) + utils.HuitacaFileName)
	projectConfig.SetConfigType("toml")
	if err := projectConfig.ReadInConfig(); err == nil {
		VerboseLogger.Println("Huitaca project file: ", projectConfig.ConfigFileUsed())
		deepMerge(Config, projectConfig.AllSettings())
		DebugLogger.Println("Effective config:\n" + spew.Sdump(Config))
	} else {
		VerboseLogger.Println("Error parsing project huitaca file: " + err.Error())
		DebugLogger.Println(err)
		os.Exit(1)
	}

}

func handleCommand(
	cmd *cobra.Command,
	services []string,
	determinantFunction func(handlers.Handler, *handlers.CommandContext) bool,
	handlerFunction func(handlers.Handler, *handlers.CommandContext) (error, int)) {

	ctx := handlers.CommandContext{
		Command:       cmd,
		Config:        Config,
		VerboseLogger: VerboseLogger,
		StdErrLogger:  StdErrLogger,
		StdOutLogger:  StdOutLogger,
	}

	for _, service := range services {
		if _, found := Config[service]; !found {
			StdErrLogger.Println("Error: Service '" + service + "' not found")
			os.Exit(1)
		}
		ctx.Service = service
		for _, handler := range HandlerChain {
			if determinantFunction(handler, &ctx) {
				if err, returnCode := handlerFunction(handler, &ctx); err != nil {
					StdErrLogger.Println(err.Error())
					os.Exit(returnCode)
				} else {
					os.Exit(returnCode)
				}
				break
			}
		}
	}
}

func deepMerge(dst map[string]interface{}, src map[string]interface{}) {

	for k, srcValue := range src {
		switch srcValue.(type) {
		case map[string]interface{}:
			if dstValue, exists := dst[k]; !exists {
				dst[k] = srcValue
			} else {
				switch dstValue.(type) {
				case string:
					dst[k] = srcValue
				case map[string]interface{}:
					deepMerge(dstValue.(map[string]interface{}),
						srcValue.(map[string]interface{}))
				}
			}
		case string:
			dst[k] = srcValue
		}
	}
}
