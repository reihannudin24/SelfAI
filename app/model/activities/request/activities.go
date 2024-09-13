package request

import "time"

type CreateActivities struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Address   string    `json:"address"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
	Remainder string    `json:"remainder"`
	Type      string    `json:"type"`
	UserId    int       `json:"user_id"`
	Token     string    `json:"token"`
}

type UpdateActivities struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Content   string    `json:"content"`
	Address   string    `json:"address"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Date      time.Time `json:"date"`
	Remainder string    `json:"remainder"`
	Type      string    `json:"type"`
	UserId    int       `json:"user_id"`
	Token     string    `json:"token"`
}

type DeleteActivities struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type ShowActivities struct {
	Slug  string `json:"slug"`
	Token string `json:"token"`
}
