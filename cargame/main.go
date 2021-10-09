package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main(){
	screenWidth := int32(600)
	screenHeight := int32(900)
	rl.InitWindow(screenWidth,screenHeight , "Car Games")
	rl.SetTargetFPS(60)
	
	road_img := rl.LoadImage("assets/road.png")
	rl.ImageResize(road_img,int32(600),int32(900))
	road := rl.LoadTextureFromImage(road_img)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(road,0,0,rl.White)
		rl.ClearBackground(rl.RayWhite)

		rl.EndDrawing()
	}
	rl.CloseWindow()

}