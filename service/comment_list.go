package service

type CommentListResponse struct {
	CommentList []*CommentActionResponse `json:"comment_list"`
}
