package components

import (
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

type AnimationType = string

const AnimationsId engine.ComponentId = "Animations"

type Animations struct {
	container  *engine.Entity
	animations map[AnimationType]*Animation
	flips      map[AnimationType]sdl.RendererFlip
	current    AnimationType
}

func NewAnimations(container *engine.Entity) *Animations {
	a := make(map[AnimationType]*Animation)
	flips := make(map[AnimationType]sdl.RendererFlip)

	return &Animations{
		container:  container,
		animations: a,
		flips:      flips,
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
	if a.current != t {
		a.animations[a.current].currentIndex = 0
		a.current = t
	}
}

func (a *Animations) Id() engine.ComponentId {
	return AnimationsId
}

func (a *Animations) Update() error {
	_ = a.animations[a.current].Update()
	return nil
}
