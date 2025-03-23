package mappers

import (
	"voting/internal/domain/dtos"
	"voting/internal/infraestructure/secundary/postgres/models"
)

func MapUserModelsToDTOs(users []models.User) []dtos.UserDTO {
	dtoUsers := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		dtoUsers[i] = dtos.UserDTO{
			ID:          user.ID,
			Name:        user.Name,
			Dni:         user.Dni,
			HouseID:     user.HouseID,
			VoteGroupID: user.VoteGroupID,
		}
	}
	return dtoUsers
}
