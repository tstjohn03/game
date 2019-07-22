package models

import (
	"net/http"
	"time"
	//"fmt"
	"log"
)


func GetAutoIncOneString() string {
	log.Println("Auto Inc One String: ", status.AutoIncOneString)
	return status.AutoIncOneString
}

func GetCoinPerSecond() int {
	log.Println("Coin Per Second: ", status.CoinPerSecond)
	return status.CoinPerSecond
}

func AutoIncOne() {
	for status.AutoOneLvl >= 0{
		if status.AutoOneLvl == 0 {
			status.AutoIncOneString = "200 Coins"
			status.CoinPerSecond = 0
		}
		if status.AutoOneLvl == 1 {
			status.AutoIncOneString = "350 Coins to upgrade"
			for status.AutoOneLvl == 1 {
				status.CoinPerSecond = 1
				balance.Coins++
				log.Println("Balance After Increment:", balance.Coins)
				time.Sleep(1000 * time.Millisecond)
			}
		}
		if status.AutoOneLvl == 2 {
			status.AutoIncOneString = "No More Upgrades"
			for status.AutoOneLvl == 2 {
				status.CoinPerSecond = 2
				balance.Coins+= 2
				log.Println("Balance After Increment:", balance.Coins)

				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
	
}
func BuildAutoIncOneHandler(bal *BalanceType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if status.AutoOneLvl == 0 && bal.Coins >= 10 {
			bal.Coins = bal.Coins - 10
			status.AutoOneLvl = 1
			log.Println("Balance:", bal.Coins)
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.AutoOneLvl == 0 && bal.Coins < 200 {
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.AutoOneLvl == 1 && bal.Coins >= 350 {
			bal.Coins = bal.Coins - 350
			status.AutoOneLvl = 2
			log.Println("Balance:", bal.Coins)
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.AutoOneLvl == 1 && bal.Coins < 350 {
			http.Redirect(w, r, "/", 302)
			return
		}
	}
}

func BuildCUpgradeHandler(bal *BalanceType) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request, ) {
		if status.ClickStatus == 0 && bal.Coins >= 20{
			bal.Coins = bal.Coins - 20
			status.ClickStatus = 1
			log.Println("Balance:", bal.Coins)
			log.Println("Click Status:", status.ClickStatus)
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.ClickStatus == 0 && bal.Coins < 20 {
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.ClickStatus == 1 && bal.Coins >= 100{
			bal.Coins = bal.Coins - 100
			status.ClickStatus = 2
			log.Println("Balance:", bal.Coins)
			log.Println("Click Status:", status.ClickStatus)
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.ClickStatus == 1 && bal.Coins < 100 {
			http.Redirect(w, r, "/", 302)
			return
		}
	
		if status.ClickStatus == 2 && bal.Coins >= 275{
			bal.Coins = bal.Coins - 275
			status.ClickStatus = 3
			log.Println("Balance:", bal.Coins)
			log.Println("Click Status:", status.ClickStatus)
			http.Redirect(w, r, "/", 302)
			return
		}
		if status.ClickStatus == 2 && bal.Coins < 275 {
			http.Redirect(w, r, "/", 302)
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}

