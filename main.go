package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"shindex-run/engine"
	gameEntities "shindex-run/engine/entities"
	"time"
)

func main() {
	w, r, err := boot()
	if err != nil {
		fmt.Printf("could not boot:\n%v", err)
		return
	}
	defer sdl.Quit()
	defer w.Destroy()
	defer r.Destroy()

	soundStream, err := initAndPlaySound("sounds/Yung_Kartz_-_05_-_Loyalty.mp3")
	if err != nil {
		fmt.Println("sound init failed:\n", err)
		return
	}
	defer soundStream.Close()
	soundStream.Close()

	game := engine.CreateSimpleGame(w, r)

	player, err := gameEntities.NewPlayer(game, r, 16, "assets/ninja/Idle__000.png")
	if err != nil {
		fmt.Println("player init failed:\n", err)
		return
	}

	enemy, err := gameEntities.NewEnemy(game, r, "assets/player/male/Idle_0.png")
	if err != nil {
		fmt.Println("enemy init failed:\n", err)
		return
	}

	background, err := gameEntities.NewBackground(game, r, "assets/background.jpg")
	if err != nil {
		fmt.Println("bg init failed:\n", err)
		return
	}

	game.AddEntity(background)
	game.AddEntity(enemy)
	game.AddEntity(player)

	game.Start()
}

func boot() (*sdl.Window, *sdl.Renderer, error) {
	// init sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, fmt.Errorf("could not init SDL: \n%v", err)
	}

	// create window
	w, err := sdl.CreateWindow("Runner", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, engine.ScreenWidth, engine.ScreenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, fmt.Errorf("could not create window: \n%v", err)
	}

	// create renderer
	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, fmt.Errorf("could not create r: \n%v", err)
	}

	return w, r, nil
}

func initAndPlaySound(path string) (beep.StreamSeekCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open mp3 file: %v", err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode mp3: %v", err)
	}

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		return nil, fmt.Errorf("could not init speaker: %v", err)
	}

	speaker.Play(streamer)

	return streamer, nil
}
