/*
 * @Author: CuiYao
 * @Date: 2021-12-10 16:31:55
 * @Last Modified by: CuiYao
 * @Last Modified time: 2022-01-28 10:25:47
 */

package main

import (
	"fmt"
	"github.com/CuiYao631/mini_program-server-go/controller"
	_ "github.com/CuiYao631/mini_program-server-go/docs"
	"github.com/CuiYao631/mini_program-server-go/repository"
	"github.com/CuiYao631/mini_program-server-go/usecase"
	"github.com/CuiYao631/mini_program-server-go/wechatGongZhong"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sashabaranov/go-openai"
	echoSwagger "github.com/swaggo/echo-swagger"
	"html/template"
	"log"
	"net"
	"os"
)

// @title 标题
// @version 版本号:(v1.0)
// @description 描述
// @contact.name 联系人
// @contact.url  联系网址
// @contact.email 联系人邮箱
// @license.name (Apache 2.0)
// @host localhost:8082
// @BasePath /
func main() {

	//ctx := context.Background()
	e := echo.New()
	e.Use(middleware.Logger())

	//sql
	//client, err := ent.Open("postgres", os.Getenv("COURSE_PLAN_POSTGRESQL_DSN"))
	//if err != nil {
	//	log.Fatal("connect sql failed", err)
	//}
	//defer client.Close()
	//if err = client.Schema.Create(ctx, migrate.WithDropIndex(true)); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}
	//redis

	//minio
	minioClient, err := minio.New(os.Getenv("ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("ACCESSKEYID"), os.Getenv("ACCESSKEY"), ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("minio failed: %v", err)
	}
	//创建一个服务端实例
	listen, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	// ChatGPT-3
	client := openai.NewClient(os.Getenv("OPEN_AI_API_KEY"))
	//设置静态资源url前缀和目录
	//这里设置 /static 为静态资源url的前缀，当前程序运行目录下面的static目录为静态资源目录
	e.Static("/static", "static")

	//初始化模版引擎
	t := &controller.Template{
		//模版引擎支持提前编译模版, 这里对views目录下以html结尾的模版文件进行预编译处理
		//预编译处理的目的是为了优化后期渲染模版文件的速度
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}

	//向echo实例注册模版引擎
	e.Renderer = t

	response := repository.MakeRepository(nil)
	pro := usecase.MakeUsecase(response, minioClient, client)
	ctrl := controller.MakeController(pro, listen, client)

	wgc := wechatGongZhong.MakeWechatGongZhong(client)
	//route

	e.GET("/", ctrl.Root)

	//公众号小程序回调
	e.GET("/wechat", wgc.EchoProcRequest)
	e.POST("/wechat", wgc.WXMsgReceive)

	e.POST("home", ctrl.Home)
	e.POST("/homeImage", ctrl.HoneWallpaper)
	//user
	g := e.Group("/user")
	ctrl.UserRoute(g)
	//resources
	r := e.Group("/recs")
	ctrl.ResourcesRoute(r)
	//wallpaper
	w := e.Group("/wallpaper")
	ctrl.WallpaperRoute(w)
	//minio
	m := e.Group("/minio")
	ctrl.MinioRoute(m)
	//ChatGPT
	e.GET("/ws", ctrl.GPTChat)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8082"))
}
