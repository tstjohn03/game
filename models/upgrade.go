package models

import (
	"net/http"
	"time"
)

var autoOneLvl = 0
var autoIncOneString = ""
func GetAutoIncOneString() string {
	return autoIncOneString
}

func AutoIncOne() {
	for autoOneLvl >= 0{
		if autoOneLvl == 0 {
			autoIncOneString = "200 coins"
		}
		if autoOneLvl == 1 {
			autoIncOneString = "350 coins to upgrade"
			for autoOneLvl == 1 {
				bal++
				time.Sleep(1 * time.Millisecond)
			}
		}
		if autoOneLvl == 2 {
			autoIncOneString = "No More Upgrades"
			for autoOneLvl == 2 {
				bal+= 2
				time.Sleep(1 * time.Millisecond)
			}
		}
	}
	
}

func AutoIncOneHandler(w http.ResponseWriter, r *http.Request) {
	if autoOneLvl == 0 && bal >= 200 {
		bal = bal - 200
		autoOneLvl = 1
		http.Redirect(w, r, "/", 302)
		return
	}
	if autoOneLvl == 0 && bal < 200 {
		http.Redirect(w, r, "/", 302)
		return
	}
	if autoOneLvl == 1 && bal >= 350 {
		bal = bal - 350
		autoOneLvl = 2
		http.Redirect(w, r, "/", 302)
		return
	}
	if autoOneLvl == 1 && bal < 350 {
		http.Redirect(w, r, "/", 302)
		return
	}
}

func CUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	if clickStatus == 0 && bal >= 20{
		bal = bal - 10
		clickStatus = 1
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 0 && bal < 20 {
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 1 && bal >= 100{
		bal = bal - 50
		clickStatus = 2
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 1 && bal < 100 {
		http.Redirect(w, r, "/", 302)
		return
	}

	if clickStatus == 2 && bal >= 275{
		bal = bal - 100
		clickStatus = 3
		http.Redirect(w, r, "/", 302)
		return
	}
	if clickStatus == 2 && bal < 275 {
		http.Redirect(w, r, "/", 302)
		return
	}
	http.Redirect(w, r, "/", 302)
}