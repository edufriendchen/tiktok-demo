package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	user "github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"

	"github.com/gofrs/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserService struct {
	ctx     context.Context
	session neo4j.SessionWithContext
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context, driver neo4j.DriverWithContext) *CreateUserService {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	return &CreateUserService{ctx: ctx, session: session}
}

func (s *CreateUserService) CreateUserNode(user *user.CreateUserRequest) (userid int64, err error) {
	user.Password = BcryptHash(user.Password)
	nickname := DefaultNickName()
	avatar := DefaultAvatar()
	userid, err = neo4j.ExecuteWrite[int64](s.ctx, s.session, func(tx neo4j.ManagedTransaction) (userid int64, err error) {
		result, err := tx.Run(s.ctx, "MATCH (n:User {username: $username}) RETURN count(*) AS count LIMIT 1", map[string]any{
			"username": user.Username,
		})
		if err != nil {
			return 0, err
		}
		record, err := result.Single(s.ctx)
		if err != nil {
			return 0, err
		}
		count, found := record.Get("count")
		op, _ := count.(int64)
		if !found {
			return 0, fmt.Errorf("could not find column")
		}
		if op != 0 {
			return 0, errno.UserAlreadyExistErr
		}
		result, err = tx.Run(s.ctx, "CREATE (n:User { username: $username, password: $password, nickname: $nickname, avatar: $avatar, follow_count: $follow_count, follower_count: $follower_count }) RETURN n", map[string]any{
			"username":       user.Username,
			"password":       user.Password,
			"nickname":       nickname,
			"avatar":         avatar,
			"follow_count":   0,
			"follower_count": 0,
		})
		if err != nil {
			return 0, err
		}
		record, err = result.Single(s.ctx)
		if err != nil {
			return 0, err
		}
		rawPerson, found := record.Get("n")
		if !found {
			return 0, fmt.Errorf("could not find column")
		}
		itemNode, ok := rawPerson.(neo4j.Node)
		if !ok {
			return 0, fmt.Errorf("expected result to be a map but was %T", rawPerson)
		}
		userid = itemNode.GetId()
		return userid, nil
	})
	if err != nil {
		return 0, err
	}
	return userid, nil
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// DefaultNickName 获得一个默认的昵称
func DefaultNickName() string {
	uid, _ := uuid.NewV4()
	return "用户" + uid.String()[0:8]
}

// DefaultAvatar 获得一个默认的头像
func DefaultAvatar() string {
	AvatarList := [...]string{
		"https://img1.baidu.com/it/u=1459539381,1684299919&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1675530000&t=264b8cfbd62ce23ee0d0a557091cc72d",
		"https://img0.baidu.com/it/u=2155033989,3634097964&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1675530000&t=a6973ebee4d25a8611a4efd711ed52a3",
		"https://img1.baidu.com/it/u=256830766,4270545878&fm=253&app=138&size=w931&n=0&f=JPG&fmt=auto?sec=1675530000&t=967010975a9f1ae6ef80e46fca09f713",
		"https://img2.baidu.com/it/u=324241668,3161137356&fm=253&app=138&size=w931&n=0&f=JPG&fmt=auto?sec=1675530000&t=7087110ff8179531e16664396e414809",
		"https://img1.baidu.com/it/u=4102089746,1733025287&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1675530000&t=e33dd1e45318b5402ae9317742c91c10",
		"https://img2.baidu.com/it/u=1659974989,4260768333&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1675530000&t=a60b66a45ff9585df06c491ae799ea74",
		"https://img2.baidu.com/it/u=346152429,3164401706&fm=253&app=138&size=w931&n=0&f=JPG&fmt=auto?sec=1675530000&t=b25ae1f4c313ef0010e95975bcc21c5d",
		"https://img1.baidu.com/it/u=1900416729,2440027599&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1675530000&t=85dfaac59b2febf3fa995735e00190b1",
	}
	rand.Seed(time.Now().UnixNano())
	return AvatarList[rand.Intn(8)]
}
