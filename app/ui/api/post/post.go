package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	ui_errors "yatter-backend-go/app/ui/api/pkg/errors"
	"yatter-backend-go/app/usecase/post"
	"yatter-backend-go/pkg/errors"
)

type Handler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

var _ Handler = (*postHandlerImpl)(nil)

type postHandlerImpl struct {
	postCreateUsecase post.CreatePostUsecase
}

// 実体にする
func NewPostHandler(postCreateUsecase post.CreatePostUsecase) Handler {
	return &postHandlerImpl{
		postCreateUsecase: postCreateUsecase,
	}
}

func (h *postHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var req PostRequst
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	ctx := r.Context()

	username := strings.Split(r.Header.Get("Authentication"), " ")[1]

	pst, err := h.postCreateUsecase.Post(ctx, username, req.Content)
	if err != nil {
		return
	}
	
	resp := toPostResponse(pst)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		ui_errors.Handle(w, errors.ErrInternal.WithDevMessage(fmt.Sprintf("failed to encode response: %s", err.Error())))
		return
	}
}
