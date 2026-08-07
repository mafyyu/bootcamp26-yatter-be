package public

import (
	"time"
	"yatter-backend-go/app/domain/object/timelines/public"
)

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

type TimelinesPublicItem struct {
	ID        uint64    `json:"id"`
	User      User      `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type TimelinesPublicResponse []TimelinesPublicItem

func toTimelinePublicResponse(p []public.TimelinesPublic) TimelinesPublicResponse {
	var response TimelinesPublicResponse
	for _, resp := range p {
		var r = TimelinesPublicItem{
			ID: resp.PostID,
			User: User{
				ID:        resp.UserID,
				Username:  resp.Username,
				CreatedAt: resp.UserCreatedAt,
			},
			Content:   resp.Content,
			CreatedAt: resp.PostCreatedAt,
		}
		response = append(response, r)
	}

	return response
}
