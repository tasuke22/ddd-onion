package health_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/go-onion/presentation/settings"
)

func HealthCheck(ctx *gin.Context) {
	res := HealthResponse{
		Status: "ok",
	}
	settings.ReturnStatusOK(ctx, res)
}
