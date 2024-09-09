package domain

type FavoriteThings struct {
	ID             int    `json:"id"`
	FavoriteThings string `json:"favorite_thing"`
	UserId         int    `json:"user_id"`
}

type FavoriteThingsResponse struct {
	ID             int    `json:"id"`
	FavoriteThings string `json:"favorite_things"`
	UserId         int    `json:"user_id"`
}

func (data FavoriteThings) ToResponse() FavoriteThingsResponse {
	return FavoriteThingsResponse{
		ID:             data.ID,
		FavoriteThings: data.FavoriteThings,
		UserId:         data.UserId,
	}
}
