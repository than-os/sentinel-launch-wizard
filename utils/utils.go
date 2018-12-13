package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/than-os/launch-wizard/models"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var (
	Password string
	Name string
)

func ReadInput(question string) (string, error) {

	var input string
	color.Green("%s", question)
	_, err := fmt.Scanln(&input)
	return input, err
}

func ReadNumber(question string) float64 {
	price, err := ReadInput(question)
	if err != nil {
		color.Red("%s", "error while reading the number")
		return 0.0
	}
	p, e := strconv.ParseFloat(price, 64)
	if e != nil {
		color.Red("%s", "please enter a number!")
		return ReadNumber(question)
	}

	return p
}

func CheckConfig(path string) (bool, string, string) {
	var c models.Config
	fi, err := ioutil.ReadFile(path)
	if err != nil {
		return false, "", ""
	}

	err = json.Unmarshal(fi, &c)
	if err != nil {
		return false, "", ""
	}
	if c.Account.Password == "" || c.Account.Name == "" {
		return false, "", ""
	}

	return true, c.Account.Password, c.Account.Name
}



func ExecCommand(cmdName string, cmdArgs []string) {
	var cmd *exec.Cmd
	if cmdName == "mongod" {
		cmd = exec.Command(cmdName)
	} else {
		cmd = exec.Command(cmdName, cmdArgs...)
	}
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		color.Red("%s", os.Stderr, "error creating output pipeline: ", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("output | %s\n", scanner.Text())
		}
	}()

	if cmdName == "mongod" {
		err = cmd.Run()
		//err = cmd.Wait()
		color.Red("%s error occurred: ", err)
	} else {
		err = cmd.Start()
		color.Red("%s error occurred: ", err)
	}

	err = cmd.Wait()

}

func WriteJsonFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 644)
}