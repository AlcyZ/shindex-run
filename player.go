package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func newPlayer(r *sdl.Renderer, speed float64, path string) (*entity, error) {
	initPos := vector{x: 50, y: screenHeight - 220}
	player := newEntity(initPos)

	idleAnimation, err := getPlayerIdleAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create player idle animation: \n%v", err)
	}
	runAnimation, err := getPlayerRunAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create player run animation: \n%v", err)
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
		return &entity{}, fmt.Errorf("could not create movement animation renderer: \n%v", err)
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
			return &animation{}, fmt.Errorf("could not load player idle texture %v, \n%v", i, err)
		}

		idleTxt = append(idleTxt, txt)
	}

	anim, err := newAnimation(container, idleTxt, time.Second, 0.25)
	if err != nil {
		return &animation{}, fmt.Errorf("could not create idle animation: \n%v", err)
	}

	return anim, nil
}

func getPlayerRunAnimation(container *entity, r *sdl.Renderer) (*animation, error) {
	var runTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Run__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load player run texture %v, \n%v", i, err)
		}

		runTxt = append(runTxt, txt)
	}

	anim, err := newAnimation(container, runTxt, time.Second, 0.25)
	if err != nil {
		return &animation{}, fmt.Errorf("could not create run animation: \n%v", err)
	}

	return anim, nil
}
