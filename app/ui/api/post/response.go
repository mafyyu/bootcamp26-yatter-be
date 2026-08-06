package post

import (
	"time"
	"yatter-backend-go/app/domain/object/post"
)

type Response struct {
	ID        uint64    `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func toPostResponse(post *post.Post) *Response {
	return &Response{
		ID:        post.ID(),
		Content:   post.Content(),
		CreatedAt: post.CreatedAt(),
	}
}
