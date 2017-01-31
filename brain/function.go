package brain

// A Function takes 0 or more Values as arguments and returns a Value
// Although a Function that takes 0 Values isn't really any different from an Expression
// The main method of a Function is 'Apply()' which is used 'at runtime'. The other
// methods are 'utility' methods that are used to manipulate the Function's internals
// In the future, they may be moved to a separate Interface
type Function interface {
	Apply(env *Environment, depth uint64, args ...Value) Value
	Arity() uint64 // How many args are needed by the function
}
