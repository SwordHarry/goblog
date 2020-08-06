package main

import (
	"context"
	"flag"
	"fmt"
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
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// 命令行配置项
var (
	port    string
	runMode string
	config  string
	// 版本信息
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

// @title 博客系统
// @version 1.0
// @description Go 编程之旅：一起用 Go 做项目
// @termOfService https://github.com/SwordHarry/goblog
func main() {
	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}
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

	// 优雅重启与关停
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("main s.ListenAndServe err: %v", err)
		}
	}()
	// 等待中断信号，注意是 os.Singal
	quit := make(chan os.Signal)
	// 接收 syscall.SIGINT 和 syscall.SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 最大时间控制，通知服务端有 5 s 时间处理原有请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// 配置初始化
func init() {
	// 先读取命令行参数
	_ = setupFlag()
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
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
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
	if err = newSetting.ReadSection("Tracer", &global.TracerSetting); err != nil {
		return err
	}
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
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
	jaegerTracer, _, err := tracer.NewJaegerTracer(global.TracerSetting.ServiceName, global.TracerSetting.HostPort)
	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}

// 初始化命令行配置
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
	return nil
}
