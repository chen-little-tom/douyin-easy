package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type FollowModel struct{}

var followModel FollowModel

type Follow struct {
	Id       uint64    `json:"id"`
	UserId   uint64    `json:"userId"`
	FollowId uint64    `json:"followId"`
	CreateAt time.Time `json:"createAt"`
}

func (f FollowModel) TableName() string {
	return "tb_follow"
}

func NewFollowModel() *FollowModel {
	return &followModel
}

// AddFollowAndCount 对用户点赞并更新计数
func (f FollowModel) AddFollowAndCount(userId, followId uint64) error {
	//执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		fol := Follow{FollowId: followId, UserId: userId, CreateAt: time.Now()}
		if err := tx.Create(fol).Error; err != nil {
			return err
		}
		//增加点赞计数
		if err := tx.Exec("UPDATE tb_user v SET v.follow_count = v.follow_count + 1 WHERE v.id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE tb_user v SET v.fans_count = v.fans_count + 1 WHERE v.id = ?", followId).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

// CancelFollowAndCount 取消点赞并更新计数
func (f FollowModel) CancelFollowAndCount(userId, followId uint64) error {

	return DB.Transaction(func(tx *gorm.DB) error {
		//执行事务
		fol := Follow{FollowId: followId, UserId: userId, CreateAt: time.Now()}
		if err := tx.Create(fol).Error; err != nil {
			return err
		}
		//增加点赞计数
		if err := tx.Exec("UPDATE tb_user v SET v.follow_count = v.follow_count - 1 WHERE v.id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE tb_user v SET v.fans_count = v.fans_count - 1 WHERE v.id = ?", followId).Error; err != nil {
			return err
		}
		return nil
	})
	return nil

}

// QueryFollowById 根据用户id查询该用户关注列表
func (f FollowModel) QueryFollowById(userId uint64, follow *[]*Follow) error {
	if follow == nil {
		return errors.New("QueryFollowById : null pointer exception")
	}
	if err := DB.Model(&Follow{}).Where("user_id = ?", userId).Find(follow).Error; err != nil {
		return err
	}
	return nil
}

// QueryFansById 根据用户id查询该用户的粉丝列表
func (f FollowModel) QueryFansById(userId uint64, fans *[]*Follow) (error, error) {
	if fans == nil {
		return errors.New("QueryFansById : null pointer exception"), nil
	}
	if err := DB.Model(&Follow{}).Where("follow_id = ?", userId).Find(fans).Error; err != nil {
		return err, nil
	}
	return nil, nil
}

func (f FollowModel) QueryIsFollow(userId, followId uint64) (bool, error) {
	follow := Follow{FollowId: followId, UserId: userId}
	var count int64
	if err := DB.Model(&follow).Where("user_id = ? AND follow_id = ?", follow.UserId, follow.FollowId).Count(&count).Error; err != nil {
		return false, err
	}
	if count >= 1 {
		return true, nil
	}
	return false, nil
}
