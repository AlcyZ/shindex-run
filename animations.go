package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

type animationType = string

const AnimationsId engine.ComponentId = "animations"

type animations struct {
	container  *engine.Entity
	animations map[animationType]*animation
	flips      map[animationType]sdl.RendererFlip
	current    animationType
}

func newAnimations(container *engine.Entity) *animations {
	a := make(map[animationType]*animation)
	flips := make(map[animationType]sdl.RendererFlip)

	return &animations{
		container:  container,
		animations: a,
		flips:      flips,
	}
}

func (a *animations) add(animation *animation, flip sdl.RendererFlip, t animationType) {
	a.animations[t] = animation
	a.flips[t] = flip

	if a.current == "" {
		a.current = t
	}
}

func (a *animations) layout(t animationType) *layout {
	return a.animations[t].layout()
}

func (a *animations) changeAnimation(t animationType) {
	if a.current != t {
		a.animations[a.current].currentIndex = 0
		a.current = t
	}
}

func (a *animations) Id() engine.ComponentId {
	return AnimationsId
}

func (a *animations) Update() error {
	_ = a.animations[a.current].Update()
	return nil
}
