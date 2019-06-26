package components

import (
	"shindex-run/engine"
)

type AnimationType = string

const AnimationsId engine.ComponentId = "Animations"

type Animations struct {
	container  *engine.Entity
	animations map[AnimationType]*Animation
	current    AnimationType
	locked     bool
}

func NewAnimations(container *engine.Entity) *Animations {
	a := make(map[AnimationType]*Animation)

	return &Animations{
		container:  container,
		animations: a,
	}
}

func (a *Animations) Add(animation *Animation, t AnimationType) {
	a.animations[t] = animation

	if a.current == "" {
		a.current = t
	}
}

func (a *Animations) Layout(t AnimationType) *engine.Layout {
	return a.animations[a.current].Layout()
}

func (a *Animations) ChangeAnimation(t AnimationType) {
	if !a.locked && a.current != t {
		a.animations[a.current].currentIndex = 0
		a.current = t
	}
}

func (a *Animations) SingleAnimation(t AnimationType) {
	a.ChangeAnimation(t)
	a.locked = true
}

func (a *Animations) Unlock() {
	a.locked = false
}

func (a *Animations) Id() engine.ComponentId {
	return AnimationsId
}

func (a *Animations) Update() error {
	_ = a.animations[a.current].Update()
	return nil
}
