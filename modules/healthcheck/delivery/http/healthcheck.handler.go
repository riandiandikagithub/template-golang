package http

import (
	"context"
	"net/http"
	"payment-simulator/models"

	"github.com/gin-gonic/gin"
)

// HealthcheckHandler  represent the httphandler for Healthcheck
type HealthcheckHandler struct {
	// HealthcheckUsecase healthcheck.Usecase
}

// NewHealthcheckHandler will initialize the Healthcheck / resources endpoint
func NewHealthcheckHandler(e *gin.Engine) {
	handler := &HealthcheckHandler{
		// HealthcheckUsecase: us,
	}
	e.GET("/healthcheck", handler.CheckService)
}

func (structHCHandler *HealthcheckHandler) CheckService(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()
	//hc := healthcheck.HealthCheck{
	//	Database: postgres.CheckStatusDb(),
	//	//Redis:    redis.CheckStatus(),
	//}

	res := models.Response{
		Rc:      "00",
		Message: "Success",
		Data:    "hc",
	}
	ctx.JSON(http.StatusOK, res)
	ctx.Done()
}
