package service

import (
	"context"
	"fmt"

	user "github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/crypto/bcrypt"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (checkUserService *CheckUserService) Login(ctx context.Context, session neo4j.SessionWithContext, user *user.CheckUserRequest) (int64, error) {
	userid, err := neo4j.ExecuteRead[int64](ctx, session, func(tx neo4j.ManagedTransaction) (int64, error) {
		result, err := tx.Run(ctx, "MATCH (n:User {username: $username}) RETURN n.password AS ps, id(n) AS i LIMIT 1", map[string]any{
			"username": user.Username,
		})
		if err != nil {
			return 0, err
		}
		record, err := result.Single(ctx)
		if err != nil {
			return 0, errno.AuthorizationFailedErr
		}
		fmt.Println(record.Values[0], record.Values[1])
		if !BcryptCheck(user.Password, record.Values[0].(string)) {
			return 0, errno.AuthorizationFailedErr
		}
		return record.Values[1].(int64), nil
	})
	if err != nil {
		return 0, err
	}
	return userid, nil
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
