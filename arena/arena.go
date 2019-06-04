package arena

import (
	"context"
	"github.com/phlipped/genbots/bot"
	"github.com/phlipped/genbots/action"
	"github.com/google/uuid"
)

type Arena struct {
	width uint32
	height uint32
	bots map[uuid.UUID]*bot.Bot
}

func New(w, h uint32) *Arena {
	return &Arena{width: w, height: h}
}

func (a *Arena) Step() {
	// Eval each bot
	ctx, cancel := context.WithCancel(context.Background())
	actions := make(chan *action.Action, len(a.bots))
	for _, bot := range a.bots {
		go bot.Eval(ctx, actions)
	}

	// Now loop over all the actions coming in on bot action chan
	// and implemenet them.
	// At some stage, set a timelimit or execution limit to force
	// the results to come back (to be implemented)
	for action := range actions {
		_ = action
	}
	_ = cancel

}
