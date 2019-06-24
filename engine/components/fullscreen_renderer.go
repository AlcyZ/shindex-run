package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

type FullscreenRenderer struct {
	container *engine.Entity
	renderer  *sdl.Renderer
	texture   *sdl.Texture
}

func NewFullscreenRenderer(container *engine.Entity, renderer *sdl.Renderer, texture *sdl.Texture) *FullscreenRenderer {
	return &FullscreenRenderer{
		container: container,
		renderer:  renderer,
		texture:   texture,
	}
}

func (r *FullscreenRenderer) Update() error {
	if err := r.renderer.Copy(r.texture, nil, nil); err != nil {
		return fmt.Errorf("could not render fullscreen texture: \n%v", err)
	}

	return nil
}

func (r *FullscreenRenderer) Id() engine.ComponentId {
	return "fullscreen_renderer"
}
