package entities

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
	"shindex-run/engine/components"
	"time"
)

const (
	idle  = "idle"
	left  = "left"
	right = "right"
)

func NewPlayer(game *engine.Game, r *sdl.Renderer, speed float64, path string) (*engine.Entity, error) {
	initPos := engine.NewVector(50, engine.ScreenHeight-220)
	player := engine.NewEntity(game)
	player.ChangePosition(initPos)

	idleAnimation, err := getPlayerIdleAnimation(player, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create player idle animation: \n%v", err)
	}
	leftAnim, rightAnim, err := getPlayerRunAnimation(player, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create player run animation: \n%v", err)
	}
	attack, attackAnim, err := newAttack(player, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create player attack: \n%v", err)
	}

	animations := components.NewAnimations(player)
	animations.Add(idleAnimation, idle)
	animations.Add(leftAnim, left)
	animations.Add(rightAnim, right)
	animations.Add(attackAnim, "attack")
	player.AddComponent(animations)

	flips := components.NewFlips(player)
	flips.Add(sdl.FLIP_NONE, idle)
	flips.Add(sdl.FLIP_HORIZONTAL, left)
	flips.Add(sdl.FLIP_NONE, right)
	player.AddComponent(flips)

	control := components.NewHorizontalControl(player, speed, getLeftKeys(), getRightKeys())
	err = control.WithAnimations(idle, left, right) // animations must be attach to player component first
	err = control.WithFlips(idle, left, right)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not add animations to horizontal control: \n%v", err)
	}
	player.AddComponent(control)

	player.AddComponent(attack)

	// the render component should be the last attached, because its very likely that other components updates the
	// internal state to be rendered
	renderer, err := components.NewAnimationsRenderer(player, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create animations renderer: \n%v", err)
	}
	player.AddComponent(renderer)

	return player, nil
}

func getPlayerIdleAnimation(container *engine.Entity, r *sdl.Renderer) (*components.Animation, error) {
	var idleTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Idle__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &components.Animation{}, fmt.Errorf("could not load player idle texture %v, \n%v", i, err)
		}

		idleTxt = append(idleTxt, txt)
	}

	anim, err := components.NewAnimation(container, idleTxt, time.Second, 0.25)
	if err != nil {
		return &components.Animation{}, fmt.Errorf("could not create idle animation: \n%v", err)
	}

	return anim, nil
}

func getPlayerRunAnimation(container *engine.Entity, r *sdl.Renderer) (left *components.Animation, right *components.Animation, err error) {
	var runTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Run__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &components.Animation{}, &components.Animation{}, fmt.Errorf("could not load player run texture %v, \n%v", i, err)
		}

		runTxt = append(runTxt, txt)
	}

	left, err = components.NewAnimation(container, runTxt, time.Second, 0.25)
	if err != nil {
		return &components.Animation{}, &components.Animation{}, fmt.Errorf("could not create run animation: \n%v", err)
	}
	right, err = components.NewAnimation(container, runTxt, time.Second, 0.25)
	if err != nil {
		return &components.Animation{}, &components.Animation{}, fmt.Errorf("could not create run animation: \n%v", err)
	}

	return left, right, nil
}

func newAttack(container *engine.Entity, r *sdl.Renderer) (*components.Attack, *components.Animation, error) {
	var attackTextures []*sdl.Texture
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Attack__00%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return &components.Attack{}, &components.Animation{}, fmt.Errorf("could not load attack texture: %v", err)
		}
		attackTextures = append(attackTextures, texture)
	}

	animation, err := components.NewAnimation(container, attackTextures, time.Second/3, 0.25)
	if err != nil {
		return &components.Attack{}, &components.Animation{}, fmt.Errorf("could not create attack animation: %v", err)
	}

	return components.NewAttack(container, animation), animation, nil
}

func getLeftKeys() []sdl.Scancode {
	keys := make([]sdl.Scancode, 2)

	keys = append(keys, sdl.SCANCODE_A)
	keys = append(keys, sdl.SCANCODE_LEFT)

	return keys
}

func getRightKeys() []sdl.Scancode {
	keys := make([]sdl.Scancode, 2)

	keys = append(keys, sdl.SCANCODE_D)
	keys = append(keys, sdl.SCANCODE_RIGHT)

	return keys
}
