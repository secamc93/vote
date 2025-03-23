package usecase

import (
	"voting/internal/domain/domainerrors"
	"voting/internal/domain/dtos"
)

func (u *UsecaseVote) CreateUser(user dtos.UserDTO) (uint, error) {

	groupExist, err := u.voteRepo.GetVoteGroupByID(user.VoteGroupID)
	if err != nil {
		u.log.Error("error : ", err)
		return 0, err
	}
	if !groupExist {
		return 0, domainerrors.ErrVoteGroupNotFound
	}

	houseExist, err := u.voteRepo.GetHouseByID(user.HouseID)
	if err != nil {
		u.log.Error("error : ", err)
		return 0, err
	}
	if !houseExist {
		return 0, domainerrors.ErrHouseNotFound
	}

	relationExist, err := u.voteRepo.GetHouseAndVoteGroupInUser(user.HouseID, user.VoteGroupID)
	if err != nil {
		u.log.Error("error retrieving relation: ", err)
		return 0, err
	}
	if relationExist {
		return 0, domainerrors.ErrHouseAndVoteGroupInUserExist
	}

	userID, err := u.voteRepo.CreateUser(user)
	if err != nil {
		u.log.Error("error : ", err)
		return 0, err
	}
	return userID, nil
}
