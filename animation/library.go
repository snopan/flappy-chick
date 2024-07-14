package animation

import (
	"fmt"

	"github.com/snopan/flappy-chick/spritesheet"
)

type AnimationKey int

var AnimationLibrary = map[AnimationKey]*Animation{}

const (
	PlayerDie AnimationKey = iota
	PlayerFall
	PlayerFly
)

var animationData = map[AnimationKey]struct {
	spriteSheetKey spritesheet.SpriteSheetKey
	frameDuration  int
}{
	PlayerDie: {
		spriteSheetKey: spritesheet.ChickenDie,
		frameDuration:  3,
	},
	PlayerFall: {
		spriteSheetKey: spritesheet.ChickenFall,
		frameDuration:  5,
	},
	PlayerFly: {
		spriteSheetKey: spritesheet.ChickenFly,
		frameDuration:  3,
	},
}

func InitAnimationLibrary() error {
	for animationKey, data := range animationData {
		spriteSheet, ok := spritesheet.SpriteSheetLibrary[data.spriteSheetKey]
		if !ok {
			return fmt.Errorf("sprite sheet %d is not loaded", data.spriteSheetKey)
		}

		AnimationLibrary[animationKey] = FromSpriteSheet(spriteSheet, data.frameDuration)
	}

	return nil
}
