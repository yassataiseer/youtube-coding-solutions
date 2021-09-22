package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type PlatForm struct {
	posX int32
	posY int32
	color rl.Color
  }

func main() {
	screenWidth := int32(450)
    screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "DoodleJump")

	rl.SetTargetFPS(60)

	background_img := rl.LoadImage("assets/background.png")
	rl.ImageResize(background_img,screenWidth,screenHeight)

	left_img := rl.LoadImage("assets/playerleft.png")
	left_char := rl.LoadTextureFromImage(left_img)

	right_img := rl.LoadImage("assets/playerright.png")
	right_char := rl.LoadTextureFromImage(right_img)

	background := rl.LoadTextureFromImage(background_img)
	current_player :=  rl.LoadTextureFromImage(right_img)
	platforms := []PlatForm{}

	platform1 := PlatForm{10,100,rl.Black}
	platform2 := PlatForm{100,200,rl.Black}
	platform3 := PlatForm{300,300,rl.Brown}
	platform4 := PlatForm{40,400,rl.Black}
	platform5 := PlatForm{250,500,rl.Black}
	
	platforms = append(platforms, platform1)
	platforms = append(platforms, platform2)
	platforms = append(platforms, platform3)
	platforms = append(platforms, platform4)
	platforms = append(platforms, platform5)

	var x_pos int32 = 0
	var y_pos int32 = 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(background, 0, 0, rl.White)
		rl.DrawTexture(current_player, x_pos, y_pos, rl.White)

		for _,current_platform := range platforms{
			rl.DrawRectangle(current_platform.posX,current_platform.posY,100,30,current_platform.color)

		}

		if (rl.IsKeyDown(rl.KeyA)&&x_pos<400&&x_pos>-30){
			x_pos-=5
			current_player = left_char
		}
		if (rl.IsKeyDown(rl.KeyD)&&x_pos<400&&x_pos>-30){
			x_pos+=5
			current_player = right_char
		}
		if(x_pos>=400){
			x_pos-=5
		}
		if(x_pos<0){
			x_pos+=5
		}
		
		rl.EndDrawing()
		y_pos+=2
	}

	rl.CloseWindow()
}