package service

import (
	"restApp/model"
	"restApp/repository"
)

type WalletService struct {
	repository.WalletRepo
}

var ws = NewWalletService(WalletService{})

func NewWalletService(w WalletService) repository.IWalletRepo {
	return &WalletService{WalletRepo: w.WalletRepo}
}

func GetData() (model.WalletDtoResponse, error) {
	return ws.GetData()
}

func GetDataByUsername(username string) (model.WalletDto, error) {
	return ws.GetDataByUsername(username)
}

func Update(username string, values ...int) (bool, error) {
	return ws.UpdateInsert(username, values[0], values[1])
}

func Save(username string, value int) (bool, error) {
	return ws.SaveBalance(username, value)
}
