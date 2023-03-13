package model

type Wallet struct {
	Id              string       `json:"id"`
	ProfileGlobalId string       `json:"profile_global_id"`
	Platform        string       `json:"platform"`
	MochiWallet     *MochiWallet `json:"mochi_wallet"`
	EvmWallet       *EvmWallet   `json:"evm_wallet"`
}

type MochiWallet struct {
	Id string `json:"id"`
}

type EvmWallet struct {
	Id      string `json:"id"`
	ChainId string `json:"chain_id"`
	Address string `json:"address"`
}
