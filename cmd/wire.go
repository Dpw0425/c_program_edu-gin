//go:build wireinject
// +build wireinject

package main

import (
	"c_program_edu-gin/internal/app"
	"c_program_edu-gin/internal/app/api"
	"c_program_edu-gin/internal/app/service"
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/internal/job"
	"c_program_edu-gin/internal/provider"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	provider.NewHttpClient,
	provider.NewRequestClient,
	provider.NewMysqlClient,
	provider.NewRedisClient,
	provider.NewEmailClient,

	service.WebProviderSet,
	app.CacheProviderSet,
	app.RepoProviderSet,
)

func NewHttpInjector(conf *config.Config) *api.AppProvider {
	panic(
		wire.Build(
			providerSet,
			api.ProviderSet,
		),
	)
}

func NewSQLInjector(conf *config.Config) *job.SQLProvider {
	panic(
		wire.Build(
			providerSet,
			app.SQLProviderSet,
		),
	)
}
