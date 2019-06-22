package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func newBackground(r *sdl.Renderer, path string) (*entity, error) {
	t, err := img.LoadTexture(r, path)
	if err != nil {
		return &entity{}, fmt.Errorf("could not load background image %v: %v", path, err)
	}

	// position, width and height to 0, 0, because fullscreen is used anyway
	position := vector{x: 0, y: 0}
	bg := newEntity(position, 0, 0)

	renderer := newFullscreenRenderer(bg, r, t)
	bg.addComponent(renderer)

	return bg, nil
}
