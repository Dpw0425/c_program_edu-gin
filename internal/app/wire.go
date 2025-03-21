package app

import (
	"c_program_edu-gin/internal/app/storage/cache"
	"c_program_edu-gin/internal/app/storage/repo"
	"c_program_edu-gin/internal/job"
	"github.com/google/wire"
)

var SQLProviderSet = wire.NewSet(
	wire.Struct(new(job.SQLProvider), "*"),
)

var AdminProviderSet = wire.NewSet(
	wire.Struct(new(job.AdminProvider), "*"),
)

var CacheProviderSet = wire.NewSet(
	cache.NewEmailStorage,
	cache.NewTokenSessionStorage,
)

var RepoProviderSet = wire.NewSet(
	repo.NewUsers,
	repo.NewAdmins,
)
