package handlers

import (
	"net/http"
	"voting/internal/domain/dtos"
	"voting/internal/infraestructure/primary/handlers/handlers/mapper"

	"github.com/gin-gonic/gin"
)

// GetHouses retrieves a list of houses based on the provided HouseDTO filter from the query parameters.
//
// @Summary Retrieve list of houses
// @Description Retrieves houses using the provided filter criteria in the query parameters.
// @Tags House
// @Accept json
// @Produce json
// @Success 200 {array} response.HouseResponse "List of houses"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /get-houses [get]
func (h *HandlerVote) GetHouses(ctx *gin.Context) {
	var houseDTO dtos.HouseDTO
	if err := ctx.ShouldBindQuery(&houseDTO); err != nil {
		h.log.Error("error : ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	houses, err := h.usecase.GetHouses(houseDTO)
	if err != nil {
		h.log.Error("error : ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	housesResponse := mapper.MapHouseDTOsToResponses(houses)

	ctx.JSON(http.StatusOK, housesResponse)
}
