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

	bg := engine.NewEntity(game)

	renderer := components.NewFullscreenRenderer(bg, r, t)
	bg.AddComponent(renderer)

	return bg, nil
}
