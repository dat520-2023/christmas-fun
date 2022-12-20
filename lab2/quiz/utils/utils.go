package utils

import (
	"bufio"
	"datxxx/lab2/quiz"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func ReadCommand(expected []string) int32 {
	return readInput(expected, "command")
}

func readInput(expected []string, operation string) int32 {
	for {
		fmt.Println()
		fmt.Printf("Select one of the %s: \n", operation)
		for num, command := range expected {
			fmt.Printf("%d) %s", num+1, command)
			fmt.Println()
		}
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Unable to read input")
			continue
		}
		selection, err := strconv.Atoi(string(input))
		if err != nil {
			fmt.Println("Please select the valid input")
			continue
		}
		if selection < 0 || selection > len(expected) {
			fmt.Println("Please select the valid command")
			continue
		}
		return int32(selection - 1)
	}
}

// ReadAnswerWithTimeout reads the participant response within a timeout
// It should return InvalidAnswer if timeout.
func ReadAnswerWithTimeout(t time.Duration, response chan int32) int32 {
	select {
	case <-time.After(t):
		log.Println("timedout")
		return quiz.InvalidAnswer
	case result := <-response:
		return result
	}
}

// ReadFromCommandLine reads the user input and return it in the resultChannel
func ReadFromCommandLine(optionLen int, resultChannel chan int32) {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Unable to read input")
			continue
		}
		selection, err := strconv.Atoi(string(input))
		if err != nil {
			fmt.Println("Please select the valid input")
			continue
		}
		if selection < 0 || selection > optionLen {
			fmt.Println("Please select the valid command")
			continue
		}
		resultChannel <- int32(selection - 1)
	}
}

func ReadEnter() bool {
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Unable to read input")
		return false
	}
	return true
}
