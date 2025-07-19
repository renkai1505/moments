package main

import (
	"github.com/kingwrcy/moments/handler"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/do/v2"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func setupRouter(injector do.Injector) {
	userHandler := handler.NewUserHandler(injector)
	memoHandler := handler.NewMemoHandler(injector)
	commentHandler := handler.NewCommentHandler(injector)
	sycConfigHandler := handler.NewSysConfigHandler(injector)
	fileHandler := handler.NewFileHandler(injector)
	tagHandler := handler.NewTagHandler(injector)
	rssHandler := handler.NewRssHandler(injector)
	childHandler := handler.NewChildHandler(injector)
	growthHandler := handler.NewGrowthHandler(injector)
	milestoneHandler := handler.NewMilestoneHandler(injector)
	e := do.MustInvoke[*echo.Echo](injector)
	cfg := do.MustInvoke[*vo.AppConfig](injector)

	apiGroup := e.Group("/api")

	userGroup := apiGroup.Group("/user")
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/reg", userHandler.Reg)
	userGroup.POST("/profile", userHandler.Profile)
	userGroup.POST("/profile/:username", userHandler.ProfileForUser)
	userGroup.POST("/saveProfile", userHandler.SaveProfile)

	memoGroup := apiGroup.Group("/memo")
	memoGroup.POST("/list", memoHandler.ListMemos)
	memoGroup.POST("/save", memoHandler.SaveMemo)
	memoGroup.POST("/remove", memoHandler.RemoveMemo)
	memoGroup.POST("/like", memoHandler.LikeMemo)
	memoGroup.POST("/get", memoHandler.GetMemo)
	memoGroup.POST("/setPinned", memoHandler.SetPinned)
	memoGroup.POST("/getFaviconAndTitle", memoHandler.GetFaviconAndTitle)
	memoGroup.POST("/getDoubanMovieInfo", memoHandler.GetDoubanMovieInfo)
	memoGroup.POST("/getDoubanBookInfo", memoHandler.GetDoubanBookInfo)
	memoGroup.POST("/removeImage", memoHandler.RemoveImage)

	commentGroup := apiGroup.Group("/comment")
	commentGroup.POST("/add", commentHandler.AddComment)
	commentGroup.POST("/remove", commentHandler.RemoveComment)

	sycConfigGroup := apiGroup.Group("/sysConfig")
	sycConfigGroup.POST("/save", sycConfigHandler.SaveConfig)
	sycConfigGroup.POST("/get", sycConfigHandler.GetConfig)
	sycConfigGroup.POST("/getFull", sycConfigHandler.GetFullConfig)

	tagGroup := apiGroup.Group("/tag")
	tagGroup.POST("/list", tagHandler.List)

	fileGroup := apiGroup.Group("/file")
	fileGroup.POST("/upload", fileHandler.Upload)
	fileGroup.POST("/s3PreSigned", fileHandler.S3PreSigned)

	uploadGroup := e.Group("/upload")
	uploadGroup.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       cfg.UploadDir,
		HTML5:      false,
		IgnoreBase: true,
		Browse:     false,
	}))

	rssGroup := e.Group("/rss")
	rssGroup.GET("", rssHandler.GetRss)

	friendHandler := handler.NewFriendHandler(injector)
	friendGroup := apiGroup.Group("/friend")
	friendGroup.POST("/list", friendHandler.GetFriendList)
	friendGroup.POST("/add", friendHandler.AddFriend)
	friendGroup.POST("/delete", friendHandler.DeleteFriend)

	// 儿童档案相关路由
	childGroup := apiGroup.Group("/child")
	childGroup.POST("/list", childHandler.List)
	childGroup.POST("/get/:id", childHandler.Get)
	childGroup.POST("/create", childHandler.Create)
	childGroup.POST("/update/:id", childHandler.Update)
	childGroup.POST("/delete/:id", childHandler.Delete)

	// 成长记录相关路由
	growthGroup := apiGroup.Group("/growth")
	growthGroup.POST("/list", growthHandler.List)
	growthGroup.POST("/get/:id", growthHandler.Get)
	growthGroup.POST("/create", growthHandler.Create)
	growthGroup.POST("/update/:id", growthHandler.Update)
	growthGroup.POST("/delete/:id", growthHandler.Delete)

	// 里程碑相关路由
	milestoneGroup := apiGroup.Group("/milestone")
	milestoneGroup.POST("/list", milestoneHandler.List)
	milestoneGroup.POST("/get/:id", milestoneHandler.Get)
	milestoneGroup.POST("/create", milestoneHandler.Create)
	milestoneGroup.POST("/update/:id", milestoneHandler.Update)
	milestoneGroup.POST("/delete/:id", milestoneHandler.Delete)

	if cfg.EnableSwagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

}
