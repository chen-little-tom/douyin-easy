package model

import "time"

type videoModel struct{}

var VideoModel videoModel

type Video struct {
	Id            string `json:"id"`            // 视频id
	Title         string `json:"title"`         // 标题
	AuthorId      uint64 `json:"authorId"`      // 作者
	PlayId        uint64 `json:"playId"`        // 播放地址
	CoverId       uint64 `json:"coverId"`       // 视频封面地址
	FavoriteCount uint64 `json:"favoriteCount"` // 收到的喜欢数目
	CommentCount  uint64 `json:"commentCount"`  // 评论数
	Model
}

func (vm videoModel) Add(video Video) error {
	return DB.Model(Video{}).Create(&video).Error
}

func (vm videoModel) ListByTime(lastTime time.Time) ([]Video, error) {
	var rows []Video
	tx := DB.Model(Video{}).Order("create_at desc")
	tx.Where("create_at < ?", lastTime).Limit(10).Find(&rows)
	return rows, tx.Error
}
