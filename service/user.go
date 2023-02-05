package service

import (
	"douyin-easy/cache"
	"douyin-easy/model"
	"douyin-easy/utils"
	"douyin-easy/utils/user_token"
	"douyin-easy/vo"
	"errors"
	"log"
	"strings"
)

type userService struct{}

var UserService userService

// UserLogin 用户登录
func (us userService) UserLogin(username, password string) (uint64, string, error) {
	user, err := model.UserModel.FindByUsername(username)
	if err != nil {
		log.Printf("查询用户失败,username->%s,err->%s\n", username, err)
		return 0, "", errors.New("未找到用户信息")
	}
	if user.Id == 0 {
		return 0, "", errors.New("未找到用户信息")
	}
	err = utils.CryptMatch(password, user.Password)
	if err != nil {
		return 0, "", errors.New("账号或密码错误")
	}
	return us.doLogin(user)
}

// UserRegister 用户注册
func (us userService) UserRegister(username, password string) (uint64, string, error) {
	user, err := model.UserModel.FindByUsername(username)
	if err != nil {
		log.Printf("查询用户失败,username->%s,err->%s\n", username, err)
		return 0, "", errors.New("注册失败")
	}
	if user.Id != 0 {
		return 0, "", errors.New("用户名已经被注册")
	}
	user.Username = username
	encode, err := utils.CryptEncode(password)
	if err != nil {
		log.Printf("用户密码加密失败,username->%s,err->%s\n", username, err)
		return 0, "", errors.New("用户注册失败")
	}
	user.Password = *encode
	err = model.UserModel.Add(user)
	if err != nil {
		log.Printf("用户信息存入数据库失败,err->%s\n", err)
		return 0, "", errors.New("用户注册失败")
	}
	return us.doLogin(user)
}

// doLogin 登录缓存操作
func (us userService) doLogin(user model.User) (uint64, string, error) {
	// 执行
	token, err := user_token.GetUserToken(user.Id)
	if err != nil {
		return 0, "", errors.New("获取token失败")
	}
	_, err = cache.UserCache.Put(token, user)
	if err != nil {
		log.Printf("缓存登录用户失败,err->%s\n", err)
		return 0, "", errors.New("登录失败")
	}
	return user.Id, token, nil
}

// Detail 用户详细信息
func (us userService) Detail(uId uint64, token string) (vo.UserVo, error) {
	var ret vo.UserVo
	// 查询是否有关联
	u, err := model.UserModel.Detail(uId)
	if err != nil {
		log.Printf("查询用户出现错误,err->%s\n", err)
		return ret, errors.New("查询用户出现错误")
	}
	if u.Id == 0 {
		return ret, errors.New("未找到用户信息")
	}
	ret.Id = u.Id
	ret.FollowCount = u.FollowCount
	ret.FollowerCount = u.FansCount
	ret.Name = u.Username
	ret.IsFollow = false
	return ret, nil
}

// DetailByIdsMap 批量查询用户详情
// Return map k->userId v->userInfo
func (us userService) DetailByIdsMap(userIds []uint64, loginUser model.User) (map[uint64]vo.UserVo, error) {
	ret := make(map[uint64]vo.UserVo, len(userIds))
	users, err := model.UserModel.ListByIds(userIds)
	if err != nil {
		log.Printf("批量查询用户信息失败,err->%s\n", err)
		return ret, err
	}
	for _, u := range users {
		var user vo.UserVo
		user.Id = u.Id
		user.Name = u.Username
		user.IsFollow = false
		user.FollowCount = u.FollowCount
		user.FollowerCount = u.FansCount
		ret[u.Id] = user
	}
	return ret, nil
}

// GetLoginUser 获取当前登录用户信息
func (us userService) GetLoginUser(token string) (model.User, error) {
	token = strings.ReplaceAll(token, " ", "+")
	user, err := cache.UserCache.Get(token)
	if err != nil {
		log.Printf("获取当前登录用户信息失败,err->%s\n", err)
		return user, errors.New("获取当前登录用户信息失败")
	}
	if user.Id == 0 {
		log.Printf("用户未登录")
		return user, errors.New("用户未登录")
	}
	return user, nil
}
