package main

import (

	"time"
	//"reflect"
	//"fmt"
	"github.com/veandco/go-sdl2/sdl"
)


var flag bool = true;

func setPixel(x, y int, c color, pixels []byte) {
	index := (x*windowWidth + y) * 4
	if index >= len(pixels)-4 || index < 0 {
		return
	}
	pixels[index] = c.r
	pixels[index+1] = c.g
	pixels[index+2] = c.b
}

func (b boxesR) draw() {
	for i := range b {
		startX := b[i].x - b[i].height/2
		endX := b[i].x + b[i].height/2
		startY := b[i].y - b[i].width/2
		endY := startY + b[i].width

		for j := startX; j <= endX; j++ {
			for k := startY; k <= endY; k++ {
				setPixel(j, k, b[i].color, pixels)
			}
		}

	}
}


func clear() {
	for i := 0; i < windowWidth*windowHeight*4; i++ {
		pixels[i] = 0
	}
}


func intializeGui() {
	window, _ = sdl.CreateWindow("Sorting algorithm", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(windowWidth), int32(windowHeight), sdl.WINDOW_SHOWN)

	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	tex, _ = renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(windowWidth), int32(windowHeight))

}

func initializeBoxes(boxes boxesR) boxesR{
	for i := 0; i < length; i++ {

		box := &rectangle{
			pos{baseLine - heigthArray[i]/2, 100 + boxWidth/2 + i*boxWidth + i*10},
			boxWidth,
			heigthArray[i],
			color{255, 255, 255, 0},
		}

		boxes = append(boxes, box)
	}
	boxes.draw()
	return boxes;
}

func getKeyBoardEvent(keyState []uint8,boxes boxesR){

	if flag{
		
		switch uint8(1){
		case keyState[sdl.SCANCODE_1]:
			flag = false;
			boxes = initializeBoxes(boxes)
			clear()
			boxes.draw()
			printString("BUBBLE SORT",100,50,7);
			tex.Update(nil, pixels, windowWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()
			time.Sleep(1*time.Second);
			bubbleSort(boxes);
			time.Sleep(3*time.Second);
			clear();
			flag = true;
		case keyState[sdl.SCANCODE_2]:
			flag = false;
			boxes = initializeBoxes(boxes);
			clear()
			boxes.draw()
			printString("QUICK SORT",100,50,7);
			tex.Update(nil, pixels, windowWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()
			time.Sleep(1*time.Second);
			quickSort(boxes,0,len(boxes));
			time.Sleep(3*time.Second);
			clear();
			flag = true;
		case keyState[sdl.SCANCODE_3]:
			flag = false;
			boxes = initializeBoxes(boxes);
			clear()
			boxes.draw()
			printString("SELECTION SORT",100,50,7);
			tex.Update(nil, pixels, windowWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()
			time.Sleep(1*time.Second);
			selectionSort(boxes);
			time.Sleep(3*time.Second);
			clear();
			flag = true;
		case keyState[sdl.SCANCODE_4]:
			flag = false;
			boxes = initializeBoxes(boxes);
			clear()
			boxes.draw()
			printString("MERGE SORT",100,50,7);
			tex.Update(nil, pixels, windowWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()
			time.Sleep(1*time.Second);
			mergeSort(boxes,0,len(boxes)-1);
			time.Sleep(3*time.Second);
			clear();
			flag = true;
		
		}
	}

	

	
}

func main() {

	intializeGui()

	defer window.Destroy()
	defer renderer.Destroy()
	defer tex.Destroy()

	keyState := sdl.GetKeyboardState()

	var boxes boxesR

	

	
	

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				
				getKeyBoardEvent(keyState,boxes);
			}
			if flag{

				printString("1 . BUBBLE SORT",200,50,7);
				printString("2 . QUICK SORT",300,50,7);
				printString("3 . SELECTION SORT",400,50,7);
				printString("4 . MERGE SORT",500,50,7);
			}
			printString("SORTING ALGORITHM DEPICTION FROM GOLANG",50,50,3);
				
			tex.Update(nil, pixels, windowWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()
			sdl.Delay(0)

		}
		
	}

}


