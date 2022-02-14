package repository

import (
	"encoding/json"
	"io"
	"os"
	"restApp/model"
)

type ILocalWalletRepo interface {
	GetLocalData() (model.LocalWalletResponse, error)
}

type LocalWalletRepo struct{}

/*
	Bu metod local.json dosyasındaki verileri alıp döndürür
*/

func (lw *LocalWalletRepo) GetLocalData() (model.LocalWalletResponse, error) {
	l := &model.LocalWalletResponse{}
	var _error error
	file, err := os.Open("./.config/local.json")

	if err != nil {
		_error = err
	}
	defer file.Close()

	read, err := io.ReadAll(file)
	if err != nil {
		_error = err
	}

	err = json.Unmarshal(read, l)
	if err != nil {
		_error = err
	}

	return *l, _error
}
