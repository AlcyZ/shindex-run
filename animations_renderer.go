package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

const AnimationsRendererId = "animations_renderer"

type animationsRenderer struct {
	container *engine.Entity
	r         *sdl.Renderer
	a         *animations
}

func newAnimationsRenderer(container *engine.Entity, r *sdl.Renderer) (*animationsRenderer, error) {
	a, err := container.GetComponent(AnimationsId)
	if err != nil {
		return &animationsRenderer{}, fmt.Errorf("could not get component:\n%v", err)
	}

	return &animationsRenderer{
		container: container,
		r:         r,
		a:         a.(*animations),
	}, nil
}

func (r *animationsRenderer) Id() engine.ComponentId {
	return AnimationsRendererId
}

func (r *animationsRenderer) Update() error {
	t := r.a.current
	position := r.container.CurrentPosition()
	layout := r.a.layout(t)
	flip := r.a.flips[t]

	dest := &sdl.Rect{
		X: int32(position.X),
		Y: int32(position.Y),
		W: layout.width,
		H: layout.height,
	}

	if err := r.r.CopyEx(layout.texture, nil, dest, 0, nil, flip); err != nil {
		return fmt.Errorf("could not render texture to window: \n%v", err)
	}

	return nil
}
