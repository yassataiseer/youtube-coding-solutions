package main

import "github.com/gen2brain/raylib-go/raylib"






func main() {
	screenWidth := int32(450)
    screenHeight := int32(800)

	rl.InitWindow(screenWidth, screenHeight, "DoodleJump")

	rl.SetTargetFPS(60)

	background_img := rl.LoadImage("assets/background.png")
	rl.ImageResize(background_img,screenWidth,screenHeight)
	background := rl.LoadTextureFromImage(background_img)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawTexture(background, 0, 0, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}