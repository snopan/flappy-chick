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

		size := sprite.Image.Bounds().Size()

		transform := component.Transform.Get(e)
		op := &ebiten.DrawImageOptions{}

		// Center on top left so we can apply rotation and scaling
		op.GeoM.Translate(-float64(size.X)/2, -float64(size.Y)/2)
		op.GeoM.Rotate(transform.Rotation.Theta)
		op.GeoM.Scale(transform.Scale.X, transform.Scale.Y)

		// Then move the image to it's coordinates
		op.GeoM.Translate(transform.Position.X, transform.Position.Y)

		screen.DrawImage(sprite.Image, op)
	})
}
