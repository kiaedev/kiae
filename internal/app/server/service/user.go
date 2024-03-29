package service

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/kiaedev/kiae/api/user"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *UserSvc) Info(ctx context.Context, empty *emptypb.Empty) (*user.User, error) {
	return s.userDao.Get(ctx, MustGetUserid(ctx))
}

func (s *UserSvc) saveFromOidcUserInfo(ctx context.Context, userInfo *oidc.UserInfo) (err error) {
	extra := make(map[string]any)
	if err := userInfo.Claims(&extra); err != nil {
		return err
	}

	u := &user.User{
		Email:   userInfo.Email,
		OuterId: userInfo.Subject,
		Roles:   []string{"member"},
	}
	username, ok := extra["name"].(string)
	if ok {
		u.Nickname = username
	}
	avatar, ok := extra["picture"].(string)
	if ok {
		u.Avatar = avatar
	}

	eu, err := s.userDao.GetByOuterId(ctx, userInfo.Subject)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	} else if err != nil {
		_, err = s.userDao.Create(ctx, u)
		return
	}

	u.Id = eu.Id
	_, err = s.userDao.Update(ctx, u)
	return
}
