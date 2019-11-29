package main

import (
	"fmt"
	"os"
	"bufio"
	
)

type choices struct{
	text string
	description string
	nextStory *storyNode
	nextChoices *choices
}

type storyNode struct{
	text string
	choices *choices
};

var scanner *bufio.Scanner;

func (s *storyNode) addChoice(choice string,description string,nextStory *storyNode){
	fmt.Println(description);
	if s.choices == nil{
		fmt.Println("inside if");
		s.choices = &choices{
			text: choice,
			description: description,
			nextStory: nextStory,
		};
	}else{
		currentChoices := s.choices;
		fmt.Println("inside else");
		for currentChoices.nextChoices != nil{
			currentChoices = currentChoices.nextChoices;
		}
		currentChoices.nextChoices = &choices{
			text: choice,
			description: description,
			nextStory: nextStory,
		}
	}
	
	
}



func (s *storyNode) showChoices(){
	
	currentChoices := s.choices;

	for currentChoices !=nil{
		fmt.Println(currentChoices.text + "   " + currentChoices.description);
		currentChoices = currentChoices.nextChoices;
	}
}


func (s *storyNode) play(){
	fmt.Println(s.text);
	if s.choices == nil{
		return
	}
	
	s.showChoices();
	for {
		scanner.Scan();
		response := scanner.Text()
		currentChoices := s.choices;
		for currentChoices !=nil{			
			if currentChoices.text == response{
				currentChoices.nextStory.play();
				return;
			}
			currentChoices = currentChoices.nextChoices;
		}
		fmt.Println("oops not an option lets try again");
	}

}

func main(){

	scanner = bufio.NewScanner(os.Stdin);
	standing := &storyNode{"you are standing in front of a door you have 2 options to go",nil};

	insideRoom := &storyNode{"now you are in a room you but this is a dark room choose options to lit up room",nil};

	youQuit := &storyNode{"you are dusgusting fuck you",nil};

	dieInDark := &storyNode{"you will die motherfucker",nil};
	
	litRoom := &storyNode{"now you can see three doors  and these will lead you to room a room full of girls or a room full of snakes or a room you dont want to go",nil}

	//threeDoor := &storyNode{"Pick a door to go ",nil}

	door1 := &storyNode{"Fooled you this is room of snakes biaaaatchhhhh ",nil};

	door2 := &storyNode{"Welcoe to the world full of weed and girls",nil}

	door3 := &storyNode{"your mother introducing you to some relatives",nil}

	standing.addChoice("A","open the door",insideRoom);
	standing.addChoice("B","go back to your home",youQuit);
	
	insideRoom.addChoice("A","turn on green lantern",litRoom);

	insideRoom.addChoice("B","you will die",dieInDark);

	litRoom.addChoice("A","this will open door for room full of girls",door1);
	litRoom.addChoice("B","room full of snakes",door2);
	litRoom.addChoice("C","this is a door that will show you wonders of nature",door3);


	standing.play();
		
}