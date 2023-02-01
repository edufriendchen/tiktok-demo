package service

import (
	"context"
	"fmt"

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
		result, err = tx.Run(s.ctx, "CREATE (n:User { username: $username, password: $password, nickname: $nickname, follow_count: $follow_count, follower_count: $follower_count }) RETURN n", map[string]any{
			"username":       user.Username,
			"password":       user.Password,
			"nickname":       nickname,
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
