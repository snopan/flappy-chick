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
		query: donburi.NewQuery(filter.Contains(component.Sprite)),
	}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	r.query.Each(w, func(e *donburi.Entry) {
		sprite := component.Sprite.Get(e)
		if sprite.Image == nil {
			return
		}

		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(sprite.Image, op)
	})
}
