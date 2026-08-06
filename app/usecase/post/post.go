package post

import (
	"context"
	"yatter-backend-go/app/domain/object/post"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/usecase/transactor"
	"yatter-backend-go/pkg/errors"
)

// FindByUsernameを呼び出して、下の層にuser_idだけを流す

// UIがusecaseの実装を見れないように隠蔽するため
type CreatePostUsecase interface {
	Post(ctx context.Context, username string, content string) (*post.Post, error)
}

type PostUsecaseImpl struct {
	postRepo   repository.Post
	userRepo   repository.User
	transactor transactor.Transactor
}

func NewPostUseCase(postRepo repository.Post, userRepo repository.User, transactor transactor.Transactor) *PostUsecaseImpl {
	return &PostUsecaseImpl{
		postRepo:   postRepo,
		userRepo:   userRepo,
		transactor: transactor,
	}
}

func (uc *PostUsecaseImpl) Post(ctx context.Context, username string, content string) (*post.Post, error) {
	result, err := uc.transactor.TransactionWithValue(ctx, func (ctx context.Context) (any, error)  {
		
		u, err := uc.userRepo.FindByUsername(ctx, username)
		if err != nil {
			return nil, err
		}

		pendingPost, err := post.NewPendingPost(u.ID(), content)
		if err != nil {
			return nil, err
		}

		pst, err := uc.postRepo.Insert(ctx, pendingPost)
		if err != nil {
			return nil, err
		}
		return pst, nil
	})
	if err != nil {
		return nil, err
	}

	post, ok:= result.(*post.Post)
	if !ok{
		return nil, errors.ErrInternal.WithDevMessage("failed to cast result to post.Post")
	}
	return post, nil
}
