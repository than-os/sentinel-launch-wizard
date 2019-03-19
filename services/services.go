package services

import (
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"github.com/than-os/launch-wizard/constants"
	"github.com/than-os/launch-wizard/models"
	"github.com/than-os/launch-wizard/utils"
	"log"
	"net/http"
	"os"
	"time"
)

func StartWizard() {
	var config models.Config
	ok, err := utils.ReadInput(constants.AskForWallet)
	if ok == "n" || ok == "N" {
		config.Account.Name, _ = utils.ReadInput(constants.AskForTMName)
	}
	config.Account.Password, err = utils.ReadInput(constants.AskForTMPass)
	// add it here
	config.Price = utils.ReadNumber(constants.AskForPrice)
	config.SOCKS5MAC, err = utils.ReadInput(constants.AskForSOCKS5Pass)
	encMethod, err := utils.ReadInput(constants.AskForEncryptionMethod)
	ok, err = utils.ReadInput(constants.AskForConfirmation)

	if ok == "y" || ok == "Y" {
		color.Green("%s", "Starting the node...")
		StartNode(config.Account.Password, config.Account.Name)
	}

	switch encMethod {
	case "1":
		config.EncMethod = "aes-256-cfb"
	case "2":
		config.EncMethod = "salsa20"
	case "3":
		config.EncMethod = "rc4-md5"
	case "4":
		config.EncMethod = "chacha20"
	case "5":
		config.EncMethod = "aes-128-cfb"

	}
	wallet := GenerateTmWallet(config.Account.Name, config.Account.Password)
	config.Account.Address = wallet.Address
	config.Account.PubKey = wallet.PubKey
	config.Account.Seed = wallet.Seed
	if err != nil {
		defer StartWizard()
		color.Red("Error while reading user input. ", err)
		os.Exit(1)
	}

	data, err := json.Marshal(config)
	if err != nil {
		color.Red("%s", "error while converting user config: ", err)
		return
	}
	err = utils.WriteJsonFile(constants.CONFIG_PATH, data)
	if err != nil {
		color.Red("%s", "error while converting user config: ", err)
		return
	}
}

func StartNode(tmPass, tmName string) {
	go utils.ExecCommand("mongod", []string{""})
	go utils.ExecCommand("nohup", []string{"gaiacli", "advanced", "rest-server", "--node", "tcp://tm-lcd.sentinelgroup.io:26657", "--chain-id=Sentinel-testnet-1.1", ">>", "tendermint.log"})
	go utils.ExecCommand("gunicorn", []string{"--reload", "-b", "0.0.0.0:3000", "--log-level", "DEBUG", "--chdir", "root", "server:server"})
	time.Sleep(time.Second * 1)
	log.Println("here it goes: ", tmPass, tmName)
	utils.ExecCommand("python", []string{"/root/app.py", tmPass, tmName})
}

func GenerateTmWallet(name, password string) models.TmAccount {
	var wallet models.TmAccount
	 body := map[string]string{"name": name, "password": password}
	 bytesData, err := json.Marshal(body)
	 if err != nil {
	 	color.Red("%s", "error while user data marshal: \n", err)
	 	return wallet
	 }
	resp, err := http.Post(constants.COSMOS_URL + "/keys", "application/json" , bytes.NewBuffer(bytesData))
	if err != nil {
		color.Red("%s", "error while generating new Tendermint Wallet for the user: \n", err)
		return wallet
	}
	if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
		color.Red("%s", "error while reading user wallet info: \n", err)
		return wallet
	}
	defer resp.Body.Close()
	return wallet

}