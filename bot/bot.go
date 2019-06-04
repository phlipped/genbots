package bot

import (
	"context"
	"github.com/phlipped/genbots/brain"
	"github.com/phlipped/genbots/action"
	"github.com/google/uuid"
)

type Bot struct {
	Id uuid.UUID
	Brain brain.Actioner
	X int32
	Y int32
	Energy int64
	env *Env
}

func NewBot() *Bot {
	// FIXME Generate a new brain
	return &Bot{Id: uuid.New()}
}

func (b *Bot) Copy() *Bot {
	return &Bot{
		Id: uuid.New(),
		Brain: b.Brain.Copy(),
		Energy: b.Energy,
	}
}

func (b *Bot) Mutate() {
	b.Brain.Mutate()
}

func (b *Bot) Eval(ctx context.Context, actionChan chan<- *action.Action) {
	// Following inline func runs Brain.Eval but recovers from a panic
	// and sets the return value to an empty Action.
	action := func() (a *action.Action) {
		defer func() {
			if r := recover(); r != nil {
				a = &action.Action{}
			}
		}()
		return b.Brain.Eval(ctx, b.env)
	}()
	action.BotId = b.Id
	actionChan <- action
}
