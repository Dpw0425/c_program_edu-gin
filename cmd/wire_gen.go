// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"c_program_edu-gin/internal/app"
	"c_program_edu-gin/internal/app/api"
	"c_program_edu-gin/internal/app/api/handler"
	"c_program_edu-gin/internal/app/api/handler/admin"
	v1_2 "c_program_edu-gin/internal/app/api/handler/admin/v1"
	"c_program_edu-gin/internal/app/api/handler/web"
	"c_program_edu-gin/internal/app/api/handler/web/v1"
	"c_program_edu-gin/internal/app/api/router"
	service2 "c_program_edu-gin/internal/app/service"
	"c_program_edu-gin/internal/app/service/admin/v1"
	"c_program_edu-gin/internal/app/service/web/v1"
	"c_program_edu-gin/internal/app/storage/cache"
	"c_program_edu-gin/internal/app/storage/repo"
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/internal/job"
	"c_program_edu-gin/internal/provider"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewHttpInjector(conf *config.Config) *api.AppProvider {
	client := provider.NewRedisClient(conf)
	emailStorage := cache.NewEmailStorage(client)
	db := provider.NewMysqlClient(conf)
	userRepo := repo.NewUsers(db)
	templateService := &service.TemplateService{}
	emailClient := provider.NewEmailClient(conf)
	emailService := service.EmailService{
		Storage:  emailStorage,
		UserRepo: userRepo,
		Template: templateService,
		Client:   emailClient,
	}
	common := &v1.Common{
		EmailService: emailService,
	}
	jwtTokenStorage := cache.NewTokenSessionStorage(client)
	userService := &service.UserService{
		UserRepo: userRepo,
	}
	serviceEmailService := &service.EmailService{
		Storage:  emailStorage,
		UserRepo: userRepo,
		Template: templateService,
		Client:   emailClient,
	}
	user := &v1.User{
		Config:          conf,
		JwtTokenStorage: jwtTokenStorage,
		UserService:     userService,
		EmailService:    serviceEmailService,
	}
	iFileSystem := provider.NewFileSystem(conf)
	upload := &v1.Upload{
		Config:     conf,
		FileSystem: iFileSystem,
	}
	webV1 := &web.V1{
		Common: common,
		User:   user,
		Upload: upload,
	}
	webHandler := &web.Handler{
		V1: webV1,
	}
	adminRepo := repo.NewAdmins(db)
	adminService := admin_service.AdminService{
		AdminRepo: adminRepo,
	}
	v1Admin := &v1_2.Admin{
		Config:       conf,
		AdminRepo:    adminRepo,
		AdminService: adminService,
	}
	adminV1 := &admin.V1{
		Admin: v1Admin,
	}
	adminHandler := &admin.Handler{
		V1: adminV1,
	}
	handlerHandler := &handler.Handler{
		Web:   webHandler,
		Admin: adminHandler,
	}
	engine := router.NewRouter(conf, handlerHandler, jwtTokenStorage)
	appProvider := &api.AppProvider{
		Config: conf,
		Engine: engine,
	}
	return appProvider
}

func NewSQLInjector(conf *config.Config) *job.SQLProvider {
	db := provider.NewMysqlClient(conf)
	sqlProvider := &job.SQLProvider{
		Config: conf,
		DB:     db,
	}
	return sqlProvider
}

// wire.go:

var providerSet = wire.NewSet(provider.NewHttpClient, provider.NewRequestClient, provider.NewMysqlClient, provider.NewRedisClient, provider.NewEmailClient, provider.NewFileSystem, service2.WebProviderSet, service2.AdminProviderSet, app.CacheProviderSet, app.RepoProviderSet)
