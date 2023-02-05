package model

import (
	"errors"
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

	return nil
}

// CancelFollowAndCount 取消点赞并更新计数
func (f FollowModel) CancelFollowAndCount(userId, followId uint64) error {

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
func (f FollowModel) QueryFansById(userId uint64) error {

	return nil
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
