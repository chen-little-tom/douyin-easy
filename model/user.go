package model

type userModel struct{}

var UserModel userModel

type User struct {
	Id               uint64 `json:"id" gorm:"column:id;primaryKey"`                    // 用户id
	Username         string `json:"username" gorm:"column:username"`                   // 用户名
	Password         string `json:"password" gorm:"column:password"`                   // 密码
	Nickname         string `json:"nickname" gorm:"column:nickname"`                   // 昵称
	FansCount        uint64 `json:"fansCount" gorm:"column:fans_count"`                // 粉丝数
	FollowCount      uint64 `json:"followCount" gorm:"column:follow_count"`            // 关注数
	ReceiveLikeCount uint64 `json:"receiveLikeCount" gorm:"column:receive_like_count"` // 收到的点赞数
}

// Add 添加用户
func (um userModel) Add(user User) error {
	return DB.Model(&User{}).Create(&user).Error
}

// FindByUsername 根据用户名称查询用户信息
func (um userModel) FindByUsername(username string) (User, error) {
	var user User
	tx := DB.Model(&User{}).Where("username = ?", username).Find(&user)
	return user, tx.Error
}

// Detail 查询用户详情
func (um userModel) Detail(id uint64) (User, error) {
	var user User
	tx := DB.Model(&User{}).Where("id = ?", id).Find(&user)
	return user, tx.Error
}

// ListByIds 批量查询用户信息
func (um userModel) ListByIds(userIds []uint64) ([]User, error) {
	var users []User
	tx := DB.Model(&User{}).Where("id IN ?", userIds).Find(&users)
	return users, tx.Error
}
