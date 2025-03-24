package repository

import "voting/internal/domain/dtos"

func (r Repository) GetGroups() ([]dtos.VoteGroupDTO, error) {
	var groups []dtos.VoteGroupDTO
	db := r.dbConnection.
		GetDB().Table("vote_groups vg").
		Select(`
			vg.id as id,
			vg.name as name,
			vg.created_at
		`).
		Scan(&groups)
	if db.Error != nil {
		return nil, db.Error
	}
	return groups, nil
}
