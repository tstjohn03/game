package models

import (
	"html/template"
	"net/http"
	"log"
)

var templates *template.Template

type Status struct {
	ClickStatus int
	CoinPerSecond int
	AutoIncOneString string
	AutoOneLvl int
}
type BalanceType struct {
	Coins int
}
var new bool = true
var balance *BalanceType
var status *Status

func New() {
	if new == true {
		new = false
		balance = &BalanceType{0}
		status = &Status{0, 0, "", 0}
	}
}

func BalanceHandler(w http.ResponseWriter, r *http.Request){
	SetBalance(status.ClickStatus, getCoins(), balance)
	http.Redirect(w, r, "/", 302)
}

func SetBalance(cs int, coins int, bal *BalanceType) {
	if cs == 0 {
		coins++
		log.Println("Coins: ", coins)
	}
	if cs == 1 {
		coins += 2
		log.Println("Coins: ", coins)
	}
	if cs == 2 {
		coins += 4
		log.Println("Coins: ", coins)
	}
	if cs == 3 {
		coins += 6
		log.Println("Coins: ", coins)
	}
	bal.Coins = coins
}

func GetBalance() *BalanceType{
	log.Println("Balance: ", balance)
	return balance
}

func getCoins() int {
	log.Println("Coins: ", balance.Coins)
	return balance.Coins
}

func GetClickStatus() int{
	log.Println("Click Status: ", status.ClickStatus)
	return status.ClickStatus
}

func GetStatus() *Status{
	log.Print("Status:", status)
	return status
}

func GetNew() bool {
	log.Print("New? ", new)
	return new
}