package main

import (
	"fmt"
	"net/http"
	"restApp/controller"
	"restApp/service"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Go Rest Example")
	fmt.Println("Go - Oracle 21c - Godror - database/sql - net/http")

	s := service.NewWalletService(service.WalletService{})
	m := controller.NewWalletController(&s)
	r := httprouter.New()
	r.GET("/", m.Get)
	r.GET("/:username", m.GetByUsername)
	r.PUT("/:username", m.Put)
	r.POST("/:username", m.Post)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
