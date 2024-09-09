package request

type PushNotification struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	RId     int    `json:"r_id"`
	UserId  int    `json:"user_id"`
}

type EmailNotification struct {
	Email   string `json:"email"`
	Title   string `json:"title"`
	Content string `json:"content"`
	RId     int    `json:"r_id"`
	UserId  int    `json:"user_id"`
}

type SmsNotification struct {
	PhoneNumber string `json:"phone_number"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	RId         int    `json:"r_id"`
	UserId      int    `json:"user_id"`
}
