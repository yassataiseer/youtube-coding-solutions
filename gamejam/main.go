package main

import "github.com/gen2brain/raylib-go/raylib"

type Leaf struct {
	posX int32
	posY int32
	width int32
	height int32
	Color rl.Color
  }




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
	
	redleaf_img := rl.LoadImage("assets/redleaf.png")
	greenleaf_img := rl.LoadImage("assets/greenleaf.png")
	purpleleaf_img := rl.LoadImage("assets/purpleleaf.png")
	yellowleaf_img := rl.LoadImage("assets/yellowleaf.png")

	rl.ImageResize(redleaf_img,int32(60),int32(60))
	rl.ImageResize(greenleaf_img,int32(60),int32(60))
	rl.ImageResize(purpleleaf_img,int32(60),int32(60))
	rl.ImageResize(yellowleaf_img,int32(60),int32(60))

	redleaf := rl.LoadTextureFromImage(redleaf_img)
	greenleaf := rl.LoadTextureFromImage(greenleaf_img)
	purpleleaf := rl.LoadTextureFromImage(purpleleaf_img)
	yellowleaf := rl.LoadTextureFromImage(yellowleaf_img)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(character,x_coords,y_coords,rl.White)
		rl.DrawTexture(redleaf, 0, 0, rl.White)
		rl.DrawTexture(greenleaf, 80, 0, rl.White)
		rl.DrawTexture(purpleleaf, 170, 0, rl.White)
		rl.DrawTexture(yellowleaf, 240, 0, rl.White)

		rl.ClearBackground(rl.White)

		if rl.IsKeyDown(rl.KeySpace)&&y_coords>-30{
			y_coords-=playerVel
		}else{
			y_coords+=5
		}
		if rl.IsKeyDown(rl.KeyA)&&x_coords>-40{
			x_coords-=5
		}
		if rl.IsKeyDown(rl.KeyD)&&x_coords<500{
			x_coords+=5
		}



		rl.EndDrawing()
	}

	rl.CloseWindow()
}


