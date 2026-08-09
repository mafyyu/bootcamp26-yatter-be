package public

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/usecase/query"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

var _ Handler = (*GetHandlerImple)(nil)

type GetHandlerImple struct {
	TimelinesPublicQueryService query.TimelinesPublicQueryService
}

func NewGetHandler(TimelinesPublicQueryService query.TimelinesPublicQueryService) Handler {
	return &GetHandlerImple{
		TimelinesPublicQueryService: TimelinesPublicQueryService,
	}
}

func (h *GetHandlerImple) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	TimelinesPublicRequest, err := NewTimePublicRequest(limit,offset)
	if err != nil {
		return
	}

	result, err := h.TimelinesPublicQueryService.TimelinesPublicQueryService(ctx, TimelinesPublicRequest.limit, TimelinesPublicRequest.offset)
	if err != nil {
		return
	}
	fmt.Print(result)

	resp := toTimelinePublicResponse(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return
	}

}
