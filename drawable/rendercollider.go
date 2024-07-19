package drawable

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/snopan/flappy-chick/component"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type RenderCollider struct {
	query *donburi.Query
}

func NewRenderCollider() *RenderCollider {
	return &RenderCollider{
		query: donburi.NewQuery(filter.Contains(component.RectangleCollider, transform.Transform)),
	}
}

func (r *RenderCollider) Draw(w donburi.World, screen *ebiten.Image) {
	r.query.Each(w, func(e *donburi.Entry) {
		collider := component.RectangleCollider.Get(e)
		position := transform.WorldPosition(e)

		topLeftX, topLeftY := collider.GetTopLeft(position.X, position.Y)
		vector.DrawFilledRect(
			screen,
			float32(topLeftX),
			float32(topLeftY),
			float32(collider.Width),
			float32(collider.Height),
			color.White,
			false,
		)
	})
}
