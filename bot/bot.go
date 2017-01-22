package bot

import (
  "github.com/phlipped/genbots/brain"
)

type Bot struct {
  uuid uint64
  brain brain.Brain
  energy uint64
}

func (*b Bot) Eval() {

}