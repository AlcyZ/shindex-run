package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

func newBackground(r *sdl.Renderer, path string) (*engine.Entity, error) {
	t, err := img.LoadTexture(r, path)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not load background image %v: \n%v", path, err)
	}

	// position, width and height to 0, 0, because fullscreen is used anyway
	position := engine.NewVector(0, 0)
	bg := engine.NewEntity(position)

	renderer := newFullscreenRenderer(bg, r, t)
	bg.AddComponent(renderer)

	return bg, nil
}
