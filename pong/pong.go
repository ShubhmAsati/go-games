package main

import (
	"fmt"
	
	"github.com/veandco/go-sdl2/sdl"
)

const width, heigth int = 800, 600

type color struct {
	r, g, b, a byte
}

type pos struct {
	x, y float32
}

type ball struct {
	pos
	radius, xv, yv float32
	color
}

type paddle struct {
	pos
	w, h float32
	color
}

func (b *ball) draw(pixel []byte) {

	for i := -b.radius; i <= b.radius; i++ {
		for j := -b.radius; j <= b.radius; j++ {
			if i*i+j*j <= b.radius*b.radius {
				setPixel(int(i+b.x), int(j+b.y), b.color, pixel)
			}
		}
	}
}

func (p *paddle) brickCollision(b *ball) bool {

	j := p.y - p.w/2
	
	if b.x >= p.x-p.h/2 && b.x <= p.x + p.h/2 && int(b.y + b.radius) == int(j){
		b.yv = -b.yv;
		return true;
	}
	return false
}

func isBrickCollision(b *ball, br []*paddle, pixels []byte) {

	if b.y <= 400 {
		return
	}
	for index, val := range br {
		if val.brickCollision(b) {
			val.color = color{0, 0, 0, 0}
			val.draw(pixels)
			br = append(br[0:index], br[index+1:len(br)]...)
			return
		}
	}
}

func (b *ball) update(pixels []byte) {

	b.color = color{0, 0, 0, 0}
	b.draw(pixels)
	b.x += b.xv
	b.y += b.yv
	b.color = color{255, 255, 255, 0}
	if b.y < 42 {
		panic("err")
	}
	if b.x <= 0 || b.x >= float32(heigth) {
		b.xv = -b.xv
	}
	if b.y >= float32(width) {
		b.yv = -b.yv
	}
	b.draw(pixels)
}

func (p *paddle) update(keystate []uint8, pixels []byte) {
	if keystate[sdl.SCANCODE_UP] != 0 {
		if p.x-p.h/2 < 0 {
			return
		}
		p.x--
		p.clear(pixels, true)
	}
	if keystate[sdl.SCANCODE_DOWN] != 0 {
		if p.x+p.h/2 > float32(heigth) {
			return
		}
		p.x++
		p.clear(pixels, false)
	}
	p.draw(pixels)
}
func (p *paddle) clear(pixels []byte, isUp bool) {
	if isUp {
		i := p.x + 1 + p.h/2
		for j := 0; j <= int(p.w); j++ {
			setPixel(int(i), int(p.y-p.w/2)+j, color{0, 0, 0, 0}, pixels)
		}
	} else {
		i := p.x - 1 - p.h/2
		for j := 0; j <= int(p.w); j++ {
			setPixel(int(i), int(p.y-p.w/2)+j, color{0, 0, 0, 0}, pixels)
		}
	}
}

func (p *paddle) isCollision(b *ball) {
	j := p.y + p.w/2;
	if b.x >= p.x - p.h/2 && b.x <= p.x+p.h/2 {
		if  int(j) == int(b.y-b.radius) {
			b.yv = -b.yv
			return
		}
	}
}

func (p *paddle) draw(pixels []byte) {
	for i := p.x - p.h/2; i <= p.x+p.h/2; i++ {
		for j := p.y - p.w/2; j <= p.y+p.w/2; j++ {
			setPixel(int(i), int(j), p.color, pixels)
		}
	}
}

func setPixel(i, j int, c color, pixel []byte) {
	index := (i*width + j) * 4
	if index >= len(pixel)-4 || index < 0 {
		return
	}
	pixel[index], pixel[index+1], pixel[index+2] = c.r, c.g, c.b

}

func main() {
	window, err := sdl.CreateWindow("ping pong", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(heigth), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(width), int32(heigth))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()
	defer renderer.Destroy()
	defer tex.Destroy()

	pixels := make([]byte, width*heigth*4)
	b := &ball{pos{200, 200}, 10, 1, 1, color{255, 255, 255, 0}}
	b.draw(pixels)
	p := &paddle{pos{80, 50}, 15, 100, color{255, 255, 255, 0}}
	p.draw(pixels)
	var bricks []*paddle
	for i := 20; i <= heigth-20; i += 50 {
		for j := 500; j <= width-10; j += 20 {
			bricks = append(bricks, &paddle{pos{float32(i), float32(j)}, 15, 40, color{254, 254, 254, 0}})
		}
	}
	for _, val := range bricks {
		val.draw(pixels)
	}
	keyState := sdl.GetKeyboardState()


	//var frameStart time.Time

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		isBrickCollision(b, bricks, pixels)
		p.isCollision(b)
		b.update(pixels)
		p.update(keyState, pixels)
		tex.Update(nil, pixels, width*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(0)
	}
	
}
