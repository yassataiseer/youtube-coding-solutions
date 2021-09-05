package main

import (
	"math/rand"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

type Leaf struct {
	posX int32
	posY int32
	points int32
	velocity int32
	color string
  }

func randomLeaf() Leaf {
	rand.Seed(time.Now().UnixNano())
	leaf_colors:= [4]string{"red","green","purple","yellow"}
	leaf_locations:= [8]int32{1,70,160,210,280,360,500,580}
	
	var index int32 = int32(rand.Intn(3))
	var points int32 
	var velocity int32
	var color string
	if(leaf_colors[index]=="red"){
		points = 3
		color = "red"
		velocity = 3
	}else if(leaf_colors[index]=="green"){
		points = 7
		color = "green"
		velocity = 7
	}else if(leaf_colors[index]=="yellow"){
		points = 5
		color = "yellow"
		velocity = 5
	}else{
		points = -5
		color = "purple"
		velocity = 1
	}
	var Xpos int32 = leaf_locations[int32(rand.Intn(8))]
	current_leaf := Leaf{Xpos,0,points,velocity,color}
	return current_leaf
}

func main() {
	screenWidth := int32(600)
	screenHeight := int32(900)

	rl.InitWindow(screenWidth,screenHeight , "Space Invaders")
	var x_coords int32 = 2
	var y_coords int32 = 325
	var playerVel int32 = 5
	Leafs := []Leaf{}
	
	var times int = 9
	for times!=0{
		current_leaf := randomLeaf()
		Leafs = append(Leafs, current_leaf)
		times--
	}
	
	rl.SetTargetFPS(60)
	character_img := rl.LoadImage("assets/character.png")
	rl.ImageResize(character_img,int32(83),int32(100))

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

		rl.ClearBackground(rl.White)

		for index,current_leaf := range Leafs{

			if(current_leaf.color=="red"){
				rl.DrawTexture(redleaf, current_leaf.posX, current_leaf.posY, rl.White)
			}else if(current_leaf.color=="green"){
				rl.DrawTexture(greenleaf, current_leaf.posX, current_leaf.posY, rl.White)
			}else if(current_leaf.color=="yellow"){
				rl.DrawTexture(yellowleaf, current_leaf.posX, current_leaf.posY, rl.White)
			}else {
				rl.DrawTexture(purpleleaf, current_leaf.posX, current_leaf.posY, rl.White)
			}
			Leafs[index].posY += current_leaf.velocity
			if(current_leaf.posY>screenHeight){
				Leafs[index] = randomLeaf()
			}									
			if(rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords),float32(y_coords),float32(character.Width),float32(character.Height)),
			rl.NewRectangle(float32(current_leaf.posX),float32(current_leaf.posY),float32(50),float32(60)))){
				Leafs[index] = randomLeaf()

			}
			
		}


		if rl.IsKeyDown(rl.KeySpace)&&y_coords>-30{
			y_coords-=playerVel
		}else{
			y_coords+=5
		}
		if rl.IsKeyDown(rl.KeyA)&&x_coords>-40{
			x_coords-=5
		}
		if rl.IsKeyDown(rl.KeyD)&&x_coords<525{
			x_coords+=5
		}

		rl.DrawText("Score: ", 0, 0, 20, rl.Black)


		rl.EndDrawing()
	}

	rl.CloseWindow()
}


