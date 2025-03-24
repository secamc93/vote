package handlers

import (
	"net/http"
	"voting/internal/infraestructure/primary/handlers/handlers/mapper"

	"github.com/gin-gonic/gin"
)

// GetGroups godoc
// @Summary Retrieve a list of groups
// @Description Get all groups for voting
// @Tags VoteGroup
// @Accept json
// @Produce json
// @Success 200 {array}  response.VoteGroupResponse "List of groups"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /get-groups [get]
func (h *HandlerVote) GetGroups(ctx *gin.Context) {
	groups, err := h.usecase.GetGroups()
	if err != nil {
		h.log.Error("error : ", err)
		return
	}

	groupsResponses := mapper.MapVoteGroupDTOToResponses(groups)

	ctx.JSON(http.StatusOK, groupsResponses)
}
