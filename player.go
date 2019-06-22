package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	PlayerWidth  = 63
	PlayerHeight = 120
)

func newPlayer(r *sdl.Renderer, speed float64, path string) (*entity, error) {
	initPos := vector{x: 50, y: screenHeight - PlayerHeight - 100}
	player := newEntity(initPos, PlayerWidth, PlayerHeight)

	idleAnimation, err := getPlayerIdleAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create player idle animation: %v\n", err)
	}
	runAnimation, err := getPlayerRunAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create player run animation: %v\n", err)
	}
	movementAnimations := newMovementAnimations()
	movementAnimations.add(idleAnimation, sdl.FLIP_NONE, MovementIdle)
	movementAnimations.add(runAnimation, sdl.FLIP_HORIZONTAL, MovementLeft)
	movementAnimations.add(runAnimation, sdl.FLIP_NONE, MovementRight)

	adControl := newAdControl(player, speed)
	lrRightControl := newLeftRightControl(player, speed)

	player.addComponent(adControl)
	player.addComponent(lrRightControl)
	player.addComponent(movementAnimations)

	// the render component should be the last attached, because its very likely that other components updates the
	// internal state to be rendered
	renderer, err := newMovementAnimationRenderer(player, r, MovementIdle)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create movement animation renderer: %v\n", err)
	}
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

	return newAnimation(container, idleTxt, time.Second), nil
}

func getPlayerRunAnimation(container *entity, r *sdl.Renderer) (*animation, error) {
	var runTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Run__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load player run texture %v, %v", i, err)
		}

		runTxt = append(runTxt, txt)
	}

	return newAnimation(container, runTxt, time.Second), nil
}
