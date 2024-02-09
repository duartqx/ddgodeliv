package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"ddgodeliv/api/server"
)

func getDb(conn string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", conn)
}

func main() {

	db, err := getDb(os.Getenv("POSTGRES_CONN_STR"))
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	mux := server.GetNewServer(
		server.ServerConfig{
			Db: db, Secret: []byte(os.Getenv("SECRET_KEY")),
		},
	).Build()

	var addr string = fmt.Sprintf(
		"%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"),
	)

	srv := &http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Listening and Serving at:", addr)

	c := make(chan os.Signal, 1)
	// Graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)

	// Block until we interrupt signal
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(ctx)

	os.Exit(0)
}
