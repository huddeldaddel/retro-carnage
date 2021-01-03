package characters

import (
	"retro-carnage/engine/graphics"
)

type EnemySpriteSupplier interface {
	Sprite(elapsedTimeInMs int64, enemy Enemy) *graphics.SpriteWithOffset
}
