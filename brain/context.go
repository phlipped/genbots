package brain

// Context contains everything 'external' to an expression that might
// be needed to evaluate an expression.
type Context struct {
	Depth uint64
	Env   *Environment
	Args  []Closure
}

// CopyClean returns a copy of the Context but with the Args reset
// Use AddArg to add more arguments
func (c *Context) CopyClean() Context {
	return Context{
		Depth: c.Depth,
		Env:   c.Env,
		Args:  make([]Closure, 0),
	}
}

func (c *Context) AddArg(a Closure) {
	c.Args = append(c.Args, a)
}