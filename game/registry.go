package game

import (
	"github.com/deckarep/golang-set"
	"gopkg.in/tomb.v2"
)

//TODO move mapset to clipperhouse.github.io/gen
type registry struct {
	sessionIds mapset.Set
	sessions   map[string]*session
	command    chan func()
	commands   map[string]func()
	t          tomb.Tomb
}

func newRegistry() *registry {
	registry := &registry{
		sessionIds: mapset.NewThreadUnsafeSet(),
		sessions:   make(map[string]*session),
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

// all but initiator
func (r *registry) publish(session *session, payload string) {
	r.command <- func() {
		s := r.sessionIds.Iter()
		for sId := range s {
			if sId != session.id {
				// whole server can be blocked by a slow client
				r.sessions[sId.(string)].toConn <- payload
			}
		}
	}
}

// broadcast - everyone
func (r *registry) broadcast(payload string) {
	r.command <- func() {
		s := r.sessionIds.Iter()
		for sId := range s {
			// whole server can be blocked by a slow client
			r.sessions[sId.(string)].toConn <- payload
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
