package game

import (
	"github.com/nu7hatch/gouuid"
)

type entity *uuid.UUID

type component interface {
	run()
}
