package constants

const (
	CONFIG_PATH = "/root/config"
	COSMOS_URL = "http://localhost:1317"
	SERVER_CONF = "/root/sentinel/shell_scripts/shadowsocks.json"
	FoundConfig = "found previously generated config file. Starting node..."
	ConfigNotFound = "Could not read config file. Starting the wizard now..."
	AskForWallet = "Do you already have a Tendermint Wallet Address? [y/n]"
	AskForTMName = "Please Enter Your Tendermint Username: "
	AskForTMPass = "Please Enter Password for your Tendermint Wallet: "
	AskForPrice = "How many SENTs do you want to charge for Per GB of Bandwidth? "
	AskForSOCKS5Pass = "Please enter a password for your SOCKS5 node: "
	AskForEncryptionMethod = `Please select an encryption method number for your SOCKS5 node:
	1) aes-256-cfb
	2) rc4-md5
	3) aes-128-cfb
	4) salsa20
	5) chacha20
	`
	AskForConfirmation = "Is Everything okay? [y/n]"
)