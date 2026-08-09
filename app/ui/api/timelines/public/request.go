package public

import "strconv"

type TimelinesPublicRequest struct {
	// offsetは何個飛ばしか、limitはいくつ取得するか

	limit  uint64
	offset uint64
}

func NewTimePublicRequest(offset string, limit string) (*TimelinesPublicRequest, error) {
	convertedLimit, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		return nil, err
	}

	convertedOffset, err := strconv.ParseUint(offset, 10, 64)
	if err != nil {
		return nil, err
	}

	return &TimelinesPublicRequest{
		limit : convertedLimit,
		offset : convertedOffset,
	}, nil
}
