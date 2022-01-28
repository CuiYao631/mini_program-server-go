/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:55
 * @Last Modified by: CuiYao
 * @Last Modified time: 2022-01-28 10:25:47
 */

package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/CuiYao631/mini_program-server-go/controller"
	"github.com/CuiYao631/mini_program-server-go/ent"
	"github.com/CuiYao631/mini_program-server-go/ent/migrate"
	"github.com/CuiYao631/mini_program-server-go/repository"
	"github.com/CuiYao631/mini_program-server-go/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const Intervaltime = 20 * time.Minute
const IntervalNum = 3

func main() {
	// res := retryablehttp.NewClient()
	// //最大时间间隔时间为20分钟
	// res.RetryWaitMax = Intervaltime
	// //重新连接次数为3次
	// res.RetryMax = IntervalNum
	// Client := res.StandardClient()
	// var rest request.GetAccTokenResponse
	// request.GetAccessToken(Client, "wxb07ca737e71b5b7b", "077af2c2fa57aa32e7295bc40f188a54", &rest)
	// request.GetWeiXinJsapi_Ticket(Client, rest.AccessToken, &rest)
	// log.Println(rest.Ticket)

	ctx := context.Background()
	e := echo.New()
	e.Use(middleware.Logger())

	//sql
	client, err := ent.Open("postgres", os.Getenv("COURSE_PLAN_POSTGRESQL_DSN"))
	if err != nil {
		log.Fatal("connect sql failed", err)
	}
	defer client.Close()
	if err = client.Schema.Create(ctx, migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//redis

	//minio
	minioClient, err := minio.New(os.Getenv("ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("ACCESSKEYID"), os.Getenv("ACCESSKEY"), ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("minio failed: %v", err)
	}

	response := repository.MakeRepository(client)
	pro := usecase.MakeUsecase(response, minioClient)
	ctrl := controller.MakeController(pro)

	//route
	e.GET("/", ctrl.Root)
	e.POST("home", ctrl.Home)
	g := e.Group("/user")
	g.POST("/add", ctrl.CreateUser)
	g.POST("/update", ctrl.UpdateUser)
	g.POST("/list", ctrl.ListUser)
	g.POST("/delete", ctrl.DeleteUser)
	r := e.Group("/recs")
	r.POST("/create", ctrl.CreateResources)
	r.POST("/update", ctrl.UpdateResources)
	r.POST("/list", ctrl.ListResources)
	r.POST("/get", ctrl.GetResources)
	r.POST("/del", ctrl.DeleteResources)

	m := e.Group("/minio")
	ctrl.MinioRoute(m)

	e.Logger.Fatal(e.Start(":8082"))
}
