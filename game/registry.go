package game

import (
	"github.com/deckarep/golang-set"
	"github.com/nu7hatch/gouuid"
	"gopkg.in/tomb.v2"
)

//TODO move mapset to clipperhouse.github.io/gen
type registry struct {
	sessionIds mapset.Set
	sessions   map[*uuid.UUID]*session
	command    chan func()
	commands   map[string]func()
	t          tomb.Tomb
}

func newRegistry() *registry {
	registry := &registry{
		sessionIds: mapset.NewThreadUnsafeSet(),
		sessions:   make(map[*uuid.UUID]*session),
		command:    make(chan func()),
		commands:   make(map[string]func()),
	}
	registry.t.Go(registry.run)
	return registry
}

func (r *registry) run() error {
	for {
		select {
		case command := <-r.command:
			command()
		case <-r.t.Dying():
			return tomb.ErrDying
		}
	}
}

func (r *registry) publish(payload string) {
	r.command <- func() {
		s := r.sessionIds.Iter()
		for sId := range s {
			// whole server can be blocked by a slow client
			r.sessions[sId.(*uuid.UUID)].toConn <- payload
		}
	}
}

func (r *registry) add(session *session) {
	r.command <- func() {
		sId := session.id
		r.sessionIds.Add(sId)
		r.sessions[sId] = session
	}
}

func (r *registry) remove(session *session) {
	r.command <- func() {
		sId := session.id
		r.sessionIds.Remove(sId)
		delete(r.sessions, sId)
	}
}

func (r *registry) send(session *session, payload string) {
	r.command <- func() {
		session.toConn <- payload
	}
}
