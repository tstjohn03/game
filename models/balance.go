package models

import (
	"html/template"
	"net/http"
)

var templates *template.Template

var bal = 0
var clickStatus = 0

func BalanceHandler(w http.ResponseWriter, r *http.Request){
	if clickStatus == 0 {
		bal++
	}
	if clickStatus == 1 {
		bal += 2
	}
	if clickStatus == 2 {
		bal += 4
	}
	if clickStatus == 3 {
		bal += 6
	}
	http.Redirect(w, r, "/", 302)
}

func GetBalance() int{
	return bal
}

func GetClickStatus() int{
	return clickStatus
}