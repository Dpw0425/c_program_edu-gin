package api

import (
	myErr "c_program_edu-gin/pkg/error"
	"c_program_edu-gin/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunHttpServer(ctx *cli.Context, app *AppProvider) error {
	gin.SetMode(app.Config.App.RunMode)

	eg, groupCtx := errgroup.WithContext(ctx.Context)

	// 用于接收操作系统信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	logger.Infof("HTTP Listen Port: %d", app.Config.Server.Http)
	logger.Infof("HTTP Server Pid: %d", os.Getpid())

	return run(c, eg, groupCtx, app)
}

func run(c chan os.Signal, eg *errgroup.Group, ctx context.Context, app *AppProvider) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Server.Http),
		Handler: app.Engine,
	}

	// 启动 http 服务
	eg.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && !myErr.Is(err, http.ErrServerClosed) {
			logger.Errorf("Starting Server Error: %s", err)
			return err
		}

		return nil
	})

	eg.Go(func() error {
		defer func() {
			logger.Info("Shutting down server...")

			// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
			timeCtx, timeCancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer timeCancel()

			if err := server.Shutdown(timeCtx); err != nil {
				logger.Fatalf("HTTP Server Shutdown Error: %s", err)
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			return nil
		}
	})

	if err := eg.Wait(); err != nil && myErr.Is(err, context.Canceled) {
		logger.Fatalf("HTTP Server Forced To Shutdown: %s", err)
	}

	logger.Info("Server Exiting")

	return nil
}
