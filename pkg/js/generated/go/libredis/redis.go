package redis

import (
	lib_redis "github.com/devilsfang/nuclei/v3/pkg/js/libs/redis"

	"github.com/devilsfang/nuclei/v3/pkg/js/gojs"
	"github.com/dop251/goja"
)

var (
	module = gojs.NewGojaModule("nuclei/redis")
)

func init() {
	module.Set(
		gojs.Objects{
			// Functions
			"Connect":           lib_redis.Connect,
			"GetServerInfo":     lib_redis.GetServerInfo,
			"GetServerInfoAuth": lib_redis.GetServerInfoAuth,
			"IsAuthenticated":   lib_redis.IsAuthenticated,
			"RunLuaScript":      lib_redis.RunLuaScript,

			// Var and consts

			// Objects / Classes

		},
	).Register()
}

func Enable(runtime *goja.Runtime) {
	module.Enable(runtime)
}
