package main

import (
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/internal/model"
	"goblog/internal/routers"
	"goblog/pkg/logger"
	"goblog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version 1.0
// @description Go 编程之旅：一起用 Go 做项目
// @termOfService https://github.com/SwordHarry/goblog
func main() {
	// 设置 gin 的运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// 配置初始化
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatal(err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatal(err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatal(err)
	}
}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = newSetting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = newSetting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = newSetting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
