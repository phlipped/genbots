package bot

import (
	"github.com/phlipped/genbots/brain"
)

type Bot struct {
	uuid   uint64
	brain  brain.Expression
	energy uint64
	retval brain.Value // last value returned by Think()
}

func NewRandom() Bot {
	b := Bot{
		brain: brain.NewRandomExpression2(),
	}
	return b
}

// Think causes the Bot to perform one iteration of Evaluation of its Brain
// Think will self-terminate under certain conditions, e.g.
//   - Running out of energy
//   - Too much recursion
// (Although at some stage they may be given the ability to catch
// Exceptions in their expression tree, in which case only Uncaught
// Exceptions will cause termination)
// It is possible that Think will never return, and yet the Bot will be healthy.
func (b *Bot) Think() {
	// FIXME actually build an environment properly
	env := brain.Environment{}

	b.retval = b.brain.Eval(
		brain.Context{
			Env:   &env,
			Depth: 0,
		},
	)
}

// Returns a copy of the original Brain
func (b *Bot) Copy() *Bot {
	return b
}

// Mutates the Brain in-place
func (b *Bot) Mutate() {

}
