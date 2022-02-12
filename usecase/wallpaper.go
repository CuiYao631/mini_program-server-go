package usecase

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/CuiYao631/mini_program-server-go/entity"
	"github.com/minio/minio-go/v7"
)

type Wallpaper interface {
	UploadWallpaper(ctx context.Context, bucketName string, file *multipart.FileHeader) error
	ListWallpaper(ctx context.Context) (entity.Wallpaper, error)
	GetWallpaper(ctx context.Context, id string) error
	DeleteWallpaper(ctx context.Context, id string) error
}

func (uc *usecase) UploadWallpaper(ctx context.Context, bucketName string, file *multipart.FileHeader) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	uploadInfo, err := uc.minioClient.PutObject(ctx, bucketName, file.Filename, src, file.Size, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(uploadInfo)

	return nil
}

func (uc *usecase) ListWallpaper(ctx context.Context) (entity.Wallpaper, error) {
	ct, cancel := context.WithCancel(ctx)

	defer cancel()

	objectCh := uc.minioClient.ListObjects(ct, "wallpaper", minio.ListObjectsOptions{
		Prefix:    "",
		Recursive: true,
	})
	links := make([]string, 0, len(objectCh))
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return entity.Wallpaper{}, object.Err
		}
		//fmt.Println(object.Key)
		links = append(links, "https://tencent.xcuitech.com:1688/wallpaper/"+object.Key)
	}
	wallpaper := entity.Wallpaper{
		Links: links,
	}
	return wallpaper, nil
}

func (uc *usecase) GetWallpaper(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}

func (uc *usecase) DeleteWallpaper(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
