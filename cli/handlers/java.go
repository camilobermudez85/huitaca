package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type JavaHandler struct{}

func (handler JavaHandler) HandleBuild(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler JavaHandler) Build(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler JavaHandler) HandleInspect(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler JavaHandler) Inspect(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler JavaHandler) HandleLog(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler JavaHandler) Log(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler JavaHandler) HandleRestart(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler JavaHandler) Restart(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler JavaHandler) HandleRun(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler JavaHandler) Run(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}

func (handler JavaHandler) HandleStop(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler JavaHandler) Stop(config *viper.Viper, cmd *cobra.Command) int {
	return 1
}
