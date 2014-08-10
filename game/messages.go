package game

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type message struct {
	id      string
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

func buildMessageConfig(id string) string {
	wc := &wireConfig{
		Event:  "game:config",
		ChunkX: chunkX,
		ChunkY: chunkY,
		WorldX: worldX,
		WorldY: worldY,
		Id:     id,
	}

	c, _ := json.Marshal(wc)
	return string(c)
}

func buildMessageWorld() string {
	msg, _ := json.Marshal(w.toJSON())
	return string(msg)
}

func buildMessageItem(ent string) string {
	i := &item{
		id:       ent,
		position: positionsSet.byEnt(ent),
		material: materialsSet.byEnt(ent),
	}
	msg, _ := json.Marshal(i.toJSON())
	return string(msg)
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
			id:      u4.String(),
		}
	default:
		err = fmt.Errorf("%s is not a recognized event", event.Event)
	}
	return
}
