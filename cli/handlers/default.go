package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type DefaultHandler struct{}

func (handler *DefaultHandler) handleBuild(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *DefaultHandler) build(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *DefaultHandler) handleInspect(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *DefaultHandler) inspect(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *DefaultHandler) handleLog(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *DefaultHandler) log(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *DefaultHandler) handleRestart(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *DefaultHandler) restart(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *DefaultHandler) handleRun(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *DefaultHandler) run(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *DefaultHandler) handleStop(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *DefaultHandler) stop(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}
