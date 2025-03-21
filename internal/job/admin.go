package job

import (
	"c_program_edu-gin/internal/app/storage/model"
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/pkg/logger"
	"c_program_edu-gin/utils/encrypt"
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type AdminProvider struct {
	Config *config.Config
	DB     *gorm.DB
}

func InitAdmin(ctx *cli.Context, conf *AdminProvider) error {
	fmt.Println("正在初始化管理系统...")
	var count int64
	conf.DB.Table("admins").Count(&count)
	if count == 0 {
		if err := conf.DB.Table("admins").Create(&model.Admin{
			UserName:   "Admin",
			TeacherID:  conf.Config.Admin.UserName,
			Password:   encrypt.HashPassword(conf.Config.Admin.Password),
			Permission: "any",
		}).Error; err != nil {
			logger.Info("Admins init failed!")
			fmt.Println("管理系统初始化失败！", err.Error())
		} else {
			fmt.Println("管理系统初始化成功！")
			logger.Info("Admins init successful!")
		}
	} else {
		logger.Info("Super administrator already exists!")
		fmt.Println("超级管理员已存在！")
	}
	return nil
}
