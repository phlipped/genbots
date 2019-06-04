package brain

import (
	"context"
	"github.com/phlipped/genbots/action"
)

// These interface definitions don't follow
// the normal Go naming convention. But so what.

type Actioner interface {
	Eval(context.Context, Env) *action.Action
	Copy() Actioner
	Mutate()
}

type Booler interface {
	Eval(context.Context, Env) bool
	Copy() Booler
	Mutate()
}

type Inter interface {
	Eval(context.Context, Env) int32
	Copy() Inter
	Mutate()
}
