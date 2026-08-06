package infra

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object/post"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/infra/transaction"
)

// おまじない
var _ repository.Post = (*PostRepoImpl)(nil)

type PostRepoImpl struct{}

func NewPostRepository() *PostRepoImpl {
	return &PostRepoImpl{}
}

// 挿入されたデータを返すための構造体
type insertedPostDTO struct {
	Id        uint64    `db:"id"`
	UserId    uint64    `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

func (pr *PostRepoImpl) Insert(ctx context.Context, pendingPost *post.PendingPost) (*post.Post, error) {
	tx, err := transaction.FetchTransaction(ctx)
	if err != nil {
		return nil, err
	}

	insertResult, err := tx.ExecContext(
		ctx,
		`INSERT INTO yweet (user_id, content) VALUES (?,?)`,
		pendingPost.UserId(),
		pendingPost.Content(),
	)
	if err != nil {
		return nil, err
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	var insertedPostDTO insertedPostDTO
	err = tx.GetContext(
		ctx,
		&insertedPostDTO,
		`SELECT id, user_id, content, created_at FROM yweet WHERE id=?`,
		id,
	)
	if err != nil {
		return nil, err
	}

	pst, err := post.ReconstructPost(
		insertedPostDTO.Id,
		insertedPostDTO.UserId,
		insertedPostDTO.Content,
		insertedPostDTO.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return pst, nil
}
