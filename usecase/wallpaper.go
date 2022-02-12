package usecase

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type Wallpaper interface {
	UploadWallpaper(ctx context.Context, bucketName string, file *multipart.FileHeader) error
	listWallpaper(ctx context.Context) error
	GetWallpaper(ctx context.Context, id string) error
	DeleteWallpaper(ctx context.Context, id string) error
}

func (uc *usecase) UploadWallpaper(ctx context.Context, bucketName string, file *multipart.FileHeader) error {

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	uploadInfo, err := uc.minioClient.PutObject(ctx, bucketName, file.Filename, src, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(uploadInfo)

	return nil
}

func (uc *usecase) listWallpaper(ctx context.Context) error {

	return nil
}

func (uc *usecase) GetWallpaper(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}

func (uc *usecase) DeleteWallpaper(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
