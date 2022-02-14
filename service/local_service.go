package service

import (
	"restApp/model"
	"restApp/repository"
)

type LocalWalletService struct {
	repository.LocalWalletRepo
}

// type ILocalWalletRepoService struct {
// 	lw repository.ILocalWalletRepo
// }

var ls = NewLocalWalletService(repository.LocalWalletRepo{})

func NewLocalWalletService(w repository.LocalWalletRepo) repository.ILocalWalletRepo {
	return &LocalWalletService{LocalWalletRepo: w}
}

func GetLocalData() (model.LocalWalletResponse, error) {
	return ls.GetLocalData()
}
