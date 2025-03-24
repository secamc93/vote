package ports

import "voting/internal/domain/dtos"

type IVote interface {
	CreateVoteGroup(group dtos.VoteGroupDTO) (uint, error)
	CreateUser(user dtos.UserDTO) (uint, error)
	GetVoteGroupByID(id uint) (bool, error)
	GetHouseByID(id uint) (bool, error)
	GetHouseAndVoteGroupInUser(houseID, voteGroupID uint) (bool, error)
	CreateHouse(house dtos.HouseDTO) (uint, error)
	GetHouseByName(name string) (bool, error)
	GetUsers(groupID uint) ([]dtos.UserDTO, error)
	GetHouses(houseDTO dtos.HouseDTO) ([]dtos.HouseDTO, error)
	GetGroups() ([]dtos.VoteGroupDTO, error)
}
