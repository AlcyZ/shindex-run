package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

type AnimationRenderer struct {
	container *engine.Entity
	renderer  *sdl.Renderer
	animation *Animation
}

func NewAnimationRenderer(container *engine.Entity, r *sdl.Renderer) (*AnimationRenderer, error) {
	anim, err := container.GetComponent(AnimationId)
	if err != nil {
		return &AnimationRenderer{}, fmt.Errorf("could not create Animation renderer: \n%v", err)
	}

	return &AnimationRenderer{
		container: container,
		renderer:  r,
		animation: anim.(*Animation),
	}, nil
}

func (r *AnimationRenderer) Id() engine.ComponentId {
	return "fullscreen_renderer"
}

func (r *AnimationRenderer) Update() error {
	layout := r.animation.Layout()
	position := r.container.CurrentPosition()
	dest := &sdl.Rect{
		X: int32(position.X),
		Y: int32(position.Y),
		W: layout.width,
		H: layout.height,
	}
	if err := r.renderer.CopyEx(layout.texture, nil, dest, 0, nil, layout.flip); err != nil {
		return fmt.Errorf("could not render fullscreen texture: \n%v", err)
	}
	return nil
}
