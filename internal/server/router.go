package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.casinovip.tech/peso_lucky_backend/pl_struct/constant"
)

// 游戏路由
func RouterGame(route *gin.Engine) {

}

// 商户路由
func RouterMerchant(route *gin.Engine) {

}

// 连接
func RouterConn(s *gin.Server) {

}

// 活动路由
func RouterAct(route *gin.Server) {
	var (
		api = route.Group(constant.RouteAPI).Group(constant.RouteAct)
	)

	api.Use(middleware.AdaptCheck)
	api.Use(middleware.AuthRequired()) //JWT验证
	api.Use(middleware.LockCheck)
}
