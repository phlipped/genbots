package brain

import (
	"github.com/phlipped/genbots/arena"
)

// Environment includes access to the external world (e.g. the Arena,
// which may be needed to get physical values) and possibly other entities
// that a brain wants to interact with (e.g. whichever entity the Brain
// will end up sending Actions to - perhaps some kind of controller, or
// perhaps the Brain's own memory)
type Environment struct {
	a *arena.Arena
}
