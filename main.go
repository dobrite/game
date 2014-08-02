package main

import (
	"github.com/dobrite/game/game"
	//"github.com/gorilla/websocket"
	"log"
	"net/http"
	//"time"
)

func setupLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	setupLogger()
	g := &game.Game{}
	g.Init()

	mux := game.NewServeMux()

	log.Fatal(http.ListenAndServe(":3000", mux))
}
