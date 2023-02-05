package service

import "douyin-easy/model"

func GetCommentList(videoId uint64) (*[]*CommentActionResponse, error) {
	var comments []*model.Comment
	if err := model.NewCommentModel().QueryCommentListByVideoId(videoId, &comments); err != nil {
		return nil, err
	}
	var commentList []*CommentActionResponse
	for _, c := range comments {
		userDB, _ := model.UserModel.Detail(c.FromUserId)
		isFollow, _ := model.NewFollowModel().QueryIsFollow(userDB.Id, c.ToUserId)
		user := User{
			Id:            userDB.Id,
			Name:          userDB.Username,
			FollowCount:   userDB.FollowCount,
			FollowerCount: userDB.FansCount,
			IsFollow:      isFollow,
		}
		comment := CommentActionResponse{
			Id:         c.Id,
			Content:    c.Content,
			User:       &user,
			CreateDate: c.CreateAt,
		}
		commentList = append(commentList, &comment)
	}
	return &commentList, nil
}
