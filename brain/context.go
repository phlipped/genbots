package brain

// Context contains everything 'external' to an expression that might
// be needed to evaluate an expression.
type Context struct {
  Depth uint64
  Env *Environment
  Args []Value
}

func (c Context) CopyClean() Context {
	return Context{
		Depth: c.Depth,
		Env: c.Env,
		Args: make([]Value, 0),
	}
}