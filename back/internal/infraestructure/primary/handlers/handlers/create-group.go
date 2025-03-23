package handlers

import (
	"net/http"
	"voting/internal/infraestructure/primary/handlers/handlers/dtos/request"
	"voting/internal/infraestructure/primary/handlers/handlers/dtos/response" // agregado
	"voting/internal/infraestructure/primary/handlers/handlers/mapper"

	"github.com/gin-gonic/gin"
)

// CreateVoteGroup godoc
// @Summary      Crea un nuevo grupo de votación
// @Description  Crea un grupo de votación utilizando los datos proporcionados en el cuerpo de la petición.
// @Tags         VoteGroup
// @Accept       json
// @Produce      json
// @Param        group  body      request.VoteGroupRequest  true  "Detalles del grupo de votación"
// @Success      200    {object}  map[string]string           "status: ok"
// @Failure      400    {object}  map[string]string           "error description"
// @Failure      500    {object}  map[string]string           "error description"
// @Router       /vote-group [post]
func (h *HandlerVote) CreateVoteGroup(ctx *gin.Context) {
	var groupRequest request.VoteGroupRequest

	if err := ctx.ShouldBindJSON(&groupRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}
	group := mapper.MapVoteGroupRequestToDTO(groupRequest)
	groupID, err := h.usecase.CreateVoteGroup(group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Creación Exitosa",
		"group_id": groupID,
	})
}
