package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	//"strconv"
	"time"
    "math/rand"
	"fmt"
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

	rl.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - basic shapes drawing")

	rl.SetTargetFPS(60)
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

	current_pipe:=Pipe{screenWidth/2-texture.Width/2, int32(pipe_loc), 25, 25, rl.Red};
	for !rl.WindowShouldClose() {
		fmt.Println(y_coords)
		rl.BeginDrawing()

		if rl.IsKeyDown(rl.KeySpace){
			texture= rl.LoadTextureFromImage(bird_up)
			y_coords=y_coords-5
		}else{
			texture= rl.LoadTextureFromImage(bird_down)
			y_coords=y_coords+5
		}
		rl.DrawRectangle(current_pipe.posX,current_pipe.posY,current_pipe.width,current_pipe.height,current_pipe.Color)

		rl.DrawTexture(texture,x_coords, y_coords, rl.White)

		if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords),float32(34), float32(24)),rl.NewRectangle(float32(current_pipe.posX),float32(current_pipe.posY),float32(current_pipe.width),float32(current_pipe.height))){
			rl.DrawText("Game is over", 100, 120, 60, rl.LightGray)
		}
		rl.ClearBackground(rl.RayWhite)
		
		



		if(y_coords>=450){
			rl.DrawText("Game is over", 100, 120, 60, rl.LightGray)
			rl.CloseWindow()

		}
		rl.EndDrawing()
		time.Sleep(50000000)

	}

	rl.CloseWindow()
}