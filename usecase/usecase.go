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
	"github.com/sashabaranov/go-openai"
	"os"
)

var (
	Host = os.Getenv("APIHOST")
)

type Usecase interface {
	Resources
	Minio
	Wallpaper
	ChatGpt
}
type usecase struct {
	repo        repository.Repository
	minioClient *minio.Client
	host        string
	opAi        *openai.Client
}

func MakeUsecase(repo repository.Repository, minioClient *minio.Client, opAi *openai.Client) *usecase {
	return &usecase{repo: repo, minioClient: minioClient, host: Host, opAi: opAi}
}
