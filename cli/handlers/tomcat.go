package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type TomcatHandler struct{}

func (handler TomcatHandler) HandleBuild(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler TomcatHandler) Build(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler TomcatHandler) HandleInspect(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler TomcatHandler) Inspect(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler TomcatHandler) HandleLog(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler TomcatHandler) Log(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler TomcatHandler) HandleRestart(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler TomcatHandler) Restart(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler TomcatHandler) HandleRun(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler TomcatHandler) Run(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler TomcatHandler) HandleStop(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler TomcatHandler) Stop(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}
