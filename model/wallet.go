package model

/*
	Buradaki yapılar veritabanından gelen değerler ile doldurulur.
	DTO(Data Transfer Object) kullanıcıya göstermek istediğimiz değerler
*/

type Wallet struct {
	Id            int    `db:"ID"`
	Username      string `db:"USERNAME"`
	BalanceAmount int    `db:"BALANCEAMOUNT"`
}

type WalletDto struct {
	Username      string `db:"USERNAME"`
	BalanceAmount int    `db:"BALANCEAMOUNT"`
}

type WalletResponse []Wallet
type WalletDtoResponse []WalletDto
