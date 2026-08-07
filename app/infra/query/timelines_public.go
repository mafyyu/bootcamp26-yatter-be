package query

import (
	"context"
	"fmt"
	"time"
	"yatter-backend-go/app/domain/object/timelines/public"
	"yatter-backend-go/app/usecase/query"

	"github.com/jmoiron/sqlx"
)

var _ query.TimelinesPublicQueryService = (*TimelinesPublicQueryServiceImpl)(nil)

type TimelinesPublicQueryServiceImpl struct {
	db *sqlx.DB
}

func NewTimelinesPublicQueryService(db *sqlx.DB) *TimelinesPublicQueryServiceImpl {
	return &TimelinesPublicQueryServiceImpl{
		db: db,
	}
}

type TimelinesPublicDTO struct {
	PostID        uint64    `db:"post_id"`
	UserID        uint64    `db:"user_id"`
	Username      string    `db:"username"`
	UserCreatedAt time.Time `db:"user_created_at"`
	Content       string    `db:"content"`
	PostCreatedAt time.Time `db:"post_created_at"`
}

func (t *TimelinesPublicQueryServiceImpl) TimelinesPublicQueryService(ctx context.Context) ([]public.TimelinesPublic, error) {
	// &でアドレスを渡してる

	var TimelinesPublicDTO []TimelinesPublicDTO
	err := t.db.SelectContext(
		ctx,
		&TimelinesPublicDTO,
		`
			SELECT 
				yweet.id AS post_id,
				user.id AS user_id,
				user.username,
				user.created_at AS user_created_at,
				yweet.content AS content,
				yweet.created_at AS post_created_at
			FROM yweet 
			LEFT JOIN user 
			ON yweet.user_id = user.id
		`,
	)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	var TimelinesPublic []public.TimelinesPublic

	for _, dto := range TimelinesPublicDTO {
		r, err := public.ReconstructTimelinesPublic(
			dto.PostID,
			dto.UserID,
			dto.Username,
			dto.UserCreatedAt,
			dto.Content,
			dto.PostCreatedAt,
		)
		if err != nil {
			return nil, err
		}
		TimelinesPublic = append(TimelinesPublic, *r)
	}

	return TimelinesPublic, nil
}
