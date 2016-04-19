package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type JavaHandler struct{}

func (handler *JavaHandler) handleBuild(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *JavaHandler) build(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *JavaHandler) handleInspect(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *JavaHandler) inspect(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *JavaHandler) handleLog(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *JavaHandler) log(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *JavaHandler) handleRestart(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *JavaHandler) restart(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *JavaHandler) handleRun(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *JavaHandler) run(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler *JavaHandler) handleStop(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler *JavaHandler) stop(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}
