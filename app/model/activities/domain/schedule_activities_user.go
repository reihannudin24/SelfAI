package domain

type ScheduleActivitiesUser struct {
	ID       int    `json:"id"`
	Schedule string `json:"schedule"`
	Message  string `json:"message"`
	UserId   int    `json:"user_id"`
}
