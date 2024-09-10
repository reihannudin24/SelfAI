package request

type CreateScheduleActivities struct {
	Schedule string `json:"schedule"`
	Message  string `json:"message"`
	UserId   int    `json:"user_id"`
}

type UpdateScheduleActivities struct {
	ID       int    `json:"id"`
	Schedule string `json:"schedule"`
	Message  string `json:"message"`
	UserId   int    `json:"user_id"`
}

type DeleteScheduleActivities struct {
	ID int `json:"id"`
}
