package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	PlayerHeight = 120
	PlayerWidth  = 63
)

func newPlayer(r *sdl.Renderer, speed float64, path string) (*entity, error) {
	initPos := vector{x: 50, y: screenHeight - PlayerHeight - 100}
	player := newEntity(initPos)

	animation, err := getPlayerIdleAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not player idle animation: %v\n", err)
	}

	adControl := newAdControl(player, speed)
	lrRightControl := newLeftRightControl(player, speed)

	player.addComponent(adControl)
	player.addComponent(lrRightControl)
	player.addComponent(animation)

	// the render component should be the last attached, because its very likely that other components updates the
	// internal state
	renderer := newAnimationRenderer(player, r, animation, PlayerWidth, PlayerHeight)
	player.addComponent(renderer)

	return player, nil
}

func getPlayerIdleAnimation(container *entity, r *sdl.Renderer) (*animation, error) {
	var idleTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Idle__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load player idle texture %v, %v", i, err)
		}

		idleTxt = append(idleTxt, txt)
	}

	return newAnimation(container, idleTxt, time.Millisecond*600), nil
}
