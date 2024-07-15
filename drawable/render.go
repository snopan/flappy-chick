package drawable

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type Render struct {
	query *donburi.Query
}

func NewRender() *Render {
	return &Render{
		query: donburi.NewQuery(filter.Contains(component.Sprite, component.Transform)),
	}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.query.Each(w, func(e *donburi.Entry) {
		sprite := component.Sprite.Get(e)
		if sprite.Image == nil {
			return
		}

		transform := component.Transform.Get(e)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(transform.Position.X/transform.Scale.X, transform.Position.Y/transform.Scale.Y)
		op.GeoM.Scale(transform.Scale.X, transform.Scale.Y)

		screen.DrawImage(sprite.Image, op)
	})
}
