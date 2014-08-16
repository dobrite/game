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
	Z int `json:"z"`
	X int `json:"x"`
}

type messageConfig struct {
	Event  string `json:"event"`
	ChunkZ int    `json:"chunkZ"`
	ChunkX int    `json:"chunkX"`
	ChunkY int    `json:"chunkY"`
	LosZ   int    `json:"losZ"`
	LosX   int    `json:"losX"`
	LosY   int    `json:"losY"`
	Id     string `json:"id"`
}

func buildMessageConfig() string {
	wc := &messageConfig{
		Event:  "game:config",
		ChunkZ: chunkZ,
		ChunkX: chunkX,
		ChunkY: chunkY,
		LosZ:   losZ,
		LosX:   losX,
		LosY:   losY,
	}
	c, _ := json.Marshal(wc)
	return string(c)
}

func buildMessageWorld(cc chunkCoords) string {
	msg, _ := json.Marshal(w.toJSON(cc))
	return string(msg)
}

type messageSpawn struct {
	Event  string `json:"event"`
	Id     string `json:"id"`
	Z      int    `json:"z"`
	X      int    `json:"x"`
	ChunkZ int    `json:"cz"`
	ChunkX int    `json:"cx"`
}

func buildMessageSpawn(ent string) string {
	id := ent
	position := positionsSet.byEnt(ent)
	ms := &messageSpawn{
		Event:  "game:spawn",
		Id:     id,
		Z:      position.z,
		X:      position.x,
		ChunkZ: position.cz,
		ChunkX: position.cx,
	}
	msg, _ := json.Marshal(ms)
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
