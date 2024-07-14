package system

import (
	"fmt"

	"github.com/snopan/flappy-chick/animation"
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Animation struct {
	query *donburi.Query
}

func NewAnimation() *Animation {
	return &Animation{
		query: donburi.NewQuery(filter.Contains(component.Sprite, component.Animation)),
	}
}

func (a *Animation) Update(w donburi.World) {
	a.query.Each(w, func(e *donburi.Entry) {
		animation := component.Animation.Get(e)
		sprite := component.Sprite.Get(e)

		changeAnimation(animation)
		updateAnimation(animation)

		animationSprite, err := animation.CurrentSheet.GetFrame(animation.CurrentFrame)
		if err != nil {
			panic(fmt.Sprintf("animation update err: %s", err.Error()))
		}

		sprite.Image = animationSprite
	})
}

func changeAnimation(a *component.AnimationData) {
	if !a.ShouldChange {
		return
	}

	newSheet, ok := animation.AnimationLibrary[a.NextAnimation]
	if !ok {
		panic(fmt.Sprintf("invalid animation key: %d", a.NextAnimation))
	}

	a.CurrentSheet = newSheet
	a.CurrentFrame = 0

	a.ShouldChange = false
}

func updateAnimation(a *component.AnimationData) {
	a.CurrentFrameDuration++
	if a.CurrentFrameDuration < a.CurrentSheet.FrameDuration() {
		return
	}

	if a.CurrentFrame == a.CurrentSheet.EndFrame() {
		a.CurrentFrame = 0
	} else {
		a.CurrentFrame++
	}

	a.CurrentFrameDuration = 0
}
