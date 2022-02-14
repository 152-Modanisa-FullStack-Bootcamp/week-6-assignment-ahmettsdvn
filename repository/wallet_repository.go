package repository

import (
	"restApp/config"
	db "restApp/config/database"
	"restApp/constants"
	"restApp/model"
)

type IWalletRepo interface {
	GetData() (model.WalletDtoResponse, error)
	GetDataByUsername(username string) (model.WalletDto, error)
	UpdateInsert(username string, values ...int) (bool, error)
	SaveBalance(username string, value int) (bool, error)
}

type WalletRepo struct{}

var orcl = db.GetOracleConfig()
var local, _ = ILocalWalletRepo.GetLocalData(&LocalWalletRepo{})

/*
	Get, Post ve Put işlemleri için gerekli olan tüm metodlar buradadır.
*/

func (wallet *WalletRepo) GetData() (model.WalletDtoResponse, error) {

	var error_ error
	var wallets model.WalletDtoResponse
	conn, err := config.OracleConn(orcl)

	if err != nil {
		error_ = err
	}

	defer conn.Close()
	sql := "select username, balanceAmount from Wallet"
	rows, err := conn.Query(sql)
	defer rows.Close()

	if err != nil {
		error_ = err
	}

	for rows.Next() {
		var wallet model.WalletDto
		rows.Scan(&wallet.Username, &wallet.BalanceAmount)
		wallets = append(wallets, wallet)
	}

	return wallets, error_
}

func (wallet *WalletRepo) GetDataByUsername(username string) (model.WalletDto, error) {
	var error_ error
	var wallets model.WalletDto
	conn, err := config.OracleConn(orcl)

	if err != nil {
		error_ = err
	}

	defer conn.Close()
	sql := "select username, balanceAmount from Wallet where username = :1"
	err = conn.QueryRow(sql, username).Scan(&wallets.Username, &wallets.BalanceAmount)

	if err != nil {
		error_ = err
	}

	return wallets, error_
}

func (wallet *WalletRepo) UpdateInsert(username string, values ...int) (bool, error) {
	var error_ error
	var b bool
	if values[0]&values[1] < -100 {
		return false, nil
	} else {
		conn, err := config.OracleConn(orcl)

		if err != nil {
			error_ = err
			b = false
		}

		defer conn.Close()
		sql := constants.UpdateBalanceQuery(username, values[0], values[1])
		_, err = conn.Query(sql)
		if err != nil {
			error_ = err
			b = false
		}

		b = true
		if err != nil {
			error_ = err
			b = false
		}
		return b, error_
	}
}

func (w *WalletRepo) SaveBalance(username string, value int) (bool, error) {
	var error_ error
	var b bool
	if value < -100 {
		return false, nil
	} else {
		conn, err := config.OracleConn(orcl)

		if err != nil {
			error_ = err
			b = false
		}

		defer conn.Close()
		sql := constants.PostBalanceQuery(username, value)
		_, err = conn.Query(sql)
		if err != nil {
			error_ = err
			b = false
		}

		b = true
		if err != nil {
			error_ = err
			b = false
		}

		return b, error_
	}
}
