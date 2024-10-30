package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Comment struct {
	PostID     uint64    `db:"post_id" json:"post_id"`
	ParentID   uint64    `db:"parent_id" json:"parent_id"`
	CommentID  string    `db:"comment_id" json:"comment_id"`
	AuthorID   uint64    `db:"author_id" json:"author_id"`
	Content    string    `db:"content" json:"content"`
	Likes      string    `db:"likes" json:"likes"`
	Author     User      `db:"author" json:"author"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}

// UnmarshalJSON 自定义解析逻辑，将字符串类型的 post_id 解析为 uint64
func (c *Comment) UnmarshalJSON(data []byte) error {
	type Alias Comment // 避免递归调用 UnmarshalJSON

	aux := &struct {
		PostID string `json:"post_id"` // 先作为字符串解析
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// 将字符串转换为 uint64
	postID, err := strconv.ParseUint(aux.PostID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid post_id: %v", err)
	}
	c.PostID = postID

	return nil
}
