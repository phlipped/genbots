package brain

// Value is the type of the value returned by Expression.Eval() and Function.Apply()
type Value float64

// Expression represents a computation of some kind.
// The <context> provides access to the external environment.
// THe <context> stores any variable definitions that are defined
// The context is passed from each Expression to any of its sub-expressions.
type Expression interface {
	Eval(context Context) Value
}

// A Mutable expression can be copied and mutated
// It still implements Expression as well
type Mutable interface {
	Expression
	Copy() Mutable
	Mutate()
}