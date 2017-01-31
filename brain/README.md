# genbots.brain

The purpose of a brain is to produce Actions (not yet decided where the actions will be posted to)
A brain is composed of one or more brain.Nodes
The top level Node of a brain will also implement brain.Expression.
Calling Eval() on the Expression will cause (possibly) Actions to be generated.