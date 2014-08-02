package main

import (
	"fmt"
	"github.com/dobrite/game/game"
	//"github.com/gorilla/websocket"
	"log"
	"net/http"
	//"time"
)

func newServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("sock/", websocket)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.HandleFunc("/", IndexHandler)

	return mux
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./public/index.html")
}

func main() {
	g := &game.Game{}
	g.Init()

	mux := newServeMux()

	log.Fatal(http.ListenAndServe(":3000", mux))

	var i string
	_, _ = fmt.Scanf("%s", &i)
}
