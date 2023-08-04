package main

import(
		"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)
)
func ProbPull(fileName string)([]probs,error ){
//reads from quiz.csv
		if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)
		if cLines, err := csvR.ReadAll(); err == nil {
			return parseProb(cLines), nil
		}
	}
//open file
//create reader
//read it
//calll the parseprobs function
}

func main(){
	//input name of file
	fName := flag.String("f", "quiz.csv", "path of csv file")
	//set the duration of timer
	timer := flag.Int("t", 20, "timer for the quiz")
	//pull the problems from the files
	flag.Parse()
	probs, err := ProbPull()

	if err != nil {
		exit(fmt.Sprintf("something went wrong:%s", err.Error()))
	}
	correct := 0

	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
}

func parseProb(lines [][]string)[]probs{
	//parse lines with problem struct
}

type probs struct{
	q string
	a string
}
