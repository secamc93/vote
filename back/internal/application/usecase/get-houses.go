package usecase

import "voting/internal/domain/dtos"

func (u *UsecaseVote) GetHouses(houseDTO dtos.HouseDTO) ([]dtos.HouseDTO, error) {
	houses, err := u.voteRepo.GetHouses(houseDTO)
	if err != nil {
		u.log.Error("error : ", err)
		return nil, err
	}
	return houses, nil
}
