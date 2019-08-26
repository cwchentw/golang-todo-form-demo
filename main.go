package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	negronilogrus "github.com/meatballhat/negroni-logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/julienschmidt/httprouter"
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
	mux.POST("/todo/", updateTODOHandler)

	// Handle HTTP 404
	mux.NotFound = http.HandlerFunc(notFoundHandler)

	/* Create the logger for the web application. */
	l := log.New()

	n := negroni.New()
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
