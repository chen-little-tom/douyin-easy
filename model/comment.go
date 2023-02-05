package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type CommentModel struct{}

var commentModel CommentModel

type Comment struct {
	Id         uint64    `json:"commentId" gorm:"primary_key;AUTO_INCREMENT"` // 评论id
	FatherId   uint64    `json:"fatherId"`                                    // 评论父id
	ToUserId   uint64    `json:"toUserId"`                                    // 评论指向用户id
	VideoId    uint64    `json:"videoId"`                                     // 评论指向视频id
	FromUserId uint64    `json:"fromUserId"`                                  // 评论者id
	Content    string    `json:"content"`                                     // 评论内容
	CreateAt   time.Time `json:"createAt" gorm:"column:create_at"`            // 创建时间
}

func (c CommentModel) TableName() string {
	return "tb_comment"
}

func NewCommentModel() *CommentModel {
	return &commentModel
}

// AddCommentAndUpdateCount 添加评论并增加计数
func (c CommentModel) AddCommentAndUpdateCount(comment *Comment) error {
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount : null pointer exception")
	}
	//执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		//添加评论
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		//增加评论计数
		if err := tx.Exec("UPDATE tb_video v SET v.comment_count = v.comment_count + 1 WHERE v.id = ?", comment.VideoId).Error; err != nil {
			return err
		}

		return nil
	})
}

// DeleteCommentAndUpdateCountById 删除评论并减少计数
func (c CommentModel) DeleteCommentAndUpdateCountById(commentId, videoId uint64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		//删除评论
		if err := tx.Exec("DELETE FROM tb_comment WHERE id= ?", commentId).Error; err != nil {
			return err
		}
		//更新计数
		if err := tx.Exec("UPDATE tb_video v SET v.comment_count = v.comment_count - 1 WHERE v.id = ? AND v.comment_count > 0", videoId).Error; err != nil {
			return err
		}
		return nil
	})
}

// QueryCommentById 按照评论id查询评论
func (c CommentModel) QueryCommentById(commentId uint64, comment *Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById : null pointer exception")
	}
	return DB.Where("id = ?", commentId).First(comment).Error
}

// QueryCommentListByVideoId 按照视频id查询评论
func (c CommentModel) QueryCommentListByVideoId(videoId uint64, comments *[]*Comment) error {
	if comments == nil {
		return errors.New("QueryCommentListByVideoId : null pointer exception")
	}
	if err := DB.Model(&Comment{}).Where("video_id = ?", videoId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}
