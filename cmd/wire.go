//go:build wireinject
// +build wireinject

package main

import (
	"c_program_edu-gin/internal/app"
	"c_program_edu-gin/internal/app/api"
	"c_program_edu-gin/internal/app/service"
	"c_program_edu-gin/internal/config"
	"c_program_edu-gin/internal/provider"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	provider.NewHttpClient,
	provider.NewRequestClient,
	// provider.NewMysqlClient,
	provider.NewRedisClient,

	service.WebProviderSet,
	app.CacheProviderSet,
)

func NewHttpInjector(conf *config.Config) *api.AppProvider {
	panic(
		wire.Build(
			providerSet,
			api.ProviderSet,
		),
	)
}
