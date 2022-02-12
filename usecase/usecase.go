/*
 * @Author: CuiYao
 * @Date: 2021-09-07 09:37:10
 * @Last Modified by: CuiYao
 * @Last Modified time: 2022-01-28 10:16:11
 */
package usecase

import (
	"github.com/CuiYao631/mini_program-server-go/repository"
	"github.com/minio/minio-go/v7"
)

type Usecase interface {
	Resources
	Minio
	Wallpaper
}
type usecase struct {
	repo        repository.Repository
	minioClient *minio.Client
}

func MakeUsecase(repo repository.Repository, minioClient *minio.Client) *usecase {
	return &usecase{repo: repo, minioClient: minioClient}
}
