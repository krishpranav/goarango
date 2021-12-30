package requests

import "fmt"

type FollowCursor struct {
	Cursor string
}

func (r *FollowCursor) Path() string {
	return fmt.Sprintf("/_api/cursor/%s", r.Cursor)
}

func (r *FollowCursor) Method() string {
	return "PUT"
}
