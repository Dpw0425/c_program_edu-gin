package app

import (
	"c_program_edu-gin/internal/app/storage/cache"
	"github.com/google/wire"
)

var CacheProviderSet = wire.NewSet(
	cache.NewTokenSessionStorage,
)
