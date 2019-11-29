package main

import ()


type color struct {
	r, g, b, a byte
}

type pos struct {
	x, y int
}

type rectangle struct {
	pos
	width, height int
	color
}

type boxesR []*rectangle