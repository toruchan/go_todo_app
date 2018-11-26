package controller

import (
	"../model"
)

type Todo struct {
}

func NewTask() Todo {
   return Todo{}
 }

// func IndexGET(w http.ResponseWriter, r *http.Request) {
// 	tpl := template.New("fieldname example")
// 	todos := models.SelectAll()
// 	tpl = template.Must(template.ParseFiles("gosample/index.tpl"))
// 	tpl.Execute(w,todos)
// }

// 全てのタスクを取得
func GetAll() interface{} {
  todos := model.SelectAll()
  return todos
}

func AddTask(todo string) {
  model.InsertTask(todo)
}
