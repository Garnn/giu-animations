package animations

// Animation is an interface implemented by each animation.
type Animation interface {
	// Init is called once, immediately on start.
	Init()
	// Reset is called whenever needs to restart animation.
	Reset()

	// BuildNormal is called every frame when animation is not running
	// starter is a link to Animator.Start
	BuildNormal(starter func())
	// BuildAnimation is called when running an animation.
	// It receives the current animation progress as a float, where
	// 0 >= animationPercentage <= 1
	// starter is a link to Animator.Start
	BuildAnimation(animationPercentage float32, starter func())
}
