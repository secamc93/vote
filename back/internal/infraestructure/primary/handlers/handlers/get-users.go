package handlers

import (
	"net/http"
	"strconv"
	"voting/internal/infraestructure/primary/handlers/handlers/mapper"

	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users associated with a specific group.
// @Summary Retrieve users for a group
// @Description Get a list of users belonging to the specified group using the provided groupID.
// @Tags Users
// @Accept json
// @Produce json
// @Param groupID path int true "Group ID"
// @Success 200 {object} map[string]interface{} "List of users"
// @Failure 400 {object} map[string]string "Bad Request - invalid groupID"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /get-users/{groupID} [get]
func (h *HandlerVote) GetUsers(ctx *gin.Context) {
	groupID, err := strconv.Atoi(ctx.Param("groupID"))
	if err != nil {
		h.log.Error("error : ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users, err := h.usecase.GetUsers(uint(groupID))
	if err != nil {
		h.log.Error("error : ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	usersResponse := mapper.MapUserDTOsToResponses(users)

	ctx.JSON(http.StatusOK, gin.H{"users": usersResponse})
}
