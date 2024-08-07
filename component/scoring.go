package component

import "github.com/yohamta/donburi"

type ScoringData struct {
	Score       int
	PassingPipe bool
}

var Scoring = donburi.NewComponentType[ScoringData]()
