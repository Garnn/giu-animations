package animations

// PlayMode represents animation play mode.
type PlayMode byte

const (
	// PlayForward plays animation from 0 to 1 percentage progress.
	PlayForward PlayMode = iota
	// PlayBackwards plays an animation from 1 to 0 percentage progress.
	PlayBackwards
)
