package cache

import (
	"github.com/go-macaron/cache"
	_ "github.com/go-macaron/cache/memcache"
	_ "github.com/go-macaron/cache/redis"
	//"github.com/jeffotoni/gmongocrud/conf"
	"os"
)

var (
	// Memory Option
	Memory = cache.Options{}

	// File Option
	File = cache.Options{
		Adapter: "file",
		//AdapterConfig: conf.Cfg.Section("").Key("cache_adpter_config").Value(),
		AdapterConfig: os.Getenv("cache_adpter_config"),
	}

	// Redis Option
	Redis = cache.Options{
		Adapter: "redis",
		//AdapterConfig: conf.Cfg.Section("").Key("cache_adpter_config").Value(),
		AdapterConfig: os.Getenv("cache_adpter_config"),
	}

	// Memcache Option
	Memcache = cache.Options{
		Adapter: "memcache",
		//AdapterConfig: conf.Cfg.Section("").Key("cache_adpter_config").Value(),
		AdapterConfig: os.Getenv("cache_adpter_config"),
	}
)
