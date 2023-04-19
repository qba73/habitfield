package main

import (
	"fmt"
	habit "github.com/RyanRalphs/habitfield"
	"log"
	"os"
)

func main() {
	userInput := os.Args[1:]
	writer := os.Stdout
	if len(userInput) == 0 {
		habit.PrintHelp(writer)
		log.Fatal("Exiting Program. Please try again after reading the above help message!")
	}
	input, err := habit.ProcessUserInput(userInput, writer)

	if err != nil {
		log.Fatal(err)
	}

	db, err := habit.SetUpDatabase()
	defer db.Close()
	if err != nil {
		fmt.Println("Setup", err)
	}

	err = habit.AddHabit(db, input)
	if err != nil {
		fmt.Println("add ", err)
	}
	record, err := habit.GetHabit(db, input)

	if err != nil {
		fmt.Println("get", err)
	}

	fmt.Printf("%+v\n", record)

}
