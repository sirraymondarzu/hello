// create a route  for greeting
// create a handler greeting that will output welcome to my page
package main

import (
	"context"
	"database/sql"
	"flag"     // allows you to co
	"log"      // display to terminal
	"net/http" // create a multiplexer and create routes/endpoints
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// create a new type
type application struct{}

func main() {
	// create a flag for specifing the port number when starting the server

	addr := flag.String("port", ":4000", "HTTP Network address") //if no port assigned // stored as a pointer
	dsn := flag.String("dsn", os.Getenv("APPLETREE_DB_DSN"), "PostgresSQL DSN")
	flag.Parse() // this should be called only once
	//create instance of application type

	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	app := &application{} // creating app type
	defer db.Close()
	log.Println("datase connection pool established.")

	//create a customized server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Server is active on port %s", *addr)

	err = srv.ListenAndServe() // using the pointer of addr -
	log.Fatal(err)

}

// Get a database connection pool (2/16/2023)

// this a pointer to the database
// get a databae connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// use a context to check if the db is reachable (Ping the database) make it safe
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // has  a state it is keeping track of
	defer cancel()                                                          // this is always done ... if something was wrong... always execute as the last thing

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
