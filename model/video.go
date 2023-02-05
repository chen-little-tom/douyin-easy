package model

import "time"

type videoModel struct{}

var VideoModel videoModel

type Video struct {
	Id            uint64 `json:"id" gorm:"column:id;primaryKey"`             // 视频id
	Title         string `json:"title" gorm:"column:title"`                  // 标题
	AuthorId      uint64 `json:"authorId" gorm:"column:author_id"`           // 作者
	PlayId        uint64 `json:"playId" gorm:"column:play_id"`               // 播放地址
	CoverId       uint64 `json:"coverId" gorm:"column:cover_id"`             // 视频封面地址
	FavoriteCount uint64 `json:"favoriteCount" gorm:"column:favorite_count"` // 收到的喜欢数目
	CommentCount  uint64 `json:"commentCount" gorm:"column:comment_count"`   // 评论数
	Model
}

func (vm videoModel) Add(video Video) error {
	return DB.Model(&Video{}).Create(&video).Error
}

func (vm videoModel) ListByTime(lastTime time.Time) ([]Video, error) {
	var rows []Video
	tx := DB.Model(&Video{}).Order("create_at desc")
	tx.Where("create_at < ?", lastTime).Limit(10).Find(&rows)
	return rows, tx.Error
}

func (vm videoModel) GetAuthorIdByVideoId(videoId uint64) (uint64, error) {
	var video Video
	tx := DB.Model(Video{})
	tx.Where("id = ?", videoId).Find(&video)
	return video.AuthorId, nil
}
