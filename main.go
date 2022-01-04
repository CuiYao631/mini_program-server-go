/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:55
 * @Last Modified by: CuiYao
 * @Last Modified time: 2021-12-22 16:15:08
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
	// e.Use(middleware.Recover())

	client, err := ent.Open("postgres", os.Getenv("COURSE_PLAN_POSTGRESQL_DSN"))
	if err != nil {
		log.Fatal("connect sql failed", err)
	}
	defer client.Close()
	if err = client.Schema.Create(ctx, migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	response := repository.MakeRepository(client)
	pro := usecase.MakeUsecase(response)
	ctrl := controller.MakeController(pro)
	e.GET("/", ctrl.Root)
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
	e.Logger.Fatal(e.Start(":8082"))

}
