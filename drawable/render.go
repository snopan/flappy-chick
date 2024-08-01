package drawable

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/snopan/flappy-chick/component"
	"github.com/snopan/flappy-chick/options"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

type Render struct {
	query *donburi.Query
}

func NewRender() *Render {
	return &Render{
		query: donburi.NewQuery(filter.Contains(component.Sprite, transform.Transform)),
	}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	layers := map[int][]*donburi.Entry{}

	r.query.Each(w, func(e *donburi.Entry) {
		sprite := component.Sprite.Get(e)
		if sprite.Image == nil {
			return
		}

		_, ok := layers[sprite.Layer]
		if !ok {
			layers[sprite.Layer] = []*donburi.Entry{e}
			return
		}
		layers[sprite.Layer] = append(layers[sprite.Layer], e)
	})

	layerArr := make([]int, 0)
	for layer, _ := range layers {
		layerArr = append(layerArr, layer)
	}
	sort.Ints(layerArr)

	for _, layer := range layerArr {
		entities := layers[layer]

		for _, e := range entities {
			render(w, screen, e)
		}
	}
}

func render(w donburi.World, screen *ebiten.Image, e *donburi.Entry) {
	sprite := component.Sprite.Get(e)

	// Get relevant information
	position := transform.WorldPosition(e)
	rotation := transform.WorldRotation(e)
	scale := transform.WorldScale(e)
	size := sprite.Image.Bounds().Size()

	// Dont render if it's out of frame
	if !SpriteInFrame(math.NewVec2(float64(size.X), float64(size.Y)), position, scale) {
		return
	}

	// Center on top left so we can apply rotation and scaling
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(size.X)/2, -float64(size.Y)/2)

	// Apply all transformations
	op.GeoM.Rotate(rotation)
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)

	screen.DrawImage(sprite.Image, op)
}

func SpriteInFrame(size, position, scale math.Vec2) bool {
	topLeft := math.NewVec2(
		position.X-size.X/2*scale.X,
		position.Y-size.Y/2*scale.Y,
	)

	bottomRight := math.NewVec2(
		position.X+size.X/2*scale.X,
		position.Y+size.Y/2*scale.Y,
	)

	topLeftInWindow := (0.0 <= topLeft.X && topLeft.X <= options.WindowWidth) || (0.0 <= topLeft.Y && topLeft.Y <= options.WindowHeight)
	bottomRightInWindow := (0.0 <= bottomRight.X && bottomRight.X <= options.WindowWidth) || (0.0 <= bottomRight.Y && bottomRight.Y <= options.WindowHeight)

	return topLeftInWindow || bottomRightInWindow
}
