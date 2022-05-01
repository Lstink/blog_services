package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lstink/blog/global"
	"github.com/lstink/blog/internal/model"
	"github.com/lstink/blog/internal/routers"
	"github.com/lstink/blog/pkg/logger"
	"github.com/lstink/blog/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version 1.0.0
// @description Go 语言编程之旅，一起做GO项目
// @termsOfService https://github.com/lstink/go
func main() {
	// 设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	// 实例化路由
	router := routers.NewRouter()
	// 设置服务参数
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	// 启动并监听服务
	err := s.ListenAndServe()
	if err != nil {
		return
	}

}

// init 初始化设置
func init() {
	// 设置全局配置信息
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// 设置数据库配置信息
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	// 设置日志驱动
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// setupSetting 读取配置信息
func setupSetting() error {
	// 初始化配置实例
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// 读取 Server 配置，并赋值给 global.ServerSetting 这个全局变量
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	// 读取 App 配置，并赋值给 global.AppSetting 这个全局变量
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	// 读取 Database 配置，并赋值给 global.DatabaseSetting 这个全局变量
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	// 设置读取的超时时间单位为秒
	global.ServerSetting.ReadTimeOut *= time.Second
	// 设置写入的超时时间单位为秒
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

// setupDBEngine 设置数据库驱动
func setupDBEngine() error {
	// 声明一个错误类型的变量
	var err error
	// 设置全局的 EBEngine 变量为数据库实例
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// setupLogger 设置日志驱动
func setupLogger() error {
	// 拼装日志名称
	filename := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	// 设置全局日志实例
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
