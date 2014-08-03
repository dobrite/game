package game

import (
	"github.com/nu7hatch/gouuid"
)

type session struct {
	*connection
	id *uuid.UUID
}

func newSession(id *uuid.UUID, transport transport, toConn chan string, toGame chan string) *session {
	conn := &connection{
		trans:  transport,
		toConn: toConn,
		toGame: toGame,
	}

	sess := &session{
		conn,
		id,
	}

	// prob put this in connection
	// and create newConnection factory function
	sess.tomb.Go(sess.sender)
	sess.tomb.Go(sess.receiver)
	return sess
}
