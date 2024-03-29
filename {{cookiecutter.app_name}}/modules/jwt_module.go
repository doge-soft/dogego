package modules

import (
	"{{cookiecutter.app_name}}/cache"
	"github.com/doge-soft/dogego_module_jwt/jwt"
)

var JWTModule *jwt.RedisJWT

func InitJWTModule() {
	JWTModule = jwt.NewRedisJWT(cache.CacheClient)
}
