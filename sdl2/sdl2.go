package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"math"
)

const width,heigth int = 800,600;

type color struct{
	r,g,b,a byte
}

type pos struct{
	x,y float32
}

type ball struct{
	pos
	radius float32
	xv float32
	yv float32
	color 
}

type paddle struct{
	pos
	w float32
	h float32
	color 
}


func setPixel(i,j int,c color,pixel []byte){

	index := (i*width + j)*4

	if index >=len(pixel)-4 || index <0{
		return;
	}
	pixel[index] = c.r;
	pixel[index+1] = c.g;
	pixel[index+2] = c.b;
}


func main(){
	window,err := sdl.CreateWindow("Testing sdl2",sdl.WINDOWPOS_UNDEFINED,sdl.WINDOWPOS_UNDEFINED,int32(width),int32(heigth),sdl.WINDOW_SHOWN)

	if err != nil{
		fmt.Println(err)
		return
	}

	defer window.Destroy()

	renderer,err := sdl.CreateRenderer(window,-1,sdl.RENDERER_ACCELERATED);
	
	if err != nil{
		fmt.Println(err)
		return
	}

	defer renderer.Destroy();

	tex,err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888,sdl.TEXTUREACCESS_STREAMING,int32(width),int32(heigth));
	
	if err != nil{
		fmt.Println(err)
		return
	}

	defer tex.Destroy();

	pixels := make([]byte,width*heigth*4);

	radius :=100;
	for i:= 0 ;i < heigth ; i++{
		for j := 0 ; j<width ; j++{
			
			if int(math.Sqrt(math.Pow(float64(i - 200),2) + math.Pow(float64(j-200),2))) == radius{ 
				setPixel(i,j,color{255,255,255,0},pixels);	
			}else{
				setPixel(i,j,color{0,0,0,0},pixels);
			}
		}
	}
	

	tex.Update(nil,pixels,width*4);
	renderer.Copy(tex,nil,nil);
	renderer.Present()

	sdl.Delay(2000);

}