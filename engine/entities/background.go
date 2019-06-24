package entities

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
	"shindex-run/engine/components"
)

func NewBackground(game *engine.Game, r *sdl.Renderer, path string) (*engine.Entity, error) {
	t, err := img.LoadTexture(r, path)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not load background image %v: \n%v", path, err)
	}

	// position, width and height to 0, 0, because fullscreen is used anyway
	position := engine.NewVector(0, 0)
	bg := engine.NewEntity(game, position)

	renderer := components.NewFullscreenRenderer(bg, r, t)
	bg.AddComponent(renderer)

	return bg, nil
}
