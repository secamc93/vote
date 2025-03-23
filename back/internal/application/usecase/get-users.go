package usecase

import (
	"voting/internal/domain/dtos"
)

func (u *UsecaseVote) GetUsers(groupID uint) ([]dtos.UserDTO, error) {
	users, err := u.voteRepo.GetUsers(groupID)
	if err != nil {
		u.log.Error("error : ", err)
		return nil, err
	}
	return users, nil
}
