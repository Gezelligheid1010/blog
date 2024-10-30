package mysql

import (
	"bluebell_backend/models"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreateComment(comment *models.Comment) (err error) {
	sqlStr := `insert into comment(
	comment_id, content, post_id, author_id, parent_id,likes)
	values(?,?,?,?,?,0)`
	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID,
		comment.AuthorID, comment.ParentID)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func GetCommentListByIDs(ids []string) (commentList []*models.Comment, err error) {
	sqlStr := `select comment_id, content, post_id, author_id, parent_id, create_time
	from comment
	where comment_id in (?)`
	// 动态填充id
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	// sqlx.In 返回带 `?` bindVar 的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&commentList, query, args...)
	return
}

func GetCommentList(postID int64, page, size int) (commentList []*models.Comment, err error) {
	// 排序条件
	orderBy := "create_time DESC" // 默认按最新排序

	// 计算偏移量
	offset := (page - 1) * size

	// 构建查询语句，包含排序和分页
	sqlStr := fmt.Sprintf(`SELECT comment_id, content, post_id, author_id, parent_id, create_time, likes
		FROM comment
		WHERE post_id = ? 
		ORDER BY %s
		LIMIT ? OFFSET ?`, orderBy)

	// 查询评论列表
	err = db.Select(&commentList, sqlStr, postID, size, offset)
	return
}

// IncrementCommentLike 增加评论的点赞数
func IncrementCommentLike(commentID int64) error {
	sqlStr := `UPDATE comment SET likes = likes + 1 WHERE comment_id = ?`

	fmt.Println(commentID)

	// 执行更新操作
	result, err := db.Exec(sqlStr, commentID)
	if err != nil {
		zap.L().Error("failed to increment comment like", zap.Error(err))
		return err
	}

	// 检查更新是否影响了任何行（确保 comment_id 存在）
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		zap.L().Error("failed to check rows affected", zap.Error(err))
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows // 表示未找到该评论
	}

	return nil
}
