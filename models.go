package main

import (
	"time"
)


// Student Object
type Student struct {
	StudentID        int       `json:"StudentID"`
	Name     string    `json:"Name"`
	Course      string    `json:"Course"`
	Timestamp time.Time `json:"Timestamp"`
}

type Students []Student
