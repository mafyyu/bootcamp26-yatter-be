package public

import "time"

type TimelinesPublic struct {
	PostID        uint64
	UserID        uint64
	Username      string
	UserCreatedAt time.Time
	Content       string
	PostCreatedAt time.Time
}

func ReconstructTimelinesPublic(
	postId uint64,
	userId uint64,
	username string,
	userCreatedAt time.Time,
	content string,
	postCreatedAt time.Time,
) (
	*TimelinesPublic,
	error,
) {
	timelinesPublic := &TimelinesPublic{}

	if err := timelinesPublic.setPostId(postId); err != nil {
		return nil, err
	}

	if err := timelinesPublic.setUserId(userId); err != nil {
		return nil, err
	}

	if err := timelinesPublic.setUsername(username); err != nil {
		return nil, err
	}

	if err := timelinesPublic.setUserCreatedAt(userCreatedAt); err != nil {
		return nil, err
	}

	if err := timelinesPublic.setContent(content); err != nil {
		return nil, err
	}

	if err := timelinesPublic.setPostCreatedAt(postCreatedAt); err != nil {
		return nil, err
	}

	return timelinesPublic, nil
}

func (t *TimelinesPublic) setPostId(postId uint64) error {
	t.PostID = postId
	return nil
}

func (t *TimelinesPublic) setUserId(userId uint64) error {
	t.UserID = userId
	return nil
}

func (t *TimelinesPublic) setUsername(username string) error {
	t.Username = username
	return nil
}

func (t *TimelinesPublic) setUserCreatedAt(userCreatedAt time.Time) error {
	t.UserCreatedAt = userCreatedAt
	return nil
}

func (t *TimelinesPublic) setContent(content string) error {
	t.Content = content
	return nil
}

func (t *TimelinesPublic) setPostCreatedAt(postCreatedAt time.Time) error {
	t.PostCreatedAt = postCreatedAt
	return nil
}
