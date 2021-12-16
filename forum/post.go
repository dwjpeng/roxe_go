package forum

import (
	roxe "github.com/dwjpeng/roxe_go"
)

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster roxe.AccountName, postUUID, content string, replyToPoster roxe.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *roxe.Action {
	a := &roxe.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []roxe.PermissionLevel{
			{Actor: poster, Permission: roxe.PermissionName("active")},
		},
		ActionData: roxe.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `roxe.forum::post` action.
type Post struct {
	Poster          roxe.AccountName `json:"poster"`
	PostUUID        string           `json:"post_uuid"`
	Content         string           `json:"content"`
	ReplyToPoster   roxe.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string           `json:"reply_to_post_uuid"`
	Certify         bool             `json:"certify"`
	JSONMetadata    string           `json:"json_metadata"`
}
