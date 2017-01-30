package brain

type Context struct {
  Depth uint64
  Env *Environment
  Args map[string]Value
}