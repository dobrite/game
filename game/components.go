package game

import (
	"github.com/nu7hatch/gouuid"
)

type entity struct {
	*uuid.UUID
}

type component interface {
	run()
}
