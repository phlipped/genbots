package bot

import (
	"testing"
)

func TestThink(t *testing.T) {
	b := NewRandom()
	b.Think()
	t.Logf("%f", b.retval)
}

func TestCopy(t *testing.T) {

}

func TestMutate(t *testing.T) {

}
