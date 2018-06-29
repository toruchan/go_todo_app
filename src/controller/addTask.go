package controller

import (
    "net/http"
    "models"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := r.Form["task"]
	models.InsertTask(todo[0])
	http.Redirect(w, r, "/", http.StatusFound)

}
