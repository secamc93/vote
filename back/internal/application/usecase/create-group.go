package usecase

import (
	"voting/internal/domain/dtos"
)

func (uc *UsecaseVote) CreateVoteGroup(group dtos.VoteGroupDTO) (uint, error) {
	groupID, err := uc.voteRepo.CreateVoteGroup(group)
	if err != nil {
		uc.log.Error("error : ", err)
		return 0, err
	}
	return groupID, nil
}
