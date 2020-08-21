/*
The quiz from Gophercises by Jon Calhoun

A simple quiz programme to test elementary maths.
Programme reads a csv file and splits each line into the question and answer variables.
For each line display the question and get the answer through keyboard input.
Use string compare to check the result.
Keep a tally of the correct answers and present the score at the end.

June 2019
*/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var option string

	fmt.Println("The quiz from Gophercises - initial programme.")
	fmt.Println("A question will be displayed on the screen and you type in your answer.")
	fmt.Println("The questions are simple maths, add +, subtract -, multiple *, divide / and exponential **")
	fmt.Println("The programme will terminate when all of the questions have been posed.")
	fmt.Println("Place the cursor below this message and press enter to start")
	fmt.Scanln(&option)

	// Open the quiz file.
	file, err := os.Open("quizFile.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Read all of the lines at once into a slice.
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)

	// Parse the lines one at a time and split them into a question and an answer.
	// Get the input from the user and compare it to the answer.
	// Keep a track of the number of questions and correct answers.
	var questions, answers = 0, 0

	for _, line := range lines {
		question := line[0]
		answer := line[1]
		questions++
		fmt.Println("Question:", question, "?")
		fmt.Scanln(&option)
		if strings.Compare(option, "z") == 0 {
			break
		}
		if strings.Compare(answer, option) == 0 {
			answers++
		} else {
			fmt.Println("Incorrect answer")
		}
	}
	fmt.Println("You got", answers, "correct answers from", questions, "questions.")
}
