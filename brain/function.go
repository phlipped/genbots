package brain

// Closure wraps up an Expression with its context so that it can be paassed
// as an Argument to some other Function without losing its context.
type Closure struct {
	Exp Expression
	Ctx Context
}

// Eval on a Closure is like Eval for Expressions, but doesn't need a Context
func (c Closure) Eval() Value {
	return c.Exp.Eval(c.Ctx)
}

// FuncCaller is used to call other Expressions as Functions.
// The number of arguments in the FuncCaller must match the
// expected number of arguments in the called Expression.
type FuncCaller struct {
	f Expression
	args []Expression
}

func (fc *FuncCaller) Eval(ctx Context) Value {
	newctx := ctx.CopyClean()
	for _, a := range fc.args {
		newctx.AddArg(Closure{a, ctx})
	}
	return fc.f.Eval(newctx)
}

// GetArg retrieves a Closure from the Context
type GetArg struct {
	pos uint64
}

func (g GetArg) Eval(ctx Context) Value {
	return ctx.Args[g.pos].Eval()
}

// I'm toying with the idea of having a "SetArg" and/or "AppendArg" such that
// Nodes can modify their current context