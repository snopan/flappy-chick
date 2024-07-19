package component

import "github.com/yohamta/donburi"

type VelocityData struct {
	X, Y float64
}

var Velocity = donburi.NewComponentType[VelocityData]()

func SetVelocity(e *donburi.Entry, x, y float64) {
	donburi.SetValue(e, Velocity, VelocityData{
		X: x,
		Y: y,
	})
}
