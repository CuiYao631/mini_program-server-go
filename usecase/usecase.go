/*
 * @Author: CuiYao
 * @Date: 2021-09-07 09:37:10
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-10 16:53:21
 */
package usecase

import "github.com/CuiYao631/mini_program-server-go/repository"

type Usecase interface {
}
type usecase struct {
	repo repository.Repository
}

func MakeUsecase(repo repository.Repository) *usecase {
	return &usecase{repo: repo}
}
