package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type HandlerError struct {
	msg  string
	code int
}

func (e *HandlerError) Error() string {
	return e.msg
}

func (e *HandlerError) Code() int {
	return e.code
}

type Handler interface {
	HandleBuild(config *viper.Viper, cmd *cobra.Command) bool
	Build(config *viper.Viper, cmd *cobra.Command) error

	HandleInspect(config *viper.Viper, cmd *cobra.Command) bool
	Inspect(config *viper.Viper, cmd *cobra.Command) error

	HandleLog(config *viper.Viper, cmd *cobra.Command) bool
	Log(config *viper.Viper, cmd *cobra.Command) error

	HandleRestart(config *viper.Viper, cmd *cobra.Command) bool
	Restart(config *viper.Viper, cmd *cobra.Command) error

	HandleRun(config *viper.Viper, cmd *cobra.Command) bool
	Run(config *viper.Viper, cmd *cobra.Command) error

	HandleStop(config *viper.Viper, cmd *cobra.Command) bool
	Stop(config *viper.Viper, cmd *cobra.Command) error
}
