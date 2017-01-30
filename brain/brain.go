package brain

type Value float64

type Brain struct {
  head Expression
}

// Think causes the Brain to perform one iteration of Evaluation
// The Brain will self-terminate under certain conditions, e.g.
//   - Running for too long
//   - Too much recursion
// (Although at some stage they may be given the ability to catch
// Exceptions in their expression tree, in which case only Uncaught
// Exceptions will cause termination)
func (b *Brain) Think(env *Environment) {
    b.head.Eval(
  		Context{
  			Env: env,
  			Depth: 0,
  		}
  	)
}

// Returns a copy of the original Brain
func (b *Brain) Duplicate() *Brain {
  return b
}

// Mutates the Brain in-place
func (b *Brain) Mutate() {

}