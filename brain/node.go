package brain

// Node is the interface that provides 'utility' functions
// for manipulating Expressions and/or Functions, without
// needing to know whether the Node is an Expression or Function.
// Although every implementor Expression or Function is expected to
// also implement Node, the interfaces are kept separate
// to clarify code that manipulates Nodes, as opposed to code that
// Evals and Applies Expression and Functions.
type Node interface {
	Copy() Node // Returns a Deep copy of the Node
	Mutate() // Mutates the Node (in place).
}