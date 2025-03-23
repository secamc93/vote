package handlers

import (
	"errors" // importación estándar para errors.Is
	"net/http"
	"voting/internal/domain/domainerrors"

	"voting/internal/infraestructure/primary/handlers/handlers/dtos/request"
	"voting/internal/infraestructure/primary/handlers/handlers/dtos/response"
	"voting/internal/infraestructure/primary/handlers/handlers/mapper"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Parses a JSON payload to create a new user and returns the created user's ID.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body request.UserRequest true "User Data"
// @Success 200 {object} map[string]string           "status: ok"
// @Failure 400 {object} response.BaseResponse "Bad request due to invalid JSON payload"
// @Failure 500 {object} response.BaseResponse "Internal server error during user creation"
// @Router /create-user [post]
func (h *HandlerVote) CreateUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}
	user := mapper.MapUserRequestToDTO(userRequest)
	userID, err := h.usecase.CreateUser(user)
	if err != nil {
		if errors.Is(err, domainerrors.ErrVoteGroupNotFound) ||
			errors.Is(err, domainerrors.ErrHouseNotFound) ||
			errors.Is(err, domainerrors.ErrHouseAndVoteGroupInUserExist) {
			ctx.JSON(http.StatusBadRequest, response.BaseResponse{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Creación Exitosa",
		"user_id": userID,
	})
}
