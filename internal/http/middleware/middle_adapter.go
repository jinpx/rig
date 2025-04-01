package middleware

import "github.com/gin-gonic/gin"

func AdaptCheck(ctx *gin.Context) {
	_ = AdaptUnityCheck(ctx)
	ctx.Next()
}

func AdaptUnityCheck(ctx *gin.Context) error {
	device := ctx.GetHeader(constant.Device)
	if device == "" {
		//return HandleMidErr(ctx, nil, dto.ErrorModel(data.DeviceEmptyErr))
		ctx.Set("device", 2)
		return nil
	}
	ctx.Set("device", device)
	return nil
}
