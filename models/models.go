package models

type Config struct {
	Account account `json:"account"`
	EncMethod string `json:"enc_method"`
	Price float64 `json:"price_per_gb"`
	Register register `json:"register"`
	SOCKS5MAC string `json:"socks5Mac"`
}

type account struct {
	Address string `json:"address"`
	Name string `json:"name"`
	Password string `json:"password"`
	PubKey string `json:"pubkey"`
	Seed string `json:"seed"`
}

type register struct {
	Hash string `json:"hash"`
	Token string `json:"token"`
}

type TmAccount struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Address string `json:"address"`
	PubKey string `json:"pub_key"`
	Seed string `json:"seed"`
}