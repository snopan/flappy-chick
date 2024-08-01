package options

const (
	WindowWidth                  = 800
	WindowHeight                 = 1000
	Graivty                      = 9.8 * 0.02
	FlyUpSpeed                   = -MaxTiltVelocity
	PlayerScale                  = 2.5
	GroundScale                  = 2.0
	GroundSpeed                  = -3.0
	PipeScale                    = 4.0
	PipeSpeed                    = -3.0
	PipeCenterWindowHeightFactor = 0.3
	PipeGap                      = 150.0
	PipesDistance                = 200.0
	MaxTiltVelocity              = 8.0
	MaxTiltAngle                 = 1.0
)

const (
	LayerBackground int = iota
	LayerPipe
	LayerGround
	LayerPlayer
	LayerUI
)
