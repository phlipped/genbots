package action

import (
	"github.com/google/uuid"
)

type Direction uint8
type Scent struct{
	strength uint8
	S1 uint8
	S2 uint8
	S3 uint8
}

type GiveFood struct {
	Amount uint32
	Dir Direction
}

type Action struct {
  BotId uuid.UUID // Id of bot taking this action
  Move uint32 // distance forward
  Turn Direction // turn to absolute direction
  Reproduce uint8 // proportion of size of offspring
  Eat Direction
  Attack Direction
  Grow uint32 // Add size -> increases stomach size, move cost, attack strength, defense strength, push strength, max speed, mouthful size
  Push Direction // tries to move an adjacent object
  DropScent Scent // Drops a scent on the ground. These are transparent signals that degrade over time.
  GiveFood GiveFood // Transfer energy to the entity in the corresponding direction.
}

