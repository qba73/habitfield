package habitfield

import (
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"io"
	"strings"
	"time"
)

type HabitTracker struct {
	db *storm.DB
}

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
			habit := strings.Join(userInput[1:], " ")
			return habit, nil

		}
		return "", fmt.Errorf("Habit is a command line tool for tracking habits. To get started, type 'habit help'")
	}
	return "", fmt.Errorf("%s is not a habit command", userInput[0])
}

func NewHabitTracker(db *storm.DB) *HabitTracker {
	return &HabitTracker{db: db}
}

func (ht *HabitTracker) AddHabit(habitName string) error {
	_, err := ht.GetHabit(habitName)
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

		err = ht.db.Save(&habit)
		if err != nil {
			return err
		}

		fmt.Println("Habit Added!")
		return nil

	}
	return err
}

func (ht *HabitTracker) GetHabit(habitName string) (Habit, error) {
	var habit Habit
	err := ht.db.One("Name", habitName, &habit)
	if err != nil {
		return habit, err
	}
	return habit, nil
}

func (ht *HabitTracker) CheckForStreakAndUpdate(habitName string) (Habit, error) {
	var habit Habit
	err := ht.db.One("Name", habitName, &habit)
	if err != nil {
		fmt.Println(err)
		return habit, errors.New("habit not found")
	}

	todaysDate := time.Now().Format(time.RFC3339)
	lastRecordedEntry, err := time.Parse(time.RFC3339, todaysDate)
	if err != nil {
		return habit, err
	}

	if lastRecordedEntry.Day() == habit.LastRecordedEntry.Day() {
		return habit, errors.New("habit already recorded for today")
	}

	habit.LastRecordedEntry = lastRecordedEntry
	habit.Streak++

	err = ht.db.Save(&habit)
	if err != nil {
		return habit, err
	}

	fmt.Println("Habit Updated!")
	return habit, nil
}

func PrintHelp(writer io.Writer) {
	fmt.Fprintf(writer, "Welcome to your personal habit tracker!! \n \nTo add a habit, simply run `go run cmd/habitfield/main.go habit <habit>` \nIf you want to see how your habit streaks are going, simply run `go run cmd/habitfield/main.go habit` to see all stored habits and the current streak.\n \n")
}

func SetUpDatabase() (*storm.DB, error) {
	db, err := storm.Open("test.db")
	if err != nil {
		return db, err
	}

	return db, nil
}
