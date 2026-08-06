package post

import (
	"time"
	"unicode/utf8"
	"yatter-backend-go/pkg/errors"
)

// 登録した投稿についてDBから生データを取得

// 挿入されたデータを返すための構造体
type Post struct {
	id        uint64
	userId    uint64
	content   string
	createdAt time.Time
}

func ReconstructPost(id uint64, userId uint64, content string, createAt time.Time) (*Post, error) {
	// 実体にする
	post := &Post{}

	if err := post.SetId(id); err != nil {
		return nil, err
	}

	if err := post.SetUserId(userId); err != nil {
		return nil, err
	}

	if err := post.SetContent(content); err != nil {
		return nil, err
	}

	if err := post.SetCreatedAt(createAt); err != nil {
		return nil, err
	}

	return post, nil
}

func (p *Post) SetId(id uint64) error {
	// idは1以上の整数であること
	if id < 1 {
		return errors.ErrInternal.WithDevMessage("id must be more than 0")
	}

	p.id = id
	return nil
}

func (p *Post) SetUserId(userId uint64) error {
	// userIdが1以上であること
	if userId < 1 {
		return errors.ErrInternal.WithDevMessage("userId must be more than 1")
	}
	p.userId = userId
	return nil
}

func (p *Post) SetContent(content string) error {
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

func (p *Post) SetCreatedAt(createdAt time.Time) error {
	// createdAtはYatterサービス開始時以降であること
	yatterLaunchedAt := time.Date(2025, 4, 1, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	if !createdAt.After(yatterLaunchedAt) {
		return errors.ErrInternal.WithDevMessage("createdAt must be after yatter launched")
	}

	// createdAtは未来の日付でないこと
	if createdAt.After(time.Now()) {
		return errors.ErrInternal.WithDevMessage("createdAt must not be in the future")
	}

	// createdAtのタイムゾーンはJSTであること
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	jstCreatedAt := createdAt.In(jst)
	if !createdAt.Equal(jstCreatedAt) {
		return errors.ErrInternal.WithDevMessage("createdAt must be in JST")
	}

	p.createdAt = createdAt

	return nil
}

func (p *Post) ID() uint64 {
	return p.id
}

func (p *Post) Content() string {
	return p.content
}

func (p *Post) CreatedAt() time.Time {
	return p.createdAt
}
