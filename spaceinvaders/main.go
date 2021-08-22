package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
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
	velocity int32
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
	var score int = 0
	var enemies_Alive int = 7
	Ship_img := rl.LoadImage("assets/Ship.png")
	Enemy1 := rl.LoadImage("assets/enemy1.png")
	Enemy2 := rl.LoadImage("assets/enemy2.png")
	EnemyUp := rl.LoadTextureFromImage(Enemy1)
	EnemyDown := rl.LoadTextureFromImage(Enemy2)
	Enemies := []Enemy{}
	var GameOver bool = false
	var enemy_int int = 7
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
		rl.DrawText("Score: "+strconv.Itoa(score), 0, 0, 20, rl.LightGray)
		fmt.Println(enemies_Alive)
		rand.Seed(time.Now().UnixNano())
		rl.DrawTexture(Ship,x_coords,y_coords,rl.White)
		if(GameOver||enemies_Alive==0){
			Enemies = nil
			bullets = nil
			rl.UnloadTexture(Ship)
			x_coords = 1000000000

			rl.DrawText("Your final score is: "+strconv.Itoa(score),30,40,30,rl.White)
			rl.DrawText("Press Enter to restart ",30,80,30,rl.White)
			if rl.IsKeyDown(rl.KeyEnter){
				rl.CloseWindow()
				main()
			}else{}
		
		}
		for index,current_enemy := range Enemies{
			
			if(Enemies[index].draw){
				var enemy_shoot int32 = int32(rand.Intn(200))
				if(enemy_shoot==10){
					current_bullet :=Bullet{ int32(current_enemy.posX+24),int32(current_enemy.posY),-5,float32(10),true,rl.Red}
					bullets = append(bullets,current_bullet )	
				}
				if(current_enemy.image_down){
					Enemies[index].image_down=false
					rl.DrawTexture(EnemyUp,current_enemy.posX,current_enemy.posY,rl.White)
					
				}else{
					rl.DrawTexture(EnemyDown,current_enemy.posX,current_enemy.posY,rl.White)
					Enemies[index].image_down=true
				}
				if(current_enemy.posX==0||current_enemy.posX==screenWidth){
					for i,_ := range Enemies{
						Enemies[i].posY+=5
					}
				}

				if(current_enemy.posX<=0){
					Enemy_speed = 1
				}else if(current_enemy.posX>=screenWidth){
					Enemy_speed = -1
				}else{}
				
				Enemies[index].posX += Enemy_speed
			}
			time.Sleep(5000000)

			
			for index1,current_bullet := range bullets {
				if(current_bullet.Draw){
					bullets[index1].posY = bullets[index1].posY - current_bullet.velocity
					rl.DrawCircle(current_bullet.posX-16, current_bullet.posY, current_bullet.radius, current_bullet.Color)
					if(current_bullet.posY<0||current_bullet.posY>screenHeight){
						shouldShoot=true
						bullets[index1].Draw=false
					}
					if rl.CheckCollisionRecs(rl.NewRectangle(float32(current_bullet.posX),float32(current_bullet.posY),float32(current_bullet.radius),float32(current_bullet.radius)),
					rl.NewRectangle(float32(current_enemy.posX),float32(current_enemy.posY),float32(60),float32(32)))&&current_bullet.Color==rl.White&&current_enemy.draw{
						Enemies[index].draw = false
						score+=50
						enemies_Alive--
					}
					if rl.CheckCollisionRecs(rl.NewRectangle(float32(current_bullet.posX),float32(current_bullet.posY),float32(current_bullet.radius),float32(current_bullet.radius)),
					rl.NewRectangle(float32(x_coords),float32(y_coords),float32(60),float32(32)))&&current_bullet.Color==rl.Red{
						fmt.Println("lost!")
						GameOver = true
						break
					}
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

			current_bullet :=Bullet{ int32(x_coords+50),int32(y_coords+25),5,float32(10),true,rl.White}
			bullets = append(bullets,current_bullet )
			shouldShoot=false

		}

		rl.EndDrawing()
	}
	rl.CloseWindow()

}




