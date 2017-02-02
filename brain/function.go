package brain

type Closure struct {
	Exp Expression
	Ctx Context
}

// Eval on a Closure ignores the Context that is passed in
func (c Closure) Eval() Value {
	return c.Exp.Eval(c.Ctx)
}

type FuncCaller struct {
	f Expression
	args []Expression
}

func (fc *FuncCaller) Eval(ctx Context) Value {
	newctx := ctx.CopyClean()
	for _, a := range fc.args {
		newctx.AddArg(Closure{Exp: a, Ctx: ctx})
	}
	return fc.f.Eval(newctx)
}

type GetArg struct {
	pos uint64
}

func (g GetArg) Eval(ctx Context) Value {
	return ctx.Args[g.pos].Eval()
}