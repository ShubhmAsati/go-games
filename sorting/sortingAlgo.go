package main

import (
	//"fmt"
	"time"
)

func bubbleSort(boxes boxesR) {

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if boxes[j].height > boxes[j+1].height {

				boxes[j], boxes[j+1] = boxes[j+1], boxes[j]
				boxes[j].y, boxes[j+1].y = boxes[j+1].y, boxes[j].y

				clear()
				printString("BUBBLE SORT",100,50,7);
				boxes.draw()
				tex.Update(nil, pixels, windowWidth*4)
				renderer.Copy(tex, nil, nil)
				renderer.Present()
				time.Sleep(100 * time.Millisecond)

			}
		}
	}
}

func quickSort(boxes boxesR, start, end int) {
	
	if start >= end {
		return
	}
	partitionIndex := partition(boxes, start, end)

	quickSort(boxes, start, partitionIndex)
	quickSort(boxes, partitionIndex+1, end)

}

func partition(boxes boxesR, start int, end int) int {

	element := boxes[end-1].height

	index := start

	for i := start; i < end; i++ {
		if boxes[i].height < element {
			boxes[i], boxes[index] = boxes[index], boxes[i]
			boxes[i].y, boxes[index].y = boxes[index].y, boxes[i].y
			clear()
			boxes.draw()
			printString("QUICK SORT",100,50,7);
			time.Sleep(200*time.Millisecond);
			tex.Update(nil, pixels, windowWidth*4)
			renderer.Copy(tex, nil, nil)
			renderer.Present()
			index++
		}
	}
	boxes[end-1], boxes[index] = boxes[index], boxes[end-1]
	boxes[end-1].y, boxes[index].y = boxes[index].y, boxes[end-1].y
	clear()
	boxes.draw()
	printString("QUICK SORT",100,50,7);
	time.Sleep(200*time.Millisecond);
	tex.Update(nil, pixels, windowWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	return index
}

func selectionSort(boxes boxesR) {
	for i := 0; i < len(boxes); i++ {
		smallest := 900
		index := i
		for j := i; j < len(boxes); j++ {
			if boxes[j].height < smallest {
				smallest = boxes[j].height
				index = j
			}
		}
		boxes[i], boxes[index] = boxes[index], boxes[i]
		boxes[i].y, boxes[index].y = boxes[index].y, boxes[i].y
		clear()
		boxes.draw()
		printString("SELECTION SORT",100,50,7);
		time.Sleep(200*time.Millisecond);
		tex.Update(nil, pixels, windowWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
	}
}

func mergeSort(boxes boxesR, start, end int) {
	
	if start >= end {
		return
	}
	mid := (start + end) / 2

	mergeSort(boxes, start, mid)
	mergeSort(boxes, mid+1, end)
	
	merge(boxes, start, mid, end)
	clear()
	boxes.draw()
	printString("MERGE SORT",100,50,7);
	time.Sleep(200*time.Millisecond);
	tex.Update(nil, pixels, windowWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()
}
func merge(boxes boxesR, start, mid, end int) {
	leftStart := start
	rightStart := mid + 1
	leftEnd := mid
	rightEnd := end
	index := start
	var la []*rectangle
	
	for i:= leftStart ; i < leftEnd+1;i++{
		
		r:= &rectangle{
			boxes[i].pos,
			boxes[i].width,
			boxes[i].height,
			boxes[i].color,
		};
		la = append(la,r);
	}

	
	var ra[]*rectangle

	for i:= rightStart ; i < rightEnd+1;i++{
		r:= &rectangle{
			boxes[i].pos,
			boxes[i].width,
			boxes[i].height,
			boxes[i].color,
		};
		ra = append(ra,r);
	}
	

	i := 0
	j := 0
	
	for i < len(la) && j < len(ra) {
		if la[i].height <= ra[j].height {
			
			y := boxes[index].y;
			boxes[index] = la[i]
			boxes[index].y = y;
			
			
			i++
		} else if ra[j].height <= la[i].height {
			y:= boxes[index].y;
			
			boxes[index] = ra[j]
			boxes[index].y = y;
		
			j++
		}
		index++

	}
	
	for i < len(la) {
		
		y:= boxes[index].y
		boxes[index] = la[i]
		boxes[index].y = y;
		i++
		index++

	}
	for j < len(ra) {
		y:= boxes[index].y
		boxes[index] = ra[j]
		boxes[index].y = y;
		j++;
		index++
	}
	
	clear()
	boxes.draw()
	printString("MERGE SORT",100,50,7);
	time.Sleep(200*time.Millisecond);
	tex.Update(nil, pixels, windowWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()	
}
