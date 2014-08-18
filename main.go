package main

import (
	_ "bitbucket.org/liamstask/goose/lib/goose" // for godeps
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/dobrite/game/game"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func setupLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	setupLogger()

	db, err := sql.Open("postgres", "user=game dbname=game password=badgam4 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}

	g := &game.Game{}

	g.Init(dbmap)

	mux := game.NewServeMux()

	log.Fatal(http.ListenAndServe(":3000", mux))
}
