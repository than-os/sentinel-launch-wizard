package main

import (
	"github.com/fatih/color"
	"github.com/than-os/launch-wizard/constants"
	"github.com/than-os/launch-wizard/services"
	"github.com/than-os/launch-wizard/utils"
	"os"
	"time"
)

func init() {
	ok, pass, name := utils.CheckConfig(constants.CONFIG_PATH)
	if ok {
		color.Green("%s", constants.FoundConfig)
		services.StartNode(pass, name)
		os.Exit(1)
	}
	color.Red("%s", constants.ConfigNotFound)
}

func main() {
	services.StartWizard()
	time.Sleep(time.Second * 2)
	wallet := services.GenerateTmWallet("qwdscfmddsdsdsfvvf", "qwerty")
	color.Cyan("%s", wallet)

}