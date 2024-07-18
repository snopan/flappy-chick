package component

import "github.com/yohamta/donburi"

type Anchor int

const (
	AnchorBottomLeft Anchor = iota
	AnchorBottomRight
	AnchorTopLeft
	AnchorTopRight
	AnchorCenter
)

type RectangleColliderData struct {
	Width  float64
	Height float64
	Anchor Anchor
}

func (r *RectangleColliderData) GetTopLeft(AnchorX, AnchorY float64) (float64, float64) {
	switch r.Anchor {
	case AnchorBottomLeft:
		return AnchorX, AnchorY - r.Height
	case AnchorBottomRight:
		return AnchorX - r.Width, AnchorY - r.Height
	case AnchorTopLeft:
		return AnchorX, AnchorY
	case AnchorTopRight:
		return AnchorX - r.Width, AnchorY
	case AnchorCenter:
		return AnchorX - r.Width/2.0, AnchorY - r.Height/2.0
	}
	panic("invalid anchor")
}

func (r *RectangleColliderData) GetBottomRight(AnchorX, AnchorY float64) (float64, float64) {
	switch r.Anchor {
	case AnchorBottomLeft:
		return AnchorX + r.Width, AnchorY
	case AnchorBottomRight:
		return AnchorX, AnchorY
	case AnchorTopLeft:
		return AnchorX + r.Width, AnchorY + r.Height
	case AnchorTopRight:
		return AnchorX, AnchorY + r.Height
	case AnchorCenter:
		return AnchorX + r.Width/2.0, AnchorY + r.Height/2.0
	}
	panic("invalid anchor")
}

var RectangleCollider = donburi.NewComponentType[RectangleColliderData]()
