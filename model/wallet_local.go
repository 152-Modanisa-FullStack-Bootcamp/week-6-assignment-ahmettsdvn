package model

/*
	WalletLocal değerlerini local.json dosyasından alır
*/

type LocalWallet struct {
	InitialBalanceAmount int `json:"initialBalanceAmount"`
	MinumumBalanceAmount int `json:"minumumBalanceAmount"`
}

type LocalWalletResponse LocalWallet
