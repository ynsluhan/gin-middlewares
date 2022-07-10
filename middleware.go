package gin_middlewares

import (
	"github.com/gin-gonic/gin"
	config "github.com/ynsluhan/go-config"
	"github.com/ynsluhan/go-gin-logger"
	handler "github.com/ynsluhan/gin-error-handler"
	jwt "github.com/ynsluhan/gin-jwt-middleware"
	rbdc "github.com/ynsluhan/go-gin-rbac-middleware"
	cross "github.com/ynsluhan/go-cross"
)

var conf *config.Config

func init() {
	conf = config.GetConf()
}

func Middleware(engine *gin.Engine) {
	// 开启防止崩溃
	if conf.Server.EnableRecover {
		engine.Use(
			gin.Recovery(),
		)
	}
	// 跨域
	engine.Use(
		cross.Cross(),
	)
	// 开启jwt
	if conf.Server.EnableJwt {
		engine.Use(
			jwt.Jwt(),
		)
	}
	// 开启rbac权限认证
	if conf.Server.EnableRbac {
		engine.Use(
			rbdc.RBACMiddle(),
		)
	}
	// 开启日志
	if conf.Server.EnableLogger {
		engine.Use(
			Logger.DefaultLogger(),
		)
	}
	// 开启异常处理
	if conf.Server.EnableErrorRecover {
		engine.Use(
			handler.ErrorRecover(),
		)
	}
}
