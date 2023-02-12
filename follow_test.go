package model

import (
	"testing"
)

var followListResult = []*Follow{
	{
		UserId:   100,
		FollowId: 101,
	},
	{
		UserId:   100,
		FollowId: 102,
	},
}

func TestFollowModel_TableName(t *testing.T) {
	if "tb_follow" != followModel.TableName() {
		t.FailNow()
	}
}

func TestFollowModel_QueryFollowById(t *testing.T) {
	var followList []*Follow
	if err := followModel.QueryFollowById(followListResult[0].UserId, &followList); err != nil {
		t.Errorf("QueryFollowById is fail, %v", err)
	}
	for i, follow := range followList {
		if follow.UserId != followListResult[i].UserId {
			t.Errorf("QueryFollowById is fail, userId is wrong !")
		}
		if follow.FollowId != followListResult[i].FollowId {
			t.Errorf("QueryFollowById is fail, followId is wrong !")
		}
	}
}

func TestFollowModel_QueryFansById(t *testing.T) {
	var followList []*Follow
	if err, _ := followModel.QueryFansById(followListResult[0].FollowId, &followList); err != nil {
		t.Errorf("QueryFansById is fail, %v", err)
	}
	for i, follow := range followList {
		if follow.UserId != followListResult[i].UserId {
			t.Errorf("QueryFansById is fail, userId is wrong !")
		}
		if follow.FollowId != followListResult[i].FollowId {
			t.Errorf("QueryFansById is fail, followId is wrong !")
		}
	}
}
