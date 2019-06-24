package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const AnimationsRendererId = "animations_renderer"

type animationsRenderer struct {
	container *entity
	r         *sdl.Renderer
	a         *animations
}

func newAnimationsRenderer(container *entity, r *sdl.Renderer) (*animationsRenderer, error) {
	a, err := container.getComponent(AnimationsId)
	if err != nil {
		return &animationsRenderer{}, fmt.Errorf("could not get component:\n%v", err)
	}

	return &animationsRenderer{
		container: container,
		r:         r,
		a:         a.(*animations),
	}, nil
}

func (r *animationsRenderer) id() componentId {
	return AnimationsRendererId
}

func (r *animationsRenderer) update() error {
	t := r.a.current
	layout := r.a.layout(t)
	flip := r.a.flips[t]

	dest := &sdl.Rect{
		X: int32(r.container.position.x),
		Y: int32(r.container.position.y),
		W: layout.width,
		H: layout.height,
	}

	if err := r.r.CopyEx(layout.texture, nil, dest, 0, nil, flip); err != nil {
		return fmt.Errorf("could not render texture to window: \n%v", err)
	}

	return nil
}
