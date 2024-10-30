package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/models"
	"bluebell_backend/pkg/snowflake"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 评论

// CommentHandler 创建评论
func CommentHandler(c *gin.Context) {
	var comment models.Comment
	//fmt.Println("comment:", comment)
	if err := c.BindJSON(&comment); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 生成评论ID
	commentID, err := snowflake.GetID()
	if err != nil {
		zap.L().Error("snowflake.GetID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	comment.CommentID = strconv.FormatUint(commentID, 10)
	comment.AuthorID = userID

	fmt.Println("comment:", comment)
	// 创建评论
	if err := mysql.CreateComment(&comment); err != nil {
		fmt.Println(err)
		zap.L().Error("mysql.CreateComment(&comment) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"comment_id": commentID,
		"message":    "Comment created successfully",
	})
	//ResponseSuccess(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {

	post_id, err := strconv.ParseInt(c.DefaultQuery("post_id", "0"), 10, 64)
	fmt.Println("post_id:", post_id)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 获取分页和排序参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	// 获取评论列表
	comments, err := mysql.GetCommentList(post_id, page, size)
	if err != nil {
		zap.L().Error("mysql.GetCommentList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 格式化返回的数据，包含用户信息
	var commentResponses []models.Comment
	for _, comment := range comments {
		user, err := mysql.GetUserByID(comment.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID failed", zap.Error(err))
			continue
		}
		commentResponses = append(commentResponses, models.Comment{
			CommentID:  comment.CommentID,
			Content:    comment.Content,
			Likes:      comment.Likes,
			CreateTime: comment.CreateTime,
			Author: models.User{
				UserID:   user.UserID,
				UserName: user.UserName,
				Avatar:   user.Avatar,
			},
		})
	}
	ResponseSuccess(c, commentResponses)
}

func LikeCommentHandler(c *gin.Context) {
	// 获取评论ID
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 点赞操作
	if err := mysql.IncrementCommentLike(commentID); err != nil {
		zap.L().Error("mysql.IncrementCommentLike failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"message": "Comment liked successfully",
	})
}
