package habitfield

import (
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"io"
	"time"
)

type Habit struct {
	ID                int    `storm:"id,increment"`
	Name              string `storm:"index,unique"`
	LastRecordedEntry time.Time
	Streak            int32 `storm:"index"`
}

func ProcessUserInput(userInput []string, writer io.Writer) (string, error) {
	if userInput[0] == "habit" {
		if len(userInput) > 1 {
			if userInput[1] == "help" {
				PrintHelp(writer)
				return "", fmt.Errorf("Hope this is helpful!")
			}
			for _, habit := range userInput[1:] {
				return habit, nil
			}
		}
		return "", fmt.Errorf("Habit is a command line tool for tracking habits. To get started, type 'habit help'")
	}
	return "", fmt.Errorf("%s is not a habit command", userInput[0])
}

func PrintHelp(writer io.Writer) {
	fmt.Fprintf(writer, "Welcome to your personal habit tracker!! \n \nTo add a habit, simply run `go run cmd/habitfield/main.go habit <habit>` \nIf you want to see how your habit streaks are going, simply run `go run cmd/habitfield/main.go habit` to see all stored habits and the current streak.\n \n")
}

func AddHabit(db *storm.DB, habitName string) error {
	_, err := GetHabit(db, habitName)
	if err != nil {

		todaysDate := time.Now().Format(time.RFC3339)
		lastRecordedEntry, err := time.Parse(time.RFC3339, todaysDate)
		if err != nil {
			return err
		}

		habit := Habit{
			Name:              habitName,
			LastRecordedEntry: lastRecordedEntry,
			Streak:            1,
		}

		err = db.Save(&habit)
		if err != nil {
			return err
		}

		fmt.Println("Habit Added!")
		return nil

	}
	return err
}

func GetHabit(db *storm.DB, habitName string) (Habit, error) {
	var habit Habit
	err := db.One("Name", habitName, &habit)
	if err != nil {
		fmt.Println(err)
		return habit, errors.New("habit not found")
	}
	return habit, err
}

func SetUpDatabase() (*storm.DB, error) {
	db, err := storm.Open("test.db")
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	return db, nil
}
