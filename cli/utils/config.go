package utils

import (
	"fmt"
	"github.com/kardianos/osext"
	//	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//	"os"
)

var projectConfig = viper.New()
var effectiveConfig = viper.New()
var wd string

func InitConfig() {
	wd, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Println("Error resolving current directory.", err)
		panic(err)
	}

	fmt.Println("WD: " + wd)
	projectConfig.SetConfigFile(wd + "/huitaca")
	projectConfig.SetConfigType("toml")

	if err := projectConfig.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Error parsing huitaca file: ", err)
	}

}

func GetConfig() *viper.Viper {
	return projectConfig
}
