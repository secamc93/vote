package dtos

type UserDTO struct {
	ID          uint
	Name        string
	Dni         string
	HouseID     uint
	House       HouseDTO
	VoteGroupID uint
	VoteGroup   VoteGroupDTO
}

type HouseDTO struct {
	ID          uint
	Name        string
	VoteGroupID uint
}

type VoteGroupDTO struct {
	ID     uint
	Name   string
	Houses []HouseDTO
}

type VotingDTO struct {
	ID          uint
	VoteGroupID uint
	Name        string
	VoteOptions []VoteOptionDTO
}

type VoteOptionDTO struct {
	ID       uint
	VotingID uint
	Name     string
	Vote     bool
}
