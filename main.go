package main

import (
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	"html/template"
	"./models"
)

var templates *template.Template
func main() {
	go models.AutoIncOne()
	r := mux.NewRouter()
	templates = template.Must(template.ParseGlob("pages/*.html"))
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("POST")
	r.HandleFunc("/upgrade-click/", models.CUpgradeHandler).Methods("POST")
	r.HandleFunc("/add/", models.BalanceHandler).Methods("POST")
	r.HandleFunc("/auto-inc-one/", models.AutoIncOneHandler).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bal := models.GetBalance()
	clickStatus := models.GetClickStatus()
	autoIncOneString := models.GetAutoIncOneString()
	if clickStatus == 0 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "20 coins to upgrade",
			"ClickPower": "1",
			"IncOneString": autoIncOneString,
		})
		if err != nil {
			return
		}
	}
	if clickStatus == 1 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "100 coins to upgrade",
			"ClickPower": "2",
			"IncOneString": autoIncOneString,

		})
		if err != nil {
			return
		}
	}
	if clickStatus == 2 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "275 coins to upgrade",
			"ClickPower": "4",
			"IncOneString": autoIncOneString,

		})
		if err != nil {
			return
		}
	}
	if clickStatus == 3 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "No more upgrades",
			"ClickPower": "6",
			"IncOneString": autoIncOneString,

		})
		if err != nil {
			return
		}
	}
}

