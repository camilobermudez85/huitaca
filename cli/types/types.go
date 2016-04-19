package types

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type Handler interface {
	HandleBuild(config *viper.Viper, cmd *cobra.Command) bool
	Build(config *viper.Viper, cmd *cobra.Command) int

	HandleInspect(config *viper.Viper, cmd *cobra.Command) bool
	Inspect(config *viper.Viper, cmd *cobra.Command) int

	HandleLog(config *viper.Viper, cmd *cobra.Command) bool
	Log(config *viper.Viper, cmd *cobra.Command) int

	HandleRestart(config *viper.Viper, cmd *cobra.Command) bool
	Restart(config *viper.Viper, cmd *cobra.Command) int

	HandleRun(config *viper.Viper, cmd *cobra.Command) bool
	Run(config *viper.Viper, cmd *cobra.Command) int

	HandleStop(config *viper.Viper, cmd *cobra.Command) bool
	Stop(config *viper.Viper, cmd *cobra.Command) int
}
