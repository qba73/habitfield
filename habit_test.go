package habitfield_test

import (
	"bytes"
	habit "github.com/RyanRalphs/habitfield"
	"testing"
)

var scenarios = []struct {
	name  string
	input []string
	want  string
}{
	{
		name:  "Prints habit command",
		input: []string{"habit", "test"},
		want:  "test",
	},
	{
		name:  "Prints not a habit command",
		input: []string{"test"},
		want:  "test is not a habit command",
	},
	{
		name:  "Directs user to help if no habit provided",
		input: []string{"habit"},
		want:  "Habit is a command line tool for tracking habits. To get started, type 'habit help'",
	},
}

func TestProcessUserInput(t *testing.T) {
	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			fakeOutput := &bytes.Buffer{}
			input, err := habit.ProcessUserInput(test.input, fakeOutput)

			got := input
			if err != nil {
				got = err.Error()
			}

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestAddHabit(t *testing.T) {
	db, err := habit.SetUpDatabase()
	defer db.Close()
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	tracker := habit.NewHabitTracker(db)

	err = tracker.AddHabit("test")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}
}

func TestGetHabit(t *testing.T) {
	db, err := habit.SetUpDatabase()
	defer db.Close()
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	tracker := habit.NewHabitTracker(db)

	err = tracker.AddHabit("test")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	habit, err := tracker.GetHabit("test")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if habit.Name != "test" {
		t.Errorf("got %v, want %v", habit.Name, "test")
	}
}

func TestGetHabitNotFound(t *testing.T) {
	db, err := habit.SetUpDatabase()
	defer db.Close()
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	tracker := habit.NewHabitTracker(db)

	_, err = tracker.GetHabit("test2")
	if err == nil {
		t.Errorf("got %v, want %v", err, nil)
	}
}

func TestCheckForStreakAndUpdate(t *testing.T) {
	db, err := habit.SetUpDatabase()
	defer db.Close()
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	tracker := habit.NewHabitTracker(db)

	err = tracker.AddHabit("test")
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	testHabit, err := tracker.CheckForStreakAndUpdate("test")
	if err == nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if testHabit.Streak != 1 {
		t.Errorf("got %v, want %v", testHabit.Streak, 1)
	}

}
