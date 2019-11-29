package main

import (
	
)

func printLetter( lt byte,x,y,size int){
	
	for i:= 0 ; i < 5 ; i++{ 
		for j :=0 ; j < 5 ;j++{
			for k:=0 ; k <=size ; k++{
				for l:= 0 ; l <=size;l++{
					setPixel(x+k + size*i,y+l+size*j,color{255*byte(letter[int(lt)-32][i][j]),255*byte(letter[int(lt)-32][i][j]),255*byte(letter[int(lt)-32][i][j]),0},pixels);
				}
			}
		}
	}

}


func printString(sentence string,x,y,size int){

	for i:= range sentence{
		printLetter(sentence[i],x,y+size*i*5+i*5,size);
	}
	
}