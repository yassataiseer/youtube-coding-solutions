package main

import (
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)


type Enemy struct{
	posX int32
	posY int32
	image_down bool
	draw bool
}

type Bullet struct{
	posX int32
	posY int32
	radius float32
	Draw bool
	Color rl.Color
}
func main(){
	screenWidth := int32(600)
	screenHeight := int32(900)
	rl.InitWindow(screenWidth,screenHeight , "Space Invaders")

	var x_coords int32 = 2
	var y_coords int32 = 800
	bullets := []Bullet{}
	var shouldShoot bool = true

	Ship_img := rl.LoadImage("assets/Ship.png")
	Enemy1 := rl.LoadImage("assets/enemy1.png")
	Enemy2 := rl.LoadImage("assets/enemy2.png")
	EnemyUp := rl.LoadTextureFromImage(Enemy1)
	EnemyDown := rl.LoadTextureFromImage(Enemy2)
	Enemies := []Enemy{}
	var enemy_int int = 5
	var current_x int32 = 10
	for enemy_int!=0{
		enemy_int--
		current_enemy := Enemy{current_x,100,true,true}
		current_x+=75
		Enemies = append(Enemies, current_enemy)
	}
	var Enemy_speed int32 = 1 
	Ship:=rl.LoadTextureFromImage(Ship_img)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
		rl.DrawTexture(Ship,x_coords,y_coords,rl.White)
		
		for index,current_enemy := range Enemies{
			if(Enemies[index].draw){
				if(current_enemy.image_down){
					Enemies[index].image_down=false
					rl.DrawTexture(EnemyUp,current_enemy.posX,current_enemy.posY,rl.White)
					
				}else{
					rl.DrawTexture(EnemyDown,current_enemy.posX,current_enemy.posY,rl.White)
					Enemies[index].image_down=true
				}
				if(current_enemy.posX<=0){
					Enemy_speed = 1
				}else if(current_enemy.posX>=screenWidth){
					Enemy_speed = -1
				}else{}

				Enemies[index].posX += Enemy_speed
			}
			time.Sleep(5000000)
		}

		for index1,current_bullet := range bullets {
			if(current_bullet.Draw){
				rl.DrawCircle(current_bullet.posX-16, current_bullet.posY, current_bullet.radius, current_bullet.Color)
				bullets[index1].posY = bullets[index1].posY-20
				if(current_bullet.posY<0){
					shouldShoot=true
					bullets[index1].Draw=false
					
				}
			}
		}

		if rl.IsKeyDown(rl.KeyD){
			if(x_coords+1>=screenWidth){
			}else{
				x_coords+=5
			}
		}
		if rl.IsKeyDown(rl.KeyA){
			if(x_coords-1<=0){
				}else{
					x_coords-=5
				}
		}
		if rl.IsKeyDown(rl.KeySpace)&&shouldShoot{

			current_bullet :=Bullet{ int32(x_coords+50),int32(y_coords+25),float32(10),true,rl.White}
			bullets = append(bullets,current_bullet )
			shouldShoot=false

		}

		
	}
	rl.CloseWindow()

}




