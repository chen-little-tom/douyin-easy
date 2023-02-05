package service

import (
	"douyin-easy/model"
	"time"
)

type User struct {
	Id            uint64 `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint64 `json:"follow_count"`
	FollowerCount uint64 `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type CommentActionResponse struct {
	Id         uint64 `json:"id"`
	User       *User  `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func PostComment(fatherId, toUserId, videoId, fromUserId uint64, content string) (*CommentActionResponse, error) {
	comment := model.Comment{
		FatherId:   fatherId,
		ToUserId:   toUserId,
		VideoId:    videoId,
		FromUserId: fromUserId,
		Content:    content,
		CreateAt:   time.Now(),
	}
	err := model.NewCommentModel().AddCommentAndUpdateCount(&comment)
	if err != nil {
		return nil, err
	}
	var user User
	userDB, _ := model.UserModel.Detail(fromUserId)
	isFollow, _ := model.FollowModel{}.QueryIsFollow(fromUserId, toUserId)
	user.Id, user.Name, user.FollowCount, user.FollowerCount, user.IsFollow = userDB.Id, userDB.Username, userDB.FollowCount, userDB.FansCount, isFollow
	return &CommentActionResponse{Id: comment.Id, User: &user, Content: comment.Content, CreateDate: comment.CreateAt.Format("2006.01.02 15:04:05")}, nil
}

func DeleteComment(commentId, videoId, fromUserId, toUserId uint64) (*CommentActionResponse, error) {
	var comment model.Comment

	//获取评论
	err := model.NewCommentModel().QueryCommentById(commentId, &comment)
	if err != nil {
		return nil, err
	}

	//删除评论
	err = model.NewCommentModel().DeleteCommentAndUpdateCountById(commentId, videoId)
	if err != nil {
		return nil, err
	}

	var user User
	userDB, _ := model.UserModel.Detail(fromUserId)
	isFollow, _ := model.FollowModel{}.QueryIsFollow(fromUserId, toUserId)
	user.Id, user.Name, user.FollowCount, user.FollowerCount, user.IsFollow = userDB.Id, userDB.Username, userDB.FollowCount, userDB.FansCount, isFollow

	return &CommentActionResponse{Id: comment.Id, User: &user, Content: comment.Content, CreateDate: comment.CreateAt.Format("2006.01.02 15:04:05")}, nil
}
