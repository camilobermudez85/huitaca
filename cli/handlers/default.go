package handlers

import (
	//	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

type DefaultHandler struct{}

func (handler DefaultHandler) HandleBuild(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler DefaultHandler) Build(config *viper.Viper, cmd *cobra.Command) error {
	return nil
}

func (handler DefaultHandler) HandleInspect(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler DefaultHandler) Inspect(config *viper.Viper, cmd *cobra.Command) error {
	return nil
}

func (handler DefaultHandler) HandleLog(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler DefaultHandler) Log(config *viper.Viper, cmd *cobra.Command) error {
	return nil
}

func (handler DefaultHandler) HandleRestart(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler DefaultHandler) Restart(config *viper.Viper, cmd *cobra.Command) error {
	return nil
}

func (handler DefaultHandler) HandleRun(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler DefaultHandler) Run(config *viper.Viper, cmd *cobra.Command) error {
	return nil
}

func (handler DefaultHandler) HandleStop(config *viper.Viper, cmd *cobra.Command) bool {
	return false
}

func (handler DefaultHandler) Stop(config *viper.Viper, cmd *cobra.Command) error {
	return nil
}
