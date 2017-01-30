package brain

// A Function takes 0 or more Values and returns a Value
type Function interface {
	Apply(env *Environment, depth uint64, args ...Value) Value
}
