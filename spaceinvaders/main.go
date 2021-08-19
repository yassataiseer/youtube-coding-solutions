package main
import (

	"github.com/gen2brain/raylib-go/raylib"
)


type Enemy struct{
	posX int32
	posY int32
	velocity int32
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
	Ship:=rl.LoadTextureFromImage(Ship_img)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.EndDrawing()
		rl.DrawTexture(Ship,x_coords,y_coords,rl.White)
		for index1,current_bullet := range bullets {
			if(current_bullet.Draw){
				rl.DrawCircle(current_bullet.posX-16, current_bullet.posY, current_bullet.radius, current_bullet.Color)
				bullets[index1].posY = bullets[index1].posY-1
				if(current_bullet.posY<0){
					shouldShoot=true
					bullets[index1].Draw=false
				}
			}
		}

		if rl.IsKeyDown(rl.KeyD){
			if(x_coords+1>=screenWidth){
			}else{
				x_coords+=1
			}
		}
		if rl.IsKeyDown(rl.KeyA){
			if(x_coords-1<=0){
				}else{
					x_coords-=1
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




