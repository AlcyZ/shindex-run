package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	EnemyHeight = 120
	EnemyWidth  = 99
)

func newEnemy(r *sdl.Renderer, path string) (*entity, error) {
	initPos := vector{x: screenWidth - 50, y: screenHeight - PlayerHeight - 100}
	player := newEntity(initPos)

	renderer, err := newRenderer(player, r, path, EnemyWidth, EnemyHeight)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create renderer: %v\n", err)
	}

	player.addComponent(renderer)

	return player, nil
}
