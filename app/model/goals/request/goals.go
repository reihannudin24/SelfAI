package domain

import (
	"time"
)

// AddGoals represents the data structure for adding a goal
type AddGoals struct {
	Goal       string    `json:"goal"`
	Type       string    `json:"type"`
	Time       time.Time `json:"time"`
	Date       time.Time `json:"date"`
	NewGroupId int       `json:"new_group_id"`
	UserId     int       `json:"user_id"`
}

// UpdateGoals represents the data structure for updating a goal
type UpdateGoals struct {
	ID         int       `json:"id"`
	Goal       string    `json:"goal"`
	Type       string    `json:"type"`
	Time       time.Time `json:"time"`
	Date       time.Time `json:"date"`
	NewGroupId int       `json:"new_group_id"`
	UserId     int       `json:"user_id"`
}

// DeleteGoals represents the data structure for deleting a goal
type DeleteGoals struct {
	ID int `json:"id"`
}

// Show represents the data structure for displaying a goal
type Show struct {
	ID         int       `json:"id"`
	Goal       string    `json:"goal"`
	Type       string    `json:"type"`
	Time       time.Time `json:"time"`
	Date       time.Time `json:"date"`
	NewGroupId int       `json:"new_group_id"`
	UserId     int       `json:"user_id"`
}
