package system

import (
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type Tilt struct {
	query *donburi.Query
}

func NewTilt() *Tilt {
	return &Tilt{
		query: donburi.NewQuery(filter.Contains(component.TagPlayer, component.Velocity, transform.Transform)),
	}
}

func (t *Tilt) Update(w donburi.World) {
	t.query.Each(w, func(e *donburi.Entry) {
		velocity := component.Velocity.Get(e)
		transform := transform.Transform.Get(e)

		currentVelocity := velocity.Y
		if currentVelocity > 0 {
			currentVelocity = min(currentVelocity, options.MaxTiltVelocity)
		} else {
			currentVelocity = max(currentVelocity, -options.MaxTiltVelocity)
		}

		angle := options.MaxTiltAngle / options.MaxTiltVelocity * currentVelocity
		transform.LocalRotation = angle
	})
}
