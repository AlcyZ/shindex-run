package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type fullscreenRenderer struct {
	container *entity
	renderer  *sdl.Renderer
	texture   *sdl.Texture
}

func newFullscreenRenderer(container *entity, renderer *sdl.Renderer, texture *sdl.Texture) *fullscreenRenderer {
	return &fullscreenRenderer{
		container: container,
		renderer:  renderer,
		texture:   texture,
	}
}

func (r *fullscreenRenderer) update() error {
	if err := r.renderer.Copy(r.texture, nil, nil); err != nil {
		return fmt.Errorf("could not render fullscreen texture: \n%v", err)
	}

	return nil
}

func (r *fullscreenRenderer) id() componentId {
	return "fullscreen_renderer"
}

type animationRenderer struct {
	container *entity
	renderer  *sdl.Renderer
	animation *animation
}

func newAnimationRenderer(container *entity, r *sdl.Renderer) (*animationRenderer, error) {
	anim, err := container.getComponent(AnimationId)
	if err != nil {
		return &animationRenderer{}, fmt.Errorf("could not create animation renderer: \n%v", err)
	}

	return &animationRenderer{
		container: container,
		renderer:  r,
		animation: anim.(*animation),
	}, nil
}

func (r *animationRenderer) id() componentId {
	return "fullscreen_renderer"
}

func (r *animationRenderer) update() error {
	layout := r.animation.layout()
	dest := &sdl.Rect{
		X: int32(r.container.position.x),
		Y: int32(r.container.position.y),
		W: layout.width,
		H: layout.height,
	}
	if err := r.renderer.Copy(layout.texture, nil, dest); err != nil {
		return fmt.Errorf("could not render fullscreen texture: \n%v", err)
	}
	return nil
}
