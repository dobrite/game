package game

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

type handler struct {
	*registry
}

func (h *handler) websocket(w http.ResponseWriter, req *http.Request) {
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

func (h *handler) handle(transport transport) {
	toGame := make(chan string)
	toConn := make(chan string)
	u4, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
		return
	}
	id := u4.String()
	log.Printf("client connected: %s", id)
	session := newSession(id, transport, toConn, toGame)
	h.registry.add(session)
	h.registry.send(session, buildMessageConfig())
	go h.listen(session)
}

func (h *handler) listen(session *session) {
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

func (h *handler) teardown(session *session) {
	log.Printf("client disconnected: %s", session.id)
	h.registry.remove(session)
}

func (h *handler) handleMessage(msg message, session *session) {
	switch msg := msg.(type) {
	case messageMove:
		log.Println("message move " + msg.Direction)
		h.registry.subscribe(msg, session)
		h.registry.send(session, buildMessageSubscriptionSucceeded(msg.Channel))
	default:
		log.Fatal("I give up")
	}
}
