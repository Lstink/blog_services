package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lstink/blog/global"
	"github.com/lstink/blog/internal/model"
	"github.com/lstink/blog/internal/routers"
	"github.com/lstink/blog/pkg/setting"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(global.ServerSetting)
	fmt.Println(global.AppSetting)
	fmt.Println(global.DatabaseSetting)
	// 设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeOut,
		WriteTimeout:   global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		return
	}

}

// init 初始化设置
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

// setupSetting 读取配置信息
func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	// 设置读取的超时时间
	global.ServerSetting.ReadTimeOut *= time.Second
	// 设置写入的超时时间
	global.ServerSetting.WriteTimeOut *= time.Second
	return nil
}

// setupDBEngine 设置数据库驱动
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
