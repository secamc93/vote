package handlers

import (
	"voting/internal/application/usecase"
	"voting/pkg/logger"

	"github.com/gin-gonic/gin"
)

type IHandlerVote interface {
	CreateVoteGroup(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	CreateHouse(c *gin.Context)
	GetUsers(ctx *gin.Context)
	GetHouses(ctx *gin.Context)
	GetGroups(ctx *gin.Context)
}

type HandlerVote struct {
	usecase usecase.IUsecaseVote
	log     logger.ILogger
}

func New(usecase usecase.IUsecaseVote) IHandlerVote {
	return &HandlerVote{
		usecase: usecase,
		log:     logger.NewLogger(),
	}
}
