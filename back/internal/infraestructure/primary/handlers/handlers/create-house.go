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

// CreateHouse godoc
// @Summary Create a new house
// @Description Creates a new house with the provided house data and returns the generated house ID.
// @Tags House
// @Accept json
// @Produce json
// @Param house body request.HouseRequest true "House data"
// @Success 201 {object} map[string]interface{} "House created successfully, returns house_id"
// @Failure 400 {object} response.BaseResponse "Bad request or house already exists"
// @Failure 500 {object} response.BaseResponse "Internal server error"
// @Router /create-house [post]
func (h *HandlerVote) CreateHouse(c *gin.Context) {
	var HouseRequest request.HouseRequest
	if err := c.ShouldBindJSON(&HouseRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}
	house := mapper.MapHouseRequestToDTO(HouseRequest)
	houseID, err := h.usecase.CreateHouse(house)
	if err != nil {
		// Manejo actualizado de errores
		if errors.Is(err, domainerrors.ErrHouseFound) ||
			errors.Is(err, domainerrors.ErrHouseNameIsRequired) ||
			errors.Is(err, domainerrors.ErrVoteGroupIDIsRequired) {
			c.JSON(http.StatusBadRequest, response.BaseResponse{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message":  "House created successfully",
		"house_id": houseID,
	})
}
