package controller

import (
	"net/http"
	"html/template"
	"models"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	tpl := template.New("fieldname example")
	todos := models.SelectAll()
	tpl = template.Must(template.ParseFiles("gosample/index.tpl"))
	tpl.Execute(w,todos)
}
