package mapper

import (
	"voting/internal/domain/dtos"
	"voting/internal/infraestructure/primary/handlers/handlers/dtos/request"
	"voting/internal/infraestructure/primary/handlers/handlers/dtos/response"
)

func MapVoteGroupRequestToDTO(req request.VoteGroupRequest) dtos.VoteGroupDTO {
	return dtos.VoteGroupDTO{
		Name: req.Name,
	}
}

// Nueva función para mapear UserRequest a UserDTO.
func MapUserRequestToDTO(req request.UserRequest) dtos.UserDTO {
	return dtos.UserDTO{
		Name:        req.Name,
		Dni:         req.Dni,
		HouseID:     req.HouseID,
		VoteGroupID: req.VoteGroupID,
	}
}

func MapHouseRequestToDTO(req request.HouseRequest) dtos.HouseDTO {
	return dtos.HouseDTO{
		Name:        req.Name,
		VoteGroupID: req.VoteGroupID,
	}
}

func MapUserDTOsToResponses(users []dtos.UserDTO) []response.UserResponse {
	responses := make([]response.UserResponse, len(users))
	for i, user := range users {
		responses[i] = response.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Dni:         user.Dni,
			HouseID:     user.HouseID,
			VoteGroupID: user.VoteGroupID,
		}
	}
	return responses
}
