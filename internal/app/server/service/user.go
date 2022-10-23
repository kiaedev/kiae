package service

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/kiaedev/kiae/api/user"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSvc struct {
	userDao *dao.UserDao
}

func NewUserSvc(userDao *dao.UserDao) *UserSvc {
	return &UserSvc{userDao: userDao}
}

func (s *UserSvc) List(ctx context.Context, in *user.UserListRequest) (*user.UserListResponse, error) {
	results, total, err := s.userDao.List(ctx, bson.M{})
	return &user.UserListResponse{Items: results, Total: total}, err
}

func (s *UserSvc) saveFromOidcUserInfo(ctx context.Context, userInfo *oidc.UserInfo) (err error) {
	u, err := s.userDao.GetByOuterId(ctx, userInfo.Subject)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err != nil {
		u := &user.User{
			Email:   userInfo.Email,
			OuterId: userInfo.Subject,
		}
		_, err = s.userDao.Create(ctx, u)
		return
	}

	u.Email = userInfo.Email
	_, err = s.userDao.Update(ctx, u)
	return
}
