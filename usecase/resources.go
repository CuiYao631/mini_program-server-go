/*
 * @Author: CuiYao
 * @Date: 2021-12-22 13:45:37
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 15:10:20
 */

package usecase

import (
	"context"
	"mime/multipart"
	"strings"

	"github.com/CuiYao631/mini_program-server-go/entity"
)

type Resources interface {
	UploadResourcesIcon(ctx context.Context, bucketName string, file *multipart.FileHeader) (string, error)
	CreateResources(ctx context.Context, resources entity.Resources) error
	UpdateResources(ctx context.Context, resources entity.Resources) error
	ListResources(ctx context.Context) ([]entity.Resources, error)
	GetResources(ctx context.Context, id string) (entity.Resources, error)
	DeleteResources(ctx context.Context, id string) error
}

func (uc *usecase) UploadResourcesIcon(ctx context.Context, bucketName string, file *multipart.FileHeader) (string, error) {
	bucket, fileName, err := uc.UploadWallpaper(ctx, bucketName, file)
	if err != nil {
		return "", err
	}
	url, err := uc.GetWallpaper(ctx, bucket, fileName)

	return url, nil
}

func (uc *usecase) CreateResources(ctx context.Context, resources entity.Resources) error {
	return uc.repo.CreateResources(ctx, resources)
}

func (uc *usecase) UpdateResources(ctx context.Context, resources entity.Resources) error {
	return uc.repo.UpdateResources(ctx, resources)
}

func (uc *usecase) ListResources(ctx context.Context) ([]entity.Resources, error) {
	entRes, err := uc.repo.ListResources(ctx)
	if err != nil {
		return nil, err
	}
	listres := make([]entity.Resources, 0, len(entRes))
	for _, v := range entRes {
		res := entity.Resources{
			ID:       v.ID,
			Icon:     v.Icon,
			Name:     v.Name,
			Tag:      "",
			Desc:     v.Desc,
			Url:      v.URL,
			Explain:  v.Explain,
			CreateAt: v.CreatedAt,
			UpdateAt: v.UpdatedAt,
		}
		listres = append(listres, res)
	}
	return listres, nil
}

func (uc *usecase) GetResources(ctx context.Context, id string) (entity.Resources, error) {
	entres, err := uc.repo.GetResources(ctx, id)
	if err != nil {
		return entity.Resources{}, err
	}
	res := entity.Resources{
		ID:       entres.ID,
		Icon:     entres.Icon,
		Name:     entres.Name,
		Tag:      "",
		Desc:     entres.Desc,
		Url:      entres.URL,
		CreateAt: entres.CreatedAt,
		UpdateAt: entres.UpdatedAt,
	}
	return res, nil
}

func (uc *usecase) DeleteResources(ctx context.Context, id string) error {
	entres, err := uc.repo.GetResources(ctx, id)
	if err != nil {
		return err
	}
	trimStr := ""
	str := entres.Icon[0:5]
	if str[len(str)-1:] == ":" {
		trimStr = strings.Trim(entres.Icon, "http://"+uc.host)
	} else {
		trimStr = strings.Trim(entres.Icon, "https://"+uc.host)
	}
	strArr := strings.Split(trimStr, "/")
	err = uc.DeleteWallpaper(ctx, strArr[0], strArr[1])
	if err != nil {
		return err
	}
	return uc.repo.DeleteResources(ctx, id)
}
