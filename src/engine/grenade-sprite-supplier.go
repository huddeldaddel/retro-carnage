package engine

import (
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
)

const grenadeImagePath = "images/weapons/grenade.png"

var (
	grenadeSprite *graphics.SpriteWithOffset
)

// GrenadeSpriteSupplier is used to provide sprites for grenades.
type GrenadeSpriteSupplier struct{}

func (gss *GrenadeSpriteSupplier) Sprite(int64) *graphics.SpriteWithOffset {
	if nil == grenadeSprite {
		var sprite = assets.SpriteRepository.Get(grenadeImagePath)
		grenadeSprite = &graphics.SpriteWithOffset{
			Offset: geometry.Point{},
			Source: grenadeImagePath,
			Sprite: sprite,
		}
	}
	return grenadeSprite
}
