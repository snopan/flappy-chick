package component

import "github.com/yohamta/donburi"

type TransformData struct {
	Position PositionData
	Scale    ScaleData
	Rotation RotationData
}

type PositionData struct {
	X, Y float64
}

type ScaleData struct {
	X, Y float64
}

type RotationData struct {
	Theta float64
}

var Transform = donburi.NewComponentType[TransformData]()
