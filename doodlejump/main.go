package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)


type PlatForm struct{
	posX int32
	posY int32
	color rl.Color
}

func main(){
	screenWidth := int32(450)
	screenHeight := int32(800)

	rl.InitWindow(screenWidth,screenHeight,"Doodle Jump")

	rl.SetTargetFPS(60)

	background_img := rl.LoadImage("assets/background.png")
	rl.ImageResize(background_img,screenWidth,screenHeight)
	background := rl.LoadTextureFromImage(background_img)

	left_img := rl.LoadImage("assets/playerleft.png")
	left_char := rl.LoadTextureFromImage(left_img)


	right_img := rl.LoadImage("assets/playerright.png")
	right_char := rl.LoadTextureFromImage(right_img)	


	character := rl.LoadTextureFromImage(right_img)
	var x_pos int32 = 0
	var y_pos int32 = 0

	platforms := []PlatForm{}

	platform1 := PlatForm{10,100,rl.Black}
	platform2 := PlatForm{100,200,rl.Black}
	platform3 := PlatForm{300,300,rl.Brown}
	platform4 := PlatForm{40,400,rl.Black}
	platform5 := PlatForm{250,500,rl.Black}
	platform6 := PlatForm{450,600,rl.Brown}
	platform7 := PlatForm{450,700,rl.Black}
	platform8 := PlatForm{250,750,rl.Black}


	platforms = append(platforms, platform1)
	platforms = append(platforms, platform2)
	platforms = append(platforms, platform3)
	platforms = append(platforms, platform4)
	platforms = append(platforms, platform5)
	platforms = append(platforms, platform6)
	platforms = append(platforms, platform7)
	platforms = append(platforms, platform8)

	rand.Seed(time.Now().UnixNano())

	var score int = 0
	for !rl.WindowShouldClose(){
		rl.BeginDrawing()
		
		rl.DrawTexture(background,0,0,rl.White)

		rl.DrawTexture(character,x_pos,y_pos,rl.White)
		rl.DrawText("Score: "+strconv.Itoa(score),0,0,20,rl.Black)
		for index,current_platform := range platforms{
			rl.DrawRectangle(current_platform.posX,current_platform.posY,100,30,current_platform.color)
			
			if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_pos),float32(y_pos),float32(52),float32(120)),rl.NewRectangle(float32(current_platform.posX-37),float32(current_platform.posY),float32(100),float32(30))){
				y_pos-=120
				score++
				if(current_platform.color==rl.Brown){
					var PosX int32 = int32(rand.Intn(350))
					platforms[index].posX = PosX
					platforms[index].posY = 0
				}
				
			}
			platforms[index].posY += 1
			if(current_platform.posY>screenHeight){
				var PosX int32 = int32(rand.Intn(350))
				platforms[index].posX = PosX
				platforms[index].posY = 0
			}
		}

		if rl.IsKeyDown(rl.KeyA)&&x_pos<400&&x_pos>-30{
			x_pos-=5
			character = left_char
		}
		if rl.IsKeyDown(rl.KeyD)&&x_pos<400&&x_pos>-30{
			x_pos+=5
			character = right_char
		}
		if(x_pos>=400){
			x_pos-=5
		}
		if(x_pos<0){
			x_pos+=5
		}
		
		if(y_pos>screenHeight){
			platforms = nil
			rl.UnloadTexture(character)
			rl.ClearBackground(rl.White)
			rl.DrawText("Your final score is: "+strconv.Itoa(score),30,40,30,rl.Black)

		}

		rl.EndDrawing()
		y_pos+=2
	}
	rl.CloseWindow()

}