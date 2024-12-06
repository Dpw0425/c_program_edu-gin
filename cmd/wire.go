//go:build wireinject
// +build wireinject

package main

import (
	"c_program_edu-gin/internal/app/api"
	"c_program_edu-gin/internal/config"
	"github.com/google/wire"
)

//var providerSet = wire.NewSet(
//	app.CacheProviderSet,
//)

func NewHttpInjector(conf *config.Config) *api.AppProvider {
	panic(
		wire.Build(
			// providerSet,
			api.ProviderSet,
		),
	)
}
