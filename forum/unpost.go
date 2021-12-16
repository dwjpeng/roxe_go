package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster roxe.AccountName, postUUID string) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("unpost"),
		Authorization: []roxe.PermissionLevel{
			{Actor: poster, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `roxe.forum::unpost` action.
type UnPost struct {
	Poster   roxe.AccountName `json:"poster"`
	PostUUID string           `json:"post_uuid"`
}
