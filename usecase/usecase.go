/*
 * @Author: CuiYao
 * @Date: 2021-09-07 09:37:10
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 14:22:03
 */
package usecase

import "github.com/CuiYao631/mini_program-server-go/repository"

type Usecase interface {
	Resources
}
type usecase struct {
	repo repository.Repository
}

func MakeUsecase(repo repository.Repository) *usecase {
	return &usecase{repo: repo}
}
