package types

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type Handler interface {
	handleBuild(config *viper.Viper, cmd *cobra.Command) bool
	build(config *viper.Viper, cmd *cobra.Command) int

	handleInspect(config *viper.Viper, cmd *cobra.Command) bool
	inspect(config *viper.Viper, cmd *cobra.Command) int

	handleLog(config *viper.Viper, cmd *cobra.Command) bool
	log(config *viper.Viper, cmd *cobra.Command) int

	handleRestart(config *viper.Viper, cmd *cobra.Command) bool
	restart(config *viper.Viper, cmd *cobra.Command) int

	handleRun(config *viper.Viper, cmd *cobra.Command) bool
	run(config *viper.Viper, cmd *cobra.Command) int

	handleStop(config *viper.Viper, cmd *cobra.Command) bool
	stop(config *viper.Viper, cmd *cobra.Command) int
}
