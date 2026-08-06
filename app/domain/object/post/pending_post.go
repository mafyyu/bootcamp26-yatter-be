package post

import (
	"unicode/utf8"
	"yatter-backend-go/pkg/errors"
)

// ユーザーが入力したyatterを登録するため

type PendingPost struct {
	userId   uint64
	content  string
}

func NewPendingPost(userId uint64, content string) (*PendingPost, error) {
	post := &PendingPost{}

	if err := post.SetUserId(userId); err != nil {
		return nil, err
	}

	if err := post.SetContent(content); err != nil {
		return nil, err
	}

	return post, nil
}


func (p *PendingPost) SetUserId(userId uint64) error {
	// userIdが1以上であること
	if userId < 1 {
		return errors.ErrInternal.WithDevMessage("userId must be more than 1")
	}
	p.userId = userId
	return nil
}

func (p *PendingPost) SetContent(content string) error {
	// lenを使うとbyte数になってしまう
	// contentが1文字以上であること
	if utf8.RuneCountInString(content) < 1 {
		return errors.ErrInternal.WithDevMessage("content must be more than 1")
	}
	// contentが120字以内であること
	if utf8.RuneCountInString(content) > 120 {
		return errors.ErrInternal.WithDevMessage("content is limited to 120 characters.")
	}

	p.content = content

	return nil
}

func (p *PendingPost) UserId() uint64{
	return p.userId
}

func (p *PendingPost) Content() string {
	return p.content
}
