package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type TomcatHandler struct{}

func (handler *TomcatHandler) handleBuild(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *TomcatHandler) build(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *TomcatHandler) handleInspect(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *TomcatHandler) inspect(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *TomcatHandler) handleLog(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *TomcatHandler) log(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *TomcatHandler) handleRestart(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *TomcatHandler) restart(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *TomcatHandler) handleRun(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *TomcatHandler) run(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *TomcatHandler) handleStop(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *TomcatHandler) stop(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}
