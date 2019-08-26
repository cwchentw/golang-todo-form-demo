package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	negronilogrus "github.com/meatballhat/negroni-logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

var db *gorm.DB

func main() {
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

	/* Use SQLite in memory for testing purpose. */
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/* Set the routes for the web application. */
	mux := httprouter.New()
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
