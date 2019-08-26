package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/jinzhu/gorm"
	negronilogrus "github.com/meatballhat/negroni-logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// TODO represents single TODO item.
type TODO struct {
	Item  string
	Index uint
}

// TODOModel represents the model of a TODO list.
type TODOModel struct {
	ID   uint `gorm:"PRIMARY_KEY,AUTO_INCREMENT"`
	Todo string
}

// TableName set the name of the table.
func (TODOModel) TableName() string {
	return "todos"
}

var db *gorm.DB

func init() {
	var err error

	/* Use SQLite in memory for testing purpose. */
	db, err = gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&TODOModel{}) {
		db.CreateTable(&TODOModel{})
	}
}

func main() {
	defer db.Close()

	host := "127.0.0.1"
	port := "8080"

	args := os.Args[1:]

	for {
		if len(args) < 2 {
			break
		} else if args[0] == "-h" || args[0] == "--host" {
			host = args[1]

			args = args[2:]
		} else if args[0] == "-p" || args[0] == "--port" {
			port = args[1]

			args = args[2:]
		} else {
			log.Fatal(fmt.Sprintf("Unknown parameter: %s", args[0]))
		}
	}

	/* Set the routes for the web application. */
	mux := httprouter.New()

	// Listen to CSS assets
	mux.ServeFiles("/css/*filepath", http.Dir("public/css"))

	// Listen to JavaScript assets
	mux.ServeFiles("/js/*filepath", http.Dir("public/js"))

	// Listen to the index page.
	mux.GET("/", indexHandler)

	// Respond to new TODO item.
	mux.POST("/todo/", newTODOHandler)

	// Handle HTTP 404
	mux.NotFound = http.HandlerFunc(notFoundHandler)

	// Handle CORS policy
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	})

	/* Create the logger for the web application. */
	l := log.New()

	n := negroni.New()
	n.Use(c)
	n.Use(negronilogrus.NewMiddlewareFromLogger(l, "web"))
	n.UseHandler(mux)

	/* Create the main server object */
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: n,
	}

	log.Println(fmt.Sprintf("Run the web server at %s:%s", host, port))
	log.Fatal(server.ListenAndServe())
}

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

func newTODOHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()

	todo := r.FormValue("todo")
	method := r.FormValue("_method")

	if todo == "" {
		r.Header.Add("Message", "Empty TODO item")
	} else if method == "update" {
		index := r.FormValue("index")

		if index == "" {
			r.Header.Add("Message", "Unable to retrieve TODO item")
		} else {
			db.Table("todos").Where("id == ?", index).Update(struct {
				Todo string `gorm:"todo"`
			}{
				Todo: todo,
			})
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
