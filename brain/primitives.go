package brain

// Constant
type Constant struct {
	V Value
}

func (c *Constant) Eval(_ Context) Value {
	return c.V
}

// PrimitiveFunc
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

// If
var If = &PrimitiveFunc{
	op: func (args ...Closure) Value {
			if args[0].Eval() > 0 {
				return args[1].Eval()
			}
			return args[2].Eval()	
		},
	arity: 3,
}

// Mul
var Mul = &PrimitiveFunc{
	op: func (args ...Closure) Value {
		return args[0].Eval() * args[1].Eval()
	},
	arity: 2,
}

// Sub
var Sub = &PrimitiveFunc{
	op: func (args ...Closure) Value {
		return args[0].Eval() - args[1].Eval()
	},
	arity: 2,
}

// While
var While = &PrimitiveFunc{
	op: func (args ...Closure) (final Value) {
		for args[0].Eval() > 0 {
			final = args[1].Eval()
		}
		return final
	},
	arity: 2,
}