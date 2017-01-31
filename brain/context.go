package brain

// Context contains everything 'external' to an expression that might
// be needed to evaluate an expression.
type Context struct {
  Depth uint64
  Env *Environment
  Args map[string]Value
}