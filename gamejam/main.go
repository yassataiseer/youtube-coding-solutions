package main

import "github.com/gen2brain/raylib-go/raylib"






func main() {
	screenWidth := int32(600)
	screenHeight := int32(900)

	rl.InitWindow(screenWidth,screenHeight , "Space Invaders")
	var x_coords int32 = 2
	var y_coords int32 = 325
	var playerVel int32 = 5

	rl.SetTargetFPS(60)
	character_img := rl.LoadImage("assets/character.png")
	character := rl.LoadTextureFromImage(character_img)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(character,x_coords,y_coords,rl.White)
		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyDown(rl.KeySpace){
			y_coords-=playerVel
		}else{
			y_coords+=5
		}
		if rl.IsKeyDown(rl.KeyA){
			x_coords-=5
		}
		if rl.IsKeyDown(rl.KeyD){
			x_coords+=5
		}



		rl.EndDrawing()
	}

	rl.CloseWindow()
}


