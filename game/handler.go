package game

import (
	"log"
	"net/http"
)

type Handler struct{}

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	h := &Handler{}
	mux.HandleFunc("/sock/", h.Websocket)
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.HandleFunc("/", IndexHandler)

	return mux
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("index")
	http.ServeFile(w, req, "./public/index.html")
}

func (h *Handler) Websocket(w http.ResponseWriter, req *http.Request) {
	log.Println("websocket connected")
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	h.handle(newWebsocketTransport(ws))
}

func (h *Handler) handle(transport transport) {
	toGame := make(chan string)
	toConn := make(chan string)

	id := newUUID()
	positions.add(id, 8, 8, 0, 0)
	materials.add(id, flesh)

	session := newSession(id, transport, toConn, toGame)

	reg.add(session)
	reg.send(session, buildMessageConfig(id))

	log.Printf("client connected: %s", id)

	materials.byType(flesh)
	go h.listen(session)
}

func (h *Handler) listen(session *session) {
	for {
		if raw, ok := <-session.toGame; ok {
			msg, err := MessageUnmarshalJSON([]byte(raw))

			if err != nil {
				log.Println("error unmarshaling json: " + err.Error())
				//break
			} else {
				h.handleMessage(msg, session)
			}
		} else {
			break
		}
	}
	h.teardown(session)
}

func (h *Handler) handleMessage(msg message, session *session) {
	id := msg.id
	switch msg := msg.message.(type) {
	case messageMove:
		reg.commands[id.String()] = func() {
			p := positions[id.String()]
			p.x = p.x + msg.X
			p.y = p.y + msg.Y
			positions[id.String()] = p
		}
	default:
		log.Fatal("I give up")
	}
}

func (h *Handler) teardown(session *session) {
	log.Printf("client disconnected: %s", session.id)
	reg.remove(session)
	// TODO also remove them from all component maps
}
