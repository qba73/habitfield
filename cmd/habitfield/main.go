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
		fmt.Println(err)
	}

	tracker := habit.NewHabitTracker(db)

	err = tracker.AddHabit(input)
	if err != nil {
		fmt.Println(err)
	}
	record, err := tracker.GetHabit(input)

	if err != nil {
		fmt.Println(err)
	}

	record, err = tracker.CheckForStreakAndUpdate(record.Name)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(record)
}
