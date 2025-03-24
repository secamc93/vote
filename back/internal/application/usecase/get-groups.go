package usecase

import "voting/internal/domain/dtos"

func (u *UsecaseVote) GetGroups() ([]dtos.VoteGroupDTO, error) {
	groups, err := u.voteRepo.GetGroups()
	if err != nil {
		u.log.Error("error : ", err)
		return nil, err
	}
	return groups, nil
}
