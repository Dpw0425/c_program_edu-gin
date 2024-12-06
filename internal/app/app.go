package app

import (
	"c_program_edu-gin/internal/config"
	"github.com/urfave/cli/v2"
	"os"
)

// 封装应用程序
type App struct {
	app *cli.App
}

// 定义命令执行行为
type Action func(ctx *cli.Context, conf *config.Config) error

// 定义命令行参数
type Command struct {
	Name        string
	Usage       string
	Flags       []cli.Flag
	Action      Action
	SubCommands []Command
}

// 创建应用程序实例
func NewApp() *App {
	return &App{
		app: &cli.App{
			Name:    "CProgramEdu",
			Usage:   "",
			Version: "v0.0.241205",
		},
	}
}

// 注册新命令
func (c *App) Register(cm Command) {
	c.app.Commands = append(c.app.Commands, c.command(cm))
}

// 解析命令行参数
func (c *App) command(cm Command) *cli.Command {
	cmd := &cli.Command{
		Name:  cm.Name,
		Usage: cm.Usage,
		Flags: make([]cli.Flag, 0),
	}

	// 批处理
	if len(cm.SubCommands) > 0 {
		for _, v := range cm.SubCommands {
			cmd.Subcommands = append(cmd.Subcommands, c.command(v))
		}
	} else {
		if cm.Flags != nil && len(cm.Flags) > 0 {
			cmd.Flags = append(cmd.Flags, cm.Flags...)
		}

		var isConfig bool

		for _, flag := range cmd.Flags {
			if flag.Names()[0] == "config" {
				isConfig = true
				break
			}
		}

		if !isConfig {
			cmd.Flags = append(cmd.Flags, &cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "./config.yaml",
				Usage:       "配置文件路径",
				DefaultText: "./config.yaml",
			})
		}

		if cm.Action != nil {
			cmd.Action = func(ctx *cli.Context) error {
				return cm.Action(ctx, config.Load(ctx.String("config")))
			}
		}
	}

	return cmd
}

// 启动命令行应用程序
func (c *App) Run() {
	err := c.app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
