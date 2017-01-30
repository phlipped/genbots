package brain

// Expression represents a computation of some kind.
// The <context> provides access to the external environment.
// THe <context> stores any variable definitions that are defined
// The context is passed from each Expression to any of its sub-expressions.
type Expression interface {
	Eval(context Context) Value
	Mutate() Expression  // Not clear at this point whether we mutate ourselves, or instead return a mutant copy.
}