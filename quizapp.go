package main

import(
	"fmt"
)
func ProbPull(fileName string)([]probs,error ){
//reads from quiz.csv
//open file
//create reader
//read it
//calll the parseprobs function
}

func main(){
	//input name of file
	//set the duration of timer
	//pull the problems from the files
	//error handling
	//var to count correct answers
	//loop for problems and printing them
	//result
}

func parseProb(lines [][]string)[]probs{
	//parse lines with problem struct
}

type probs struct{
	q string
	a string
}