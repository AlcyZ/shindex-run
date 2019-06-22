package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type renderer struct {
	container     *entity
	texture       *sdl.Texture
	renderer      *sdl.Renderer
	width, height int32
}

func newRenderer(container *entity, r *sdl.Renderer, path string, width int32, height int32) (*renderer, error) {
	texture, err := img.LoadTexture(r, path)
	if err != nil {
		return &renderer{}, fmt.Errorf("could not create texture: %v\n", err)
	}

	return &renderer{
		container: container,
		renderer:  r,
		texture:   texture,
		width:     width,
		height:    height,
	}, nil
}

func (r *renderer) id() componentId {
	return "renderer"
}

func (r *renderer) update() error {
	dest := destFromEntity(r.container)
	_ = r.renderer.Copy(r.texture, nil, dest)
	return nil
}

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
		return fmt.Errorf("could not render fullscreen texture: %v", err)
	}

	return nil
}

func (r *fullscreenRenderer) id() componentId {
	return "fullscreen_renderer"
}

type animationRenderer struct {
	container     *entity
	renderer      *sdl.Renderer
	animation     *animation
	width, height int32
}

func newAnimationRenderer(container *entity, r *sdl.Renderer, animation *animation, width int32, height int32) *animationRenderer {
	return &animationRenderer{
		container: container,
		renderer:  r,
		animation: animation,
		width:     width,
		height:    height,
	}
}

func (r *animationRenderer) id() componentId {
	return "fullscreen_renderer"
}

func (r *animationRenderer) update() error {
	dest := destFromEntity(r.container)
	if err := r.renderer.Copy(r.animation.texture(), nil, dest); err != nil {
		return fmt.Errorf("could not render fullscreen texture: %v", err)
	}
	return nil
}

func destFromEntity(e *entity) *sdl.Rect {
	return &sdl.Rect{
		X: int32(e.position.x),
		Y: int32(e.position.y),
		W: e.width,
		H: e.height,
	}
}
