/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:30:56
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-10 16:51:22
 */
package repository

import (
	"context"

	"github.com/CuiYao631/mini_program-server-go/ent"
	"github.com/CuiYao631/mini_program-server-go/ent/user"
	"github.com/CuiYao631/mini_program-server-go/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (*ent.User, error)
	UpdateUser(ctx context.Context, id string, user entity.User) (int, error)
	ListUser(ctx context.Context) ([]*ent.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func (repo *repository) CreateUser(ctx context.Context, u entity.User) (*ent.User, error) {
	entUser, err := repo.db.User.Create().
		SetName(u.Name).
		SetPassword(u.Password).
		SetUserType(user.UserType(u.UserType.String())).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return entUser, nil
}
func (repo *repository) UpdateUser(ctx context.Context, id string, u entity.User) (int, error) {
	len, err := repo.db.User.Update().
		SetName(u.Name).
		SetPassword(u.Password).
		SetUserType(user.UserType(u.UserType.String())).
		Where(user.IDEQ(id)).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return len, nil
}
func (repo *repository) ListUser(ctx context.Context) ([]*ent.User, error) {
	ListEntUser, err := repo.db.User.Query().All(ctx)
	if err != nil {
		return []*ent.User{}, err
	}
	return ListEntUser, nil
}

func (repo *repository) DeleteUser(ctx context.Context, id string) error {
	if err := repo.db.User.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}
