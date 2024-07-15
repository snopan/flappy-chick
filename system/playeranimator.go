package system

import (
	"github.com/snopan/flappy-chick/animation"
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerAnimator struct {
	query *donburi.Query
}

func NewPlayerAnimator() *PlayerAnimator {
	return &PlayerAnimator{
		query: donburi.NewQuery(filter.Contains(component.PlayerTag, component.Animation, component.Velocity)),
	}
}

func (p *PlayerAnimator) Update(w donburi.World) {
	p.query.Each(w, func(e *donburi.Entry) {
		animationComponent := component.Animation.Get(e)
		velocity := component.Velocity.Get(e)

		if animationComponent.ShouldChange {
			return
		}

		if velocity.Y < 0 {
			if animationComponent.NextAnimation == animation.PlayerFly {
				return
			}

			animationComponent.NextAnimation = animation.PlayerFly
		} else {
			if animationComponent.NextAnimation == animation.PlayerFall {
				return
			}

			animationComponent.NextAnimation = animation.PlayerFall
		}

		animationComponent.ShouldChange = true
	})
}
