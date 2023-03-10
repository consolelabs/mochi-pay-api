package model

type Wallet struct {
	Id              string
	ProfileGlobalId string
	Platform        string
	MochiWallet     *MochiWallet
	EvmWallet       *EvmWallet
}

type MochiWallet struct {
	Id string
}

type EvmWallet struct {
	Id      string
	ChainId string
	Address string
}
