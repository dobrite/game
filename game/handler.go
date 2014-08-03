package game

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

type Handler struct {
	*registry
}

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	h := &Handler{
		registry: newRegistry(),
	}
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
	u4, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return
	}
	id := u4
	log.Printf("client connected: %s", id)
	session := newSession(id, transport, toConn, toGame)
	h.registry.add(session)
	h.registry.send(session, buildMessageConfig(id))
	h.registry.send(session, buildMessageWorld())
	p := position{
		x: 8,
		y: 8,
	}
	positions[id] = p
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

func (h *Handler) teardown(session *session) {
	log.Printf("client disconnected: %s", session.id)
	h.registry.remove(session)
}

func (h *Handler) handleMessage(msg message, session *session) {
	id := msg.id
	switch msg := msg.message.(type) {
	case messageMove:
		p := positions[id]
		positions[id] = position{
			x: p.x + msg.X,
			y: p.y + msg.Y,
		}
		h.registry.publish(buildMessageWorld())
		//h.registry.subscribe(msg, session)
		//h.registry.send(session, buildMessageSubscriptionSucceeded(msg.Channel))
	default:
		log.Fatal("I give up")
	}
}
