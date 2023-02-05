package cache

import (
	"douyin-easy/model"
	"encoding/json"
	"errors"
)

type userRedisCache struct{}

var UserCache userRedisCache

func (uc userRedisCache) Get(key string) (model.User, error) {
	var user model.User
	redisKey := uc.getPrefix() + key
	u, err := Conn.Do("Get", redisKey)
	if err != nil {
		return user, err
	}
	user, err = uc.DescSerialize(u)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Put 加入缓存
func (uc userRedisCache) Put(key string, user model.User) (bool, error) {
	redisKey := uc.getPrefix() + key
	userJson, err := uc.Serialize(user)
	if err != nil {
		return false, err
	}
	reply, err := Conn.Do("set", redisKey, userJson)
	if err != nil {
		return false, err
	}
	// 判断是否设置成功
	status := reply.(string)
	if status == ReplyOk {
		return true, nil
	}
	return false, nil
}

func (uc userRedisCache) getPrefix() string {
	return "USER_"
}

// Serialize 序列化user
func (uc userRedisCache) Serialize(user model.User) (string, error) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	ret := string(userJson)
	return ret, nil
}

// DescSerialize 反序列化
func (uc userRedisCache) DescSerialize(userJson interface{}) (model.User, error) {
	user := model.User{}
	if userJson == nil {
		return user, errors.New("用户信息为空，未登录")
	}
	data := string(userJson.([]byte))
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
