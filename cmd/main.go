package main

import (
	"c_program_edu-gin/internal/app"
	"c_program_edu-gin/internal/app/api"
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/internal/job"
	"c_program_edu-gin/pkg/logger"
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewHttpCommand() app.Command {
	return app.Command{
		Name:  "http",
		Usage: "http command - Http API 接口服务",
		Action: func(ctx *cli.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/logs/app_log", conf.Log.Path), "http")
			return api.RunHttpServer(ctx, NewHttpInjector(conf))
		},
	}
}

func NewSQLCommand() app.Command {
	return app.Command{
		Name:  "sql",
		Usage: "sql command - 数据库初始化",
		Action: func(ctx *cli.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/logs/app_log", conf.Log.Path), "sql")
			return job.Run(ctx, NewSQLInjector(conf))
		},
	}
}

func main() {
	newApp := app.NewApp()
	newApp.Register(NewHttpCommand())
	newApp.Register(NewSQLCommand())
	newApp.Run()
}
