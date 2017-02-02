package brain

type Constant struct {
	V Value
}

func (c *Constant) Eval(ctx Context) Value {
	return c.V
}

type If struct {
	pred Expression
	iftrue Expression
	iffalse Expression
}

func (i *If) Eval(ctx Context) Value {
	if i.pred.Eval(ctx) > 0 {
		return i.iftrue.Eval(ctx)
	}
	return i.iffalse.Eval(ctx)
}