package request

type AccSessionAuth struct {
	Token string `validate:"required" json:"token"`
}
