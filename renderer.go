package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

type fullscreenRenderer struct {
	container *engine.Entity
	renderer  *sdl.Renderer
	texture   *sdl.Texture
}

func newFullscreenRenderer(container *engine.Entity, renderer *sdl.Renderer, texture *sdl.Texture) *fullscreenRenderer {
	return &fullscreenRenderer{
		container: container,
		renderer:  renderer,
		texture:   texture,
	}
}

func (r *fullscreenRenderer) Update() error {
	if err := r.renderer.Copy(r.texture, nil, nil); err != nil {
		return fmt.Errorf("could not render fullscreen texture: \n%v", err)
	}

	return nil
}

func (r *fullscreenRenderer) Id() engine.ComponentId {
	return "fullscreen_renderer"
}

type animationRenderer struct {
	container *engine.Entity
	renderer  *sdl.Renderer
	animation *animation
}

func newAnimationRenderer(container *engine.Entity, r *sdl.Renderer) (*animationRenderer, error) {
	anim, err := container.GetComponent(AnimationId)
	if err != nil {
		return &animationRenderer{}, fmt.Errorf("could not create animation renderer: \n%v", err)
	}

	return &animationRenderer{
		container: container,
		renderer:  r,
		animation: anim.(*animation),
	}, nil
}

func (r *animationRenderer) Id() engine.ComponentId {
	return "fullscreen_renderer"
}

func (r *animationRenderer) Update() error {
	layout := r.animation.layout()
	position := r.container.CurrentPosition()
	dest := &sdl.Rect{
		X: int32(position.X),
		Y: int32(position.Y),
		W: layout.width,
		H: layout.height,
	}
	if err := r.renderer.Copy(layout.texture, nil, dest); err != nil {
		return fmt.Errorf("could not render fullscreen texture: \n%v", err)
	}
	return nil
}
