package main

import (
	"./models"
	//"log"
	"github.com/gorilla/websocket"
	//"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"time"
)



var templates *template.Template

func main() {
	setupRoutes()
}

func setupRoutes() {
	r := mux.NewRouter()
	templates = template.Must(template.ParseGlob("pages/*.html"))
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("POST")
	r.HandleFunc("/upgrade-click/", models.CUpgradeHandler).Methods("POST")
	r.HandleFunc("/add/", models.BalanceHandler).Methods("POST")
	r.HandleFunc("/auto-inc-one/", models.AutoIncOneHandler).Methods("POST")
	//r.HandleFunc("/ws", wsEndpoint)
	r.HandleFunc("/ws", wsHandler)
	//http.HandleFunc("/ws", wsEndpoint)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}



/*func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}*/

/*func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
	}

	log.Println("Client Successfully Connected...")

	reader(ws)
}*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	go models.AutoIncOne()
	bal := models.GetBalance()
	clickStatus := models.GetClickStatus()
	autoIncOneString := models.GetAutoIncOneString()
	coinPerSecond := models.GetCoinPerSecond()
	if clickStatus == 0 {
		err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
			"Balance": bal,
			"ClickUpString": "20 coins to upgrade",
			"ClickPower": "1",
			"IncOneString": autoIncOneString,
			"CoinPerSecond": coinPerSecond,
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
			"CoinPerSecond": coinPerSecond,

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
			"CoinPerSecond": coinPerSecond,

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
			"CoinPerSecond": coinPerSecond,

		})
		if err != nil {
			return
		}
	}
}



