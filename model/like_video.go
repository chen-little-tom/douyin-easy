package model

import (
	"errors"
	"gorm.io/gorm"
)

type LikeVideo struct {
	Id       uint64 `json:"id" gorm:"column:id;primaryKey"`
	UserId   uint64 `json:"user_id" gorm:"column:user_id;"`
	VideoId  uint64 `json:"video_id" gorm:"column:video_id;"`
	CreateAt uint64 `json:"creat_at" gorm:"column:creat_at"`
}

// LikeVideo 增加视频喜欢数并将视频id加入视频喜欢列表
func (vm videoModel) LikeVideo(video LikeVideo) error {
	return DB.Transaction(func(lo *gorm.DB) error {
		//添加视频喜欢
		if err := lo.Create(lo).Error; err != nil {
			return err
		}

		if err := lo.Exec("UPDATE tb_video v SET v.favorite_count = v.favorite_count + 1 WHERE v.id = ?", video).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

// UnlikeVideo 减少视频喜欢数并将视频id从视频喜欢列表里删除
func (vm videoModel) UnlikeVideo(video LikeVideo) error {
	return DB.Transaction(func(lo *gorm.DB) error {

		if err := lo.Create(lo).Error; err != nil {
			return err
		}

		if err := lo.Exec("UPDATE tb_video v SET v.favorite_count = v.favorite_count - 1 WHERE v.id = ?", lo).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

// QueryLikeVideoById 根据id查询视频喜欢列表
func (vm videoModel) QueryLikeVideoById(Id uint64, likevideo *[]*LikeVideo) error {
	if likevideo == nil {
		return errors.New("QueryLikeVideoById : null pointer exception")
	}
	if err := DB.Model(&LikeVideo{}).Where("id = ?", Id).Find(likevideo).Error; err != nil {
		return err
	}
	return nil
}
