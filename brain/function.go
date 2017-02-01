package brain

// A Function points to an Expression that takes Arguments
// Arguments must be added to the Context which is then
// passed to the Expression's Eval()
// This "Context modification" step is performed by a special
// Expression called a FunctionCaller.
// The Function struct itself isn't really needed at Eval time, but
// it is used for bookkeeping purposes - tracking functions, tracking how many
// FuncCallers exist for a particular Function
type Function struct {
	Exp Expression
	Arity uint64
}

// A FunctionCaller expression calls some Function with arguments
// FunctionCaller creates a new Context and places it's Eval'd children into the Args section of the context
// It then Eval's the Function's Expression 
// Any Expression that contains an "ArgLookup" expression somewhere in its immediate tree is considered
// to be a function.
type FunctionCaller struct {
  f *Function
  children []Expression // the FuncCaller's children, which will be Eval'd before Eval'ing the function
}

func (fc *FunctionCaller) Eval(ctx Context) Value {
	// FIXME check depth, and anything else we should check (kill signal from controller or arena perhaps)
	
	// Evaluate our children using our current context
	childvals := make([]Value, 0)
	for _, child := range fc.children {
      childvals = append(childvals, child.Eval(ctx))
	}

	// Build new context
	newctx := ctx
	newctx.Args = childvals

	return fc.f.Exp.Eval(newctx)
}

func (fc *FunctionCaller) Copy() Expression {
	return *fc
}

// ArgGetter retrieves corresponding Value from the Context
type ArgGetter struct {
	pos uint64
}

func (a *ArgGetter) Eval(ctx Context) Value {
  return ctx.Args[a.pos]
}

func (a *ArgGetter) Copy() Expression {
  return *a
}

func (a *ArgGetter) Mutate() {

}