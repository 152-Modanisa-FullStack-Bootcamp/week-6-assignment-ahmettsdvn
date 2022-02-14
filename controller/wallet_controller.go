package controller

import (
	"encoding/json"
	"net/http"
	"restApp/model"
	"restApp/repository"
	"restApp/service"

	"github.com/julienschmidt/httprouter"
)

/*
	Burası controller sayfası. Burada servisten gelen değerler alınır
	http istekleri için metodlar oluşturulur.
*/

type IWalletController interface {
	Get(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetByUsername(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Put(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Post(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type WalletController struct {
	ws service.WalletService
	ls service.LocalWalletService
}

func (c *WalletController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	cc, err := c.ws.GetData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(cc)
}

func (c *WalletController) GetByUsername(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	username := p.ByName("username")
	w.Header().Add("Content-Type", "application/json")
	cc, err := c.ws.GetDataByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(cc)
}

func (c *WalletController) Put(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	localJson, _ := c.ls.LocalWalletRepo.GetLocalData()

	var e model.WalletDto
	_ = json.NewDecoder(r.Body).Decode(&e)
	res, _ := json.Marshal(e)

	l := &model.WalletDto{}
	_ = json.Unmarshal(res, l)

	c.ws.UpdateInsert(username, l.BalanceAmount, localJson.InitialBalanceAmount)
}

func (c *WalletController) Post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	var e model.WalletDto
	_ = json.NewDecoder(r.Body).Decode(&e)
	res, _ := json.Marshal(e)

	l := &model.WalletDto{}
	_ = json.Unmarshal(res, l)
	c.ws.SaveBalance(username, l.BalanceAmount)
}

func NewWalletController(c *repository.IWalletRepo) IWalletController {
	return &WalletController{ws: service.WalletService{}}
}
