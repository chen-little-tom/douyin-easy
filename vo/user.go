package vo

type UserVo struct {
	Id            uint64 `json:"id"`             // 用户id
	Name          string `json:"name"`           // 用户名称
	FollowCount   uint64 `json:"follow_count"`   // 关注总数
	FollowerCount uint64 `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // 是否关注
}
