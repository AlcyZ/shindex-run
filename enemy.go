package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	EnemyWidth  = 99
	EnemyHeight = 120
)

func newEnemy(r *sdl.Renderer, path string) (*entity, error) {
	initPos := vector{x: screenWidth - 50, y: screenHeight - EnemyHeight - 100}
	player := newEntity(initPos, EnemyWidth, EnemyHeight)

	renderer, err := newRenderer(player, r, path, EnemyWidth, EnemyHeight)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create renderer: %v\n", err)
	}

	player.addComponent(renderer)

	return player, nil
}
