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

	redcar_img := rl.LoadImage("assets/redcar.png")
	rl.ImageResize(redcar_img,int32(82),int32(158))
	redcar := rl.LoadTextureFromImage(redcar_img)

	var x_coords int32 = 300
	var y_coords int32 = 450


	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(road,0,0,rl.White)
		rl.DrawTexture(redcar,x_coords,y_coords,rl.White)
		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyDown(rl.KeyA)&&x_coords>-40{
			x_coords-=5
		}
		if rl.IsKeyDown(rl.KeyD)&&x_coords<525{
			x_coords+=5
		}
		if rl.IsKeyDown(rl.KeyW)&&y_coords>0{
			y_coords-=5
		}
		if rl.IsKeyDown(rl.KeyS)&&y_coords<800{
			y_coords+=5
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()

}