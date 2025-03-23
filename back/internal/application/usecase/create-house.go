package usecase

import (
	"voting/internal/domain/domainerrors"
	"voting/internal/domain/dtos"
)

func (u *UsecaseVote) CreateHouse(house dtos.HouseDTO) (uint, error) {

	if house.Name == "" {
		return 0, domainerrors.ErrHouseNameIsRequired
	}
	if house.VoteGroupID == 0 {
		return 0, domainerrors.ErrVoteGroupIDIsRequired
	}

	existHouse, err := u.voteRepo.GetHouseByName(house.Name)
	if err != nil {
		u.log.Error("error : ", err)
		return 0, err
	}
	if existHouse {
		return 0, domainerrors.ErrHouseFound
	}

	houseID, err := u.voteRepo.CreateHouse(house)
	if err != nil {
		u.log.Error("error : ", err)
		return 0, err
	}
	return houseID, nil
}
