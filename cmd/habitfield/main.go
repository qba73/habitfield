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

	db, err := habit.OpenDatabase("habits")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	tracker := habit.NewTracker(db)

	record, err := tracker.GetHabit(input)

	if err != nil {
		err = tracker.AddHabit(input)
		os.Exit(0)
	}

	if err != nil {
		fmt.Println(err)
	}

	record, err = tracker.UpdateHabit(record.Name)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", record)
}
