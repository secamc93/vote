package repository

import (
	"voting/internal/domain/dtos"
)

func (r *Repository) GetHouses(houseDTO dtos.HouseDTO) ([]dtos.HouseDTO, error) {
	var houses []dtos.HouseDTO
	db := r.dbConnection.
		GetDB().Table("houses h").
		Select(`
			h.id as id,
			h.name as name,
			h.created_at
		`).
		Scan(&houses)
	if db.Error != nil {
		return nil, db.Error
	}
	r.log.Info("houses : ", houses)
	return houses, nil
}
