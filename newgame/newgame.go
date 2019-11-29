package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	
)

const width,heigth int = 800,600;

type color struct{
	r,g,b,a byte
}

type pos struct{
	x,y float32
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

	for{
		
		tex.Update(nil,pixels,width*4);
		renderer.Copy(tex,nil,nil);
		renderer.Present()
	}
	sdl.Delay(2000);

}