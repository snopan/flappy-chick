package animation

import (
	"fmt"

	"github.com/snopan/flappy-chick/helper/spritesheet"
)

type AnimationKey int

var animationLibrary = map[AnimationKey]*Animation{}

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
		frameDuration:  5,
	},
}

func InitAnimationLibrary() {
	for animationKey, data := range animationData {
		spriteSheet := spritesheet.GetSpriteSheet(data.spriteSheetKey)
		animationLibrary[animationKey] = FromSpriteSheet(spriteSheet, data.frameDuration)
	}
}

func GetAnimation(key AnimationKey) *Animation {
	animation, ok := animationLibrary[key]
	if !ok {
		panic(fmt.Sprintf("failed to get animation: %d", key))
	}

	return animation
}
