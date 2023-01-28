package model

import "time"

type videoModel struct{}

var VideoModel videoModel

type Video struct {
	Id            string    `json:"id"`            // 视频id
	Title         string    `json:"title"`         // 标题
	AutherId      int64     `json:"autherId"`      // 作者
	PlayId        int64     `json:"playId"`        // 播放地址
	CoverId       int64     `json:"coverId"`       // 视频封面地址
	FavoriteCount int64     `json:"favoriteCount"` // 收到的喜欢数目
	CommentCount  int64     `json:"commentCount"`  // 评论数
	CreateAt      time.Time `json:"createAt"`      // 创建时间
	UpdateAt      time.Time `json:"updateAt"`      // 更新时间
}
