package main

import (
	"fmt"
	//"time"
)

func quickSort(arr []int, start, end int) {
		
	if start >= end {
		return;
	}
		partitionIndex := partition(arr, start, end)

		quickSort(arr, start, partitionIndex)
		quickSort(arr, partitionIndex+1, end)
	

}

func partition(arr []int, start int, end int) int {

	element := arr[end-1]

	index := start

	for i := start; i < end; i++ {
		if arr[i] < element {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[end-1], arr[index] = arr[index], arr[end-1]

	return index;
}

func selectionSort(arr []int){

	for i:=0 ; i < len(arr);i++{
		smallest := 900;
		index := i;
		for j:= i ; j < len(arr); j++{
			if arr[j] < smallest{
				smallest = arr[j];
				index = j;
			}
		}
		arr[i],arr[index]  = arr[index],arr[i];
	}


}

func mergeSort(arr []int,start,end int){
	
	if start >= end {
		return
	}
	mid := (start + end)/2

	mergeSort(arr,start,mid);
	mergeSort(arr,mid+1,end);
	merge(arr,start,mid,end);
}
func merge(arr []int,start,mid,end int){
	leftStart := start;
	rightStart := mid+1;
	leftEnd := mid;
	rightEnd := end;
	index := start;
	var la = make([]int,leftEnd-leftStart +1); 
	copy(la,arr[leftStart:leftEnd+1]);
	var ra = make([]int,rightEnd-rightStart+1) 
	copy(ra,arr[rightStart:rightEnd+1]);
	
	i:=0;
	j:=0
	for i < len(la) && j < len(ra){
		if la[i] <= ra[j]{
			arr[index] = la[i];
			i++;
		}else if ra[j] <= la[i]{
			arr[index] = ra[j]
			j++;
		}
		index++;
		
	}
	
	for i < len(la){
	
		arr[index] = la[i]
		i++;
		index++;
		
	}
	for j < len(ra){
		
		arr[index] = ra[j]
		j++;
		index++;
	}

	
	
}


func main() {
	arr := []int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 23,
	}

	//quickSort(arr, 0, len(arr))
	//selectionSort(arr);
	
	mergeSort(arr,0,len(arr)-1);
	
	fmt.Println(arr)

}
