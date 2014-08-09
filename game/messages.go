package game

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"log"
)

type message struct {
	id      *uuid.UUID
	message interface{}
}

type messageEvent struct {
	Event string          `json:"event"`
	Id    string          `json:"id"`
	Data  json.RawMessage `json:"data"`
}

type messageMove struct {
	Y int `json:"y"`
	X int `json:"x"`
}

func buildMessageConfig(id *uuid.UUID) string {
	log.Println(id)
	wc := &wireConfig{
		Event:  "game:config",
		ChunkX: chunkX,
		ChunkY: chunkY,
		WorldX: worldX,
		WorldY: worldY,
		Id:     id.String(),
	}

	c, _ := json.Marshal(wc)
	return string(c)
}

func buildMessageWorld() string {
	ww := &wireWorld{
		Event: "game:world",
		Data:  blah(),
	}

	w, _ := json.Marshal(ww)
	return string(w)
}

func blah() [][]*chunk {
	base := make([]*chunk, worldY*worldX)
	outer := make([][]*chunk, worldY)
	for i := range outer {
		outer[i] = base[i*worldX : (i+1)*worldX]
		for j := range outer[i] {
			outer[i][j] = w[j][i]
		}
	}
	return outer
}

func MessageUnmarshalJSON(b []byte) (msg message, err error) {
	var event messageEvent
	err = json.Unmarshal(b, &event)
	if err != nil {
		return
	}
	switch event.Event {
	case "game:move":
		var msgMove messageMove
		err = json.Unmarshal(event.Data, &msgMove)
		if err != nil {
			return
		}
		var u4 *uuid.UUID
		u4, err = uuid.ParseHex(event.Id)
		if err != nil {
			return
		}
		msg = message{
			message: msgMove,
			id:      u4,
		}
	default:
		err = fmt.Errorf("%s is not a recognized event", event.Event)
	}
	return
}
