package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
	"github.com/gen2brain/raylib-go/raylib"
)

type Ground struct {
	posX int32
	posY int32
	width int32
	height int32
	Color rl.Color
  }
type Bullet struct{
	posX int32
	posY int32
	radius float32
	Right_direction bool
	Draw bool
	Color rl.Color
}

type Enemy struct{
	posX int32
	posY int32
	velocity int32
	Damage float32
	Direction bool
	Draw bool
	Color rl.Color
}


func RemoveIndex(bullets []Bullet, index int) []Bullet {
    ret := make([]Bullet, 0)
    ret = append(ret, bullets[:index]...)
    return append(ret, bullets[index+1:]...)
}


func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	rl.InitWindow(screenWidth,screenHeight , "FPS-Game")
	var Health int = 100
	var Score int = 0
	grounds := []Ground{}
	bullets := []Bullet{}


	ground := Ground{0,420,screenWidth,30,rl.Brown}
	table := Ground{screenWidth/2,screenHeight/2,screenWidth/4,30,rl.Brown}
	grounds = append(grounds,ground)
	grounds = append(grounds,table)
	Character := rl.LoadImage("assets/mario.png")
	LeftCharacter := rl.LoadImage("assets/leftmario.png")
	Enemies :=[]Enemy{}
	RightGoomba := rl.LoadImage("assets/leftgoomba.png")
	LeftGoomba := rl.LoadImage("assets/rightgoomba.png")
	first_enemy :=Enemy{0,370,5,0.5,true,true,rl.White}
	Enemies = append(Enemies, first_enemy)
	Enemy_texture:=rl.LoadTextureFromImage(RightGoomba)
	texture := rl.LoadTextureFromImage(Character)

    var x_coords int32 = screenWidth/2-texture.Width/2
    var y_coords int32 = screenHeight/2-texture.Height/2-40
	var should_shoot bool = true
	rl.SetTargetFPS(60)
	var isRight bool = true
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		fmt.Println(should_shoot)
		rl.ClearBackground(rl.RayWhite)
		if(Health>50){
			rl.DrawText("Health: "+strconv.Itoa(Health), 0, 0, 20, rl.Green)
		}else if(Health<50&&Health>10){
			rl.DrawText("Health: "+strconv.Itoa(Health), 0, 0, 20, rl.Orange)
		}else{
			rl.DrawText("Health: "+strconv.Itoa(Health), 0, 0, 20, rl.Red)
		}
		rl.DrawText("Score: "+strconv.Itoa(Score), 2, 300, 20, rl.Blue)
		rl.DrawTexture(texture,x_coords,y_coords,rl.White)
		
		for _,current_ground := range grounds{
			rl.DrawRectangle(current_ground.posX,current_ground.posY,current_ground.width,current_ground.height,current_ground.Color)
            if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords),float32(y_coords),float32(50),float32(50)),rl.NewRectangle(float32(current_ground.posX),float32(current_ground.posY),float32(current_ground.width),float32(current_ground.height))){
				y_coords=current_ground.posY-50
				if rl.IsKeyDown(rl.KeyW)||rl.IsKeyDown(rl.KeySpace){
					
					y_coords-=300
				}
				
				}
			
		}
		for index,current_enemy := range Enemies{
			if(Enemies[index].Draw){
				if(current_enemy.posX>screenWidth){
					Enemies[index].velocity = -5 
				}
				if(current_enemy.posX<0){
					Enemies[index].velocity =5 
				}			
				
				if(current_enemy.Direction){
					Enemy_texture = rl.LoadTextureFromImage(RightGoomba)
					Enemies[index].posX = int32(Enemies[index].posX +Enemies[index].velocity)
					Enemies[index].Direction = false
				}else{
					Enemy_texture = rl.LoadTextureFromImage(LeftGoomba)
					Enemies[index].Direction  = true
	
				}
				if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords),float32(y_coords),float32(50),float32(50)),rl.NewRectangle(float32(current_enemy.posX),float32(current_enemy.posY),float32(50),float32(50))){
					Health=Health-int(current_enemy.Damage)
					fmt.Println("detect")
				}
				rl.DrawTexture(Enemy_texture,current_enemy.posX,current_enemy.posY,current_enemy.Color)
				for index1,current_bullet := range bullets {
					if rl.CheckCollisionRecs(rl.NewRectangle(float32(current_bullet.posX),float32(current_bullet.posY),float32(current_bullet.radius),float32(current_bullet.radius)),rl.NewRectangle(float32(current_enemy.posX),float32(current_enemy.posY),float32(50),float32(50))){
						Enemies[index].Draw=false
						rand.Seed(time.Now().UnixNano())
						var side int = rand.Intn(2)
						fmt.Println("============")
						fmt.Println(side)
						if(side==1){
							new_enemy1 := Enemy{0,370,5,current_enemy.Damage*2,true,true,rl.White}
							Enemies = append(Enemies, new_enemy1)
						}else{
							new_enemy1 := Enemy{800,370,5,current_enemy.Damage*2,true,true,rl.White}
							Enemies = append(Enemies, new_enemy1)
						}
						Score++
						bullets[index1]=Bullet{}
						should_shoot=true
					}
	
					if(current_bullet.posX<0||current_bullet.posX>screenWidth){
						if(current_bullet.Draw){
							fmt.Println(current_bullet.posX)
							fmt.Println(current_bullet.posY)
							bullets[index1].Draw = false
							if(should_shoot){
								should_shoot = false
								
							}else{
								should_shoot=true
							}	
						}else{
							current_bullet.Draw=false
						}
		
					}else{
						if(current_bullet.Right_direction){
							bullets[index1].posX = int32(bullets[index1].posX+5)
							rl.DrawCircle(current_bullet.posX, current_bullet.posY, current_bullet.radius, current_bullet.Color)
							
						}else{
							bullets[index1].posX = int32(bullets[index1].posX-5)
							rl.DrawCircle(current_bullet.posX, current_bullet.posY, current_bullet.radius, current_bullet.Color)
							
						}
					}
		
		
				}
			}

		}





		if(y_coords+45<ground.posY){
			
			if(y_coords>ground.posY){

			}else{
				y_coords+=5
			}
		}
		if rl.IsKeyDown(rl.KeyD){
			texture = rl.LoadTextureFromImage(Character)
			x_coords+=5
			isRight=true
		}
		if rl.IsKeyDown(rl.KeyA){
			texture = rl.LoadTextureFromImage(LeftCharacter)
			x_coords-=5
			isRight=false
		}
		if rl.IsKeyDown(rl.KeyF){
			if(isRight){
				if(should_shoot){
					current_bullet :=Bullet{ int32(x_coords+50),int32(y_coords+25),float32(10),true,true,rl.Red}
					bullets = append(bullets,current_bullet )
					//-time.Sleep(500000)
					should_shoot = false
				}	
			}else{
				if(should_shoot){
				current_bullet := Bullet{ int32(x_coords-1),int32(y_coords+25),float32(10),false,true,rl.Red}
				bullets  = append(bullets,current_bullet )
				should_shoot = false
				}
			}
		}
		 

		if (y_coords<=0){
			y_coords+=15
		}
		if(x_coords<=0){
			x_coords+=15
		}
		if(x_coords>=screenWidth){
			x_coords-=15
		}
		if(Health<=0){
			Enemies=nil
			grounds=nil
			bullets=nil
			rl.UnloadTexture(texture)
            rl.DrawText("Your final score is: "+strconv.Itoa(Score),30,40,30,rl.Red)

		}
		rl.EndDrawing()

	}

	rl.CloseWindow()
}