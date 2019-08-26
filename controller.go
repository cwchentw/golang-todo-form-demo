package main

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var tmpl = template.Must(
		template.ParseFiles("views/layout.html", "views/index.html", "views/head.html"),
	)

	var msg string

	if r.Header.Get("Message") != "" {
		msg = r.Header.Get("Message")

		// Clean current message.
		r.Header.Set("Message", "")
	}

	rows, err := db.Table("todos").Select("*").Rows()
	if err != nil {
		msg = "Unable to retrieve database"
	}

	var todos []TODO

	todos = make([]TODO, 0)

	for rows.Next() {
		var todo struct {
			ID   uint
			Todo string `gorm:"todo"`
		}

		db.ScanRows(rows, &todo)

		todos = append(todos, TODO{
			Index: todo.ID,
			Item:  todo.Todo,
		})
	}

	data := struct {
		Title   string
		TODOs   []TODO
		Message string
	}{
		Title:   "TODO List",
		TODOs:   todos,
		Message: msg,
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func updateTODOHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()

	todo := r.FormValue("todo")
	method := r.FormValue("_method")

	if todo == "" {
		r.Header.Set("Message", "Empty TODO item")
	} else if method == "update" {
		index := r.FormValue("index")

		if index == "" {
			r.Header.Set("Message", "Unable to retrieve TODO item")
		} else {
			db.Table("todos").Where("id == ?", index).Update(struct {
				Todo string `gorm:"todo"`
			}{
				Todo: todo,
			})
		}
	} else if method == "delete" {
		index := r.FormValue("index")

		if index == "" {
			r.Header.Set("Message", "Unable to retrieve TODO item")
		} else {
			db.Table("todos").Where("id == ?", index).Delete(struct {
				ID   uint
				Todo string
			}{})
		}
	} else {
		db.Table("todos").Create(struct {
			Todo string `gorm:"todo"`
		}{
			Todo: todo,
		})
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
