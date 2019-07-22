package main

import (
	"./models"
	//"github.com/gorilla/websocket"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"

)



var templates *template.Template

func main() {
	models.New()
	setupRoutes()
}

func setupRoutes() {
	r := mux.NewRouter()
	bal := models.GetBalance()
	templates = template.Must(template.ParseGlob("pages/*.html"))
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("POST")
	r.HandleFunc("/upgrade-click/", models.BuildCUpgradeHandler(bal)).Methods("POST")
	r.HandleFunc("/add/", models.BalanceHandler).Methods("POST")
	r.HandleFunc("/auto-inc-one/", models.BuildAutoIncOneHandler(bal)).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	go models.AutoIncOne()
	bal := models.GetBalance()
	status := models.GetStatus()
	/*clickStatus := models.GetClickStatus()
	autoIncOneString := models.GetAutoIncOneString()
	CoinPerSecond := models.GetCoinPerSecond()*/
	if status.ClickStatus == 0 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "20 coins to upgrade",
			"ClickPower": "1",
			"IncOneString": status.AutoIncOneString,
			"CoinPerSecond": status.CoinPerSecond,
		})
		if err != nil {
			return
		}
		
	}
	if status.ClickStatus == 1 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "100 coins to upgrade",
			"ClickPower": "2",
			"IncOneString": status.AutoIncOneString,
			"CoinPerSecond": status.CoinPerSecond,

		})
		if err != nil {
			return
		}
		
	}
	if status.ClickStatus == 2 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "275 coins to upgrade",
			"ClickPower": "4",
			"IncOneString": status.AutoIncOneString,
			"CoinPerSecond": status.CoinPerSecond,

		})
		if err != nil {
			return
		}
	}
	if status.ClickStatus == 3 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "No more upgrades",
			"ClickPower": "6",
			"IncOneString": status.AutoIncOneString,
			"CoinPerSecond": status.CoinPerSecond,

		})
		if err != nil {
			return
		}
	}
}



