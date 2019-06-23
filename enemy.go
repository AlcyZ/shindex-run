package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func newEnemy(r *sdl.Renderer, path string) (*entity, error) {
	initPos := vector{x: screenWidth - 150, y: screenHeight - 240}
	player := newEntity(initPos)

	animation, err := getEnemyIdleAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create enemy idle animation: \n%v", err)
	}
	player.addComponent(animation)

	renderer, err := newAnimationRenderer(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create animation renderer: \n%v", err)
	}
	player.addComponent(renderer)

	return player, nil
}

func getEnemyIdleAnimation(container *entity, r *sdl.Renderer) (*animation, error) {
	var textures []*sdl.Texture

	for i := 0; i < 15; i++ {
		path := fmt.Sprintf("assets/player/male/Idle_%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load enemy idle texture: \n%v", err)
		}
		textures = append(textures, texture)
	}

	anim, err := newAnimation(container, textures, time.Second, 0.25)
	if err != nil {
		return &animation{}, fmt.Errorf("could not create enemy idle animation: \n%v", err)
	}

	return anim, nil
}
