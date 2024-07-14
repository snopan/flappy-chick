package component

import (
	"github.com/snopan/flappy-chick/animation"
	"github.com/yohamta/donburi"
)

type AnimationData struct {
	CurrentSheet         *animation.Animation
	CurrentFrame         int
	CurrentFrameDuration int
	NextAnimation        animation.AnimationKey
	ShouldChange         bool
}

var Animation = donburi.NewComponentType[AnimationData]()
