package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cuiyuanxin/airuisi/pkg/logger"
	"github.com/cuiyuanxin/airuisi/pkg/tracer"

	"github.com/cuiyuanxin/airuisi/internal/router"

	"github.com/cuiyuanxin/airuisi/global"
	"github.com/gin-gonic/gin"

	"github.com/cuiyuanxin/airuisi/pkg/setting"
)

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	//err = setting.ReadSection("Database", &global.DatabaseSetting)
	//if err != nil {
	//	return err
	//}
	err = setting.ReadSection("Logger", &global.LoggerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	//err = setting.ReadSection("JWT", &global.JwtSetting)
	//if err != nil {
	//	return err
	//}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Tracer", &global.TracerSetting)
	if err != nil {
		return err
	}
	//err = setting.ReadSection("Sms", &global.SmsSetting)
	//if err != nil {
	//	return err
	//}
	return nil
}

func setupLogger() error {
	var err error
	global.Logger, err = logger.NewLogger(global.LoggerSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("beifang-service", global.TracerSetting.AgentHostPort)
	if err != nil {
		return err
	}

	global.Tracet = jaegerTracer
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err：%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err：%v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracet err：%v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.GinMode)
	router := router.NewRouter()

	s := &http.Server{
		Addr:           global.ServerSetting.HttpAddr + ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err：%v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal)
	// 接受syscall.SIGINT和syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting tdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exiting")
}
