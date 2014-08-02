package game

import (
	"encoding/json"
	"fmt"
)

type message interface{}

type messageEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type messageMove struct {
	Direction string `json:"dir"`
}

func buildMessageConfig() string {
	wc := &wireConfig{
		Chunk_x: Chunk_x,
		Chunk_y: Chunk_y,
	}

	c, _ := json.Marshal(wc)
	return string(c)
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
		msg = msgMove
	default:
		err = fmt.Errorf("%s is not a recognized event", event.Event)
	}
	return
}
