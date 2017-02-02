package brain

type Constant struct {
	V Value
}

func (c *Constant) Eval(ctx Context) Value {
	return c.V
}

type PrimitiveFunc struct {
	op func (...Closure) Value
	arity uint64
}

func (p *PrimitiveFunc) Eval(ctx Context) Value {
	return p.op(ctx.Args...)
}

func (p *PrimitiveFunc) Arity() uint64 {
	return p.arity
}

var If = PrimitiveFunc{
	op: func (args ...Closure) Value {
			if args[0].Eval() > 0 {
				return args[1].Eval()
			}
			return args[2].Eval()	
		},
	arity: 3,
}

var Mul = PrimitiveFunc{
	op: func (args ...Closure) Value {
		return args[0].Eval() * args[1].Eval()
	},
	arity: 2,
}