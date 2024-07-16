package sprite

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/helper/spritesheet"
)

type SpriteKey int

var spriteLibrary = map[SpriteKey]*ebiten.Image{}

const (
	GroundTop1 SpriteKey = iota
	GroundTop2
	GroundTop3
	Dirt
	DirtLightLine
	DirtDarkLine
)

var spriteData = map[SpriteKey]struct {
	spriteSheetKey spritesheet.SpriteSheetKey
	row            int
	col            int
}{
	GroundTop1: {
		spriteSheetKey: spritesheet.Ground,
		row:            0,
		col:            0,
	},
	GroundTop2: {
		spriteSheetKey: spritesheet.Ground,
		row:            0,
		col:            1,
	},
	GroundTop3: {
		spriteSheetKey: spritesheet.Ground,
		row:            0,
		col:            2,
	},
	Dirt: {
		spriteSheetKey: spritesheet.Ground,
		row:            1,
		col:            0,
	},
	DirtLightLine: {
		spriteSheetKey: spritesheet.Ground,
		row:            1,
		col:            1,
	},
	DirtDarkLine: {
		spriteSheetKey: spritesheet.Ground,
		row:            1,
		col:            2,
	},
}

func InitSpriteLibrary() error {
	for key, data := range spriteData {

		sprite, err := spritesheet.GetSpriteSheet(data.spriteSheetKey).GetFrame(data.row, data.col)
		if err != nil {
			return fmt.Errorf("sprite sheet %d unable to get frame: %w", data.spriteSheetKey, err)
		}

		spriteLibrary[key] = sprite
	}

	return nil
}

func GetSprite(key SpriteKey) *ebiten.Image {
	sprite, ok := spriteLibrary[key]
	if !ok {
		panic(fmt.Sprintf("failed to get sprite: %d", key))
	}

	return sprite
}
