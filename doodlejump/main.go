package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)






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
	
	var x_pos int32= 0
	var y_pos int32 = 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(background, 0, 0, rl.White)
		rl.DrawTexture(current_player, x_pos, y_pos, rl.White)
		if rl.IsKeyDown(rl.KeyA){
			x_pos-=5
			current_player = left_char
		}
		if rl.IsKeyDown(rl.KeyD){
			x_pos+=5
			current_player = right_char
		}

		
		
		rl.EndDrawing()
		y_pos+=2
	}

	rl.CloseWindow()
}