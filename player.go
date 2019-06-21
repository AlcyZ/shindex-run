package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	PlayerHeight = 120
	PlayerWidth  = 99
)

func newPlayer(r *sdl.Renderer, path string) (*entity, error) {
	initPos := vector{x: 50, y: screenHeight - PlayerHeight - 100}
	player := newEntity(initPos)

	renderer, err := newRenderer(player, r, path, PlayerWidth, PlayerHeight)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create renderer: %v\n", err)
	}
	control := newAdControl(player)

	player.addComponent(control)

	// the render component should be the last attached, because its very likely that other components updates the
	// internal state
	player.addComponent(renderer)

	return player, nil
}
