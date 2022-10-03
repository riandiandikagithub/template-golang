package http

import (
	"net/http"
	"payment-simulator/models"
	"payment-simulator/modules/users"

	"github.com/gin-gonic/gin"
)

// UserHandler  represent the httphandler for users
type UserHandler struct {
	UserUsecase users.Usecase
}

// NewUserHandler will initialize the user/ resources endpoint
func NewUserHandler(e *gin.Engine, us users.Usecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	// e.GET("/articles", handler.FetchArticle)
	// e.POST("/articles", handler.Store)
	e.GET("/user/:username", handler.GetByUsername)
	// e.DELETE("/articles/:id", handler.Delete)
}

func (u *UserHandler) GetByUsername(ctx *gin.Context) {
	ctxContext := ctx.Request.Context()

	defer ctxContext.Done()

	username := ctx.Param("username")

	generalResponse := models.Response{

		Rc:      models.ERR_CODE_00,
		Message: models.ERR_CODE_00_MSG,
	}
	users, err := u.UserUsecase.GetByUsername(ctxContext, username)

	if err != nil {
		generalResponse.Rc = models.ERR_CODE_01
		generalResponse.Message = models.ERR_CODE_01_MSG
		ctx.JSON(http.StatusOK, generalResponse)
		ctx.Done()

	}
	generalResponse.Data = users
	ctx.JSON(http.StatusOK, generalResponse)
	ctx.Done()
}
