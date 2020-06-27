package router

import (
	"fmt"
	"github.com/KubeOperator/KubeOperator/pkg/middleware"
	"github.com/KubeOperator/KubeOperator/pkg/router/proxy"
	v1 "github.com/KubeOperator/KubeOperator/pkg/router/v1"
	"github.com/KubeOperator/KubeOperator/pkg/router/xpack"
	"github.com/kataras/iris/v12"
)

func Server() *iris.Application {
	app := iris.New()
	err := app.I18n.Load("./locales/*/*.yml", "en-US", "zh-CN")
	if err != nil {
		fmt.Println(err.Error())
	}
	app.I18n.SetDefault("zh-CN")
	app.Post("/api/auth/login", middleware.LoginHandler)
	app.Get("/api/auth/profile", middleware.JWTMiddleware().Serve, middleware.GetAuthUser)
	proxy.RegisterProxy(app)
	api := app.Party("/api")
	api.Use(middleware.PagerMiddleware)
	//api.Use(middleware.LogMiddleware)
	api.Use(middleware.JWTMiddleware().Serve)
	v1.V1(api)
	xpack.XPack(api)
	return app
}
