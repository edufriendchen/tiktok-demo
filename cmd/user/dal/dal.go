package dal

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/pkg/global"
	"gorm.io/gorm"
)

// User Gorm Data Structures
type User struct {
	gorm.Model
	UserName      string `gorm:"index:idx_username,unique;type:varchar(32);not null" json:"username"`
	Password      string `gorm:"type:varchar(32);not null" json:"password"`
	NickName      string `gorm:"type:varchar(32);not null" json:"nickName"`
	Avatar        string `gorm:"type:varchar(256)" json:"avatar"`
	FollowCount   int64  `gorm:"default:0" json:"follow_count"`
	FollowerCount int64  `gorm:"default:0" json:"follower_count"`
}

func (User) TableName() string {
	return "user"
}

// func (s *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	uid, err := uuid.NewV4()
// 	i := uid.String()[0:8]
// 	return
// }

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return global.DB.WithContext(ctx).Create(users).Error
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := global.DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetUserByID query list of user info by user_id
func GetUserByID(ctx context.Context, userID int64) (*User, error) {
	res := new(User)
	if err := global.DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := global.DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
