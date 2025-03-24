package usecase

import (
	"voting/internal/domain/dtos"
	"voting/internal/domain/ports"
	"voting/pkg/logger"
)

type IUsecaseVote interface {
	CreateVoteGroup(group dtos.VoteGroupDTO) (uint, error)
	CreateUser(user dtos.UserDTO) (uint, error)
	CreateHouse(house dtos.HouseDTO) (uint, error)
	GetUsers(groupID uint) ([]dtos.UserDTO, error)
	GetHouses(houseDTO dtos.HouseDTO) ([]dtos.HouseDTO, error)
	GetGroups() ([]dtos.VoteGroupDTO, error)
}

type UsecaseVote struct {
	voteRepo ports.IVote
	log      logger.ILogger
}

func New(voteRepo ports.IVote) IUsecaseVote {
	return &UsecaseVote{
		voteRepo: voteRepo,
		log:      logger.NewLogger(),
	}
}
