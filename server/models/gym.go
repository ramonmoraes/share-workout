package models

import "time"

type Exercise struct {
	Name       string
	Weight     int
	Sets       int
	Repetition int
	Obs        string
}

type Workout struct {
	Name        string
	Exercises   []Exercise
	LastUpdated time.Time
	Obs         string
}

type Routine struct {
	Workouts  []Workout
	CreatedAt time.Time
	obs       string
}
