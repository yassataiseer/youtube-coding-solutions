package main

import (
	"math/rand"
	"time"
	"github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type Cars struct {
	posX int32
	posY int32
	velocity int32
	color string
  }

type Coins struct {
	posX int32
	posY int32
	velocity int32
}

func randomCar() Cars {
	rand.Seed(time.Now().UnixNano())
	car_colours:= [3]string{"blue","green","yellow"}
	car_locations:= [7]int32{1,83,167,250,333,417,501}

	var index int32 = int32(rand.Intn(3))
	var velocity int32
	var color string
	if(car_colours[index]=="green"){
		color = "green"
		velocity = 7
	}else if(car_colours[index]=="yellow"){
		color = "yellow"
		velocity = 5
	}else{
		color = "blue"
		velocity = 3
	}
	var Xpos int32 = car_locations[int32(rand.Intn(7))]
	current_car := Cars{Xpos,0,velocity,color}
	return current_car
}

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

	bluecar_img := rl.LoadImage("assets/bluecar.png")
	rl.ImageResize(bluecar_img,int32(82),int32(158))
	rl.ImageFlipVertical(bluecar_img)
	bluecar := rl.LoadTextureFromImage(bluecar_img)

	greencar_img := rl.LoadImage("assets/greencar.png")
	rl.ImageResize(greencar_img,int32(82),int32(158))
	rl.ImageFlipVertical(greencar_img)
	greencar := rl.LoadTextureFromImage(greencar_img)

	yellowcar_img := rl.LoadImage("assets/yellowcar.png")
	rl.ImageResize(yellowcar_img,int32(82),int32(158))
	rl.ImageFlipVertical(yellowcar_img)
	yellowcar := rl.LoadTextureFromImage(yellowcar_img)

	coin_img := rl.LoadImage("assets/coin.png")
	rl.ImageResize(coin_img,int32(82),int32(82))
	coin := rl.LoadTextureFromImage(coin_img)
	
	car_locations:= [7]int32{1,83,167,250,333,417,501}
	Coins := Coins{0,car_locations[int32(rand.Intn(7))],2}


	var x_coords int32 = 300
	var y_coords int32 = 450
	var roadY int32 = 0
	var score int = 0
	Cars := []Cars{}
	var times int = 3
	for times!=0{
		current_car := randomCar()
		Cars = append(Cars, current_car)
		times--
	}
	var isAlive bool= true

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(road,0,roadY,rl.White)

		rl.DrawTexture(redcar,x_coords,y_coords,rl.White)
		rl.ClearBackground(rl.RayWhite)

		for index,current_car := range Cars{
			if(current_car.color=="green"){
				rl.DrawTexture(greencar, current_car.posX, current_car.posY, rl.White)
			}else if(current_car.color=="yellow"){
				rl.DrawTexture(yellowcar, current_car.posX, current_car.posY, rl.White)
			}else {
				rl.DrawTexture(bluecar, current_car.posX, current_car.posY, rl.White)
			}
			Cars[index].posY += current_car.velocity
			if (current_car.posY>screenHeight){
				Cars[index] = randomCar()
			}
			if(rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords),float32(y_coords),float32(82),float32(158)),
			rl.NewRectangle(float32(current_car.posX),float32(current_car.posY),float32(82),float32(158)))){
				isAlive = false
			}
	
		}
		rl.DrawTexture(coin,Coins.posX,Coins.posY,rl.White)
		Coins.posY +=2
		if(Coins.posY>screenHeight){
			Coins.posY=0
			Coins.posX = car_locations[int32(rand.Intn(7))]
		}
		if(rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords),float32(y_coords),float32(82),float32(158)),
		rl.NewRectangle(float32(Coins.posX),float32(Coins.posY),float32(82),float32(82)))){
			Coins.posY=0
			Coins.posX = car_locations[int32(rand.Intn(7))]
			score++
		}

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

		if(!isAlive){
			Cars = nil
			x_coords=1000000
			rl.UnloadTexture(redcar)
			rl.UnloadTexture(greencar)
			rl.UnloadTexture(yellowcar)
			rl.UnloadTexture(bluecar)
			rl.UnloadTexture(road)
			rl.UnloadTexture(coin)
			//rl.ClearBackground(rl.White)
			rl.DrawText("Your final score is: "+strconv.Itoa(score),30,40,30,rl.White)

		}
		rl.DrawText("Score: "+strconv.Itoa(score), 0, 0, 20, rl.Black)
		rl.EndDrawing()
	}
	rl.CloseWindow()

}