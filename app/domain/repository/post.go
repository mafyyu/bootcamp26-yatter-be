package repository

import (
	"context"
	"yatter-backend-go/app/domain/object/post"
)

type Post interface {
	// 新規投稿を保存してする
	// 成功した場合には投稿を返す
	// contextとpostパッケージにあるPendingPost構造体のポインタ型を渡す
	// 返り値としてpostパッケージにあるPendingPost構造体のポインタ型とエラーを返す
	Insert(ctx context.Context, pendingPost *post.PendingPost) (*post.Post, error)
}