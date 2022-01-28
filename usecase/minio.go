/*
 * @Author: CuiYao
 * @Date: 2022-01-28 09:08:01
 * @Last Modified by: CuiYao
 * @Last Modified time: 2022-01-28 10:57:30
 */

package usecase

import (
	"context"
	"log"
	"net/url"
	"time"
)

type Minio interface {

	//*******Bucket操作
	//创建Bucket
	MakeBucket(ctx context.Context) error
	//列出Bucket
	ListBucket(ctx context.Context) error
	//Bucket是否存在
	ExistsBucket(ctx context.Context) error
	//删除Bucket
	Delete(ctx context.Context) error
	//列出Bucket中的对象
	ListObject(ctx context.Context) error

	//******Object 对象操作
	//文件下载
	FGetObject(ctx context.Context) error
	//文件上传
	FPutObject(ctx context.Context) error
	//流上传
	PutObject(ctx context.Context) error
	//流下载
	GetObject(ctx context.Context) error
	//获取URL
	GetObjectUrl(ctx context.Context, bucketName, objectName string) (string, error)
	//删除对象
	DeleteObject(ctx context.Context) error
}

//*******Bucket操作
//创建Bucket
func (uc *usecase) MakeBucket(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//列出Bucket
func (uc *usecase) ListBucket(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//Bucket是否存在
func (uc *usecase) ExistsBucket(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//删除Bucket
func (uc *usecase) Delete(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//列出Bucket中的对象
func (uc *usecase) ListObject(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//******Object 对象操作
//文件下载
func (uc *usecase) FGetObject(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//文件上传
func (uc *usecase) FPutObject(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//流上传
func (uc *usecase) PutObject(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//流下载
func (uc *usecase) GetObject(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

//获取URL
func (uc *usecase) GetObjectUrl(ctx context.Context, bucketName, objectName string) (string, error) {
	URL1, err := uc.minioClient.PresignedGetObject(ctx, bucketName, objectName, time.Second*24*60*60, make(url.Values))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return URL1.String(), nil
}

//删除对象
func (uc *usecase) DeleteObject(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}
