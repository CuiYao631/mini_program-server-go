/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:53:33
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-10 17:04:32
 */
package usecase

import (
	"context"

	"github.com/CuiYao631/mini_program-server-go/entity"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, id string, user entity.User) (int, error)
	ListUser(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func (uc *usecase) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *usecase) UpdateUser(ctx context.Context, id string, user entity.User) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *usecase) ListUser(ctx context.Context) ([]entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (uc *usecase) DeleteUser(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
