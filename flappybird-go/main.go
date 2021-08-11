package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)




type Pipe struct {
	posX int32
	posY int32
	width int32
	height int32
	Color  rl.Color
	
}

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)
	rl.InitAudioDevice() 
	eat_noise := rl.LoadSound("sound/eat.wav") 
	rl.InitWindow(screenWidth, screenHeight, "FlappyApples")
	
	rl.SetTargetFPS(60)
	rl.SetSoundVolume(eat_noise, 0.2)
	bird_down := rl.LoadImage("assets/bird-down.png")
	bird_up := rl.LoadImage("assets/bird-up.png")	
	texture:= rl.LoadTextureFromImage(bird_up)
	var x_coords int32 = screenWidth/2-texture.Width/2
	var y_coords int32 = screenHeight/2-texture.Height/2-40
	//var pipes []Pipe
	rand.Seed(time.Now().UnixNano())
	var pipe_loc int = rand.Intn(450 -2+1)-2
	fmt.Println("======================")
	fmt.Println(pipe_loc)
	fmt.Println("======================")
	Pipes:= []Pipe{}
	current_pipes:=Pipe{800, int32(pipe_loc), 25, 25, rl.Red};
	Pipes = append(Pipes,current_pipes)
	var score int = 0
	for !rl.WindowShouldClose() {
		//fmt.Println(y_coords)
		rl.BeginDrawing()
		rl.DrawText("The current score is: "+strconv.Itoa(score), 0, 0, 30, rl.LightGray)
		if rl.IsKeyDown(rl.KeySpace){
			texture= rl.LoadTextureFromImage(bird_up)
			y_coords=y_coords-5
		}else{
			texture= rl.LoadTextureFromImage(bird_down)
			y_coords=y_coords+5
		}
		for io,current_pipe := range Pipes{

		fmt.Println(current_pipe.posX)
		fmt.Println(io)
		rl.DrawRectangle(current_pipe.posX,current_pipe.posY,current_pipe.width,current_pipe.height,current_pipe.Color)

		rl.DrawTexture(texture,x_coords, y_coords, rl.White)
		Pipes[io].posX = Pipes[io].posX-5
		if(current_pipe.posX <0){

			//Pipes[io] = Pipe{800, int32(pipe_loc), 25, 25, rl.Red}
			Pipes[io].posX = 800
			Pipes[io].posY = int32(rand.Intn(450 -2+1)-2)
			score--
			//Pipes = append(Pipes,current_pipes)
		}
		if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords),float32(34), float32(24)),rl.NewRectangle(float32(current_pipe.posX),float32(current_pipe.posY),float32(current_pipe.width),float32(current_pipe.height))){
			rl.PlaySoundMulti(eat_noise)
			Pipes[io].posX = 800
			Pipes[io].posY = int32(rand.Intn(450 -2+1)-2)
			score++
		}
	}
		rl.ClearBackground(rl.RayWhite)
		
		



		if(y_coords>=450){
			rl.UnloadTexture(texture)
			Pipes = nil
			rl.DrawText("Game is over"+" your score was: "+strconv.Itoa(score), 30, 40, 40, rl.Red)

		}
		rl.EndDrawing()
		time.Sleep(50000000)

	}
	rl.StopSoundMulti() // We must stop the buffer pool before unloading

	rl.UnloadSound(eat_noise) // Unload sound data

	rl.CloseAudioDevice()
	rl.CloseWindow()
}