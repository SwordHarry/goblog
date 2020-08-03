package main

import (
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/internal/model"
	"goblog/internal/routers"
	"goblog/pkg/logger"
	"goblog/pkg/setting"
	"goblog/pkg/tracer"
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
	if err = setupDBEngine(); err != nil {
		log.Fatal(err)
	}
	if err = setupLogger(); err != nil {
		log.Fatal(err)
	}
	if err = setupTracer(); err != nil {
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
	if err = newSetting.ReadSection("JWT", &global.JWTSetting); err != nil {
		return err
	}
	if err = newSetting.ReadSection("Email", &global.EmailSetting); err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
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

// 初始化链路追踪器
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("go-blog", "127.0.0.1:6831")
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}
