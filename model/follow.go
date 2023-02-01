package model

import (
	"errors"
	"time"
)

type FollowModel struct{}

var followModel FollowModel

type Follow struct {
	Id       int64     `json:"id"`
	UserId   int64     `json:"userId"`
	FollowId int64     `json:"followId"`
	CreateAt time.Time `json:"createAt"`
}

func (f FollowModel) TableName() string {
	return "tb_follow"
}

func NewFollowModel() *FollowModel {
	return &followModel
}

// AddFollowAndCount 对用户点赞并更新计数
func (f FollowModel) AddFollowAndCount(userId, followId int64) error {

	return nil
}

// CancelFollowAndCount 取消点赞并更新计数
func (f FollowModel) CancelFollowAndCount(userId, followId int64) error {

	return nil
}

// QueryFollowById 根据用户id查询该用户关注列表
func (f FollowModel) QueryFollowById(userId int64, follow *[]*Follow) error {
	if follow == nil {
		return errors.New("QueryFollowById : null pointer exception")
	}
	if err := DB.Model(&Follow{}).Where("user_id = ?", userId).Find(follow).Error; err != nil {
		return err
	}
	return nil
}

// QueryFansById 根据用户id查询该用户的粉丝列表
func (f FollowModel) QueryFansById(userId int64) error {

	return nil
}
