package repository

import (
	"errors"
	"sync"
	"voting/internal/domain/dtos"
	"voting/internal/infraestructure/secundary/postgres/connectpostgres"
	"voting/internal/infraestructure/secundary/postgres/mappers"
	"voting/internal/infraestructure/secundary/postgres/models"
	"voting/pkg/logger"

	"gorm.io/gorm" // import agregado
)

type Repository struct {
	dbConnection connectpostgres.DBConnection
	log          logger.ILogger
}

var (
	instance *Repository
	once     sync.Once
)

func New(db connectpostgres.DBConnection) *Repository {
	once.Do(func() {
		instance = &Repository{
			dbConnection: db,
			log:          logger.NewLogger(),
		}
	})
	return instance
}
func (r *Repository) CreateVoteGroup(group dtos.VoteGroupDTO) (uint, error) {
	voteGroup := models.VoteGroup{
		Name: group.Name,
	}
	result := r.dbConnection.GetDB().Create(&voteGroup)
	return voteGroup.ID, result.Error
}

func (r *Repository) CreateUser(user dtos.UserDTO) (uint, error) {
	userModel := models.User{
		Name:        user.Name,
		Dni:         user.Dni,
		HouseID:     user.HouseID,     // asignación agregada
		VoteGroupID: user.VoteGroupID, // asignación agregada
	}
	result := r.dbConnection.GetDB().Create(&userModel)
	return userModel.ID, result.Error
}

func (r *Repository) GetUsers(groupID uint) ([]dtos.UserDTO, error) {
	var usersModel []models.User
	result := r.dbConnection.GetDB().Where("vote_group_id = ?", groupID).Find(&usersModel)

	userDto := mappers.MapUserModelsToDTOs(usersModel)
	return userDto, result.Error
}

func (r *Repository) GetVoteGroupByID(id uint) (bool, error) {
	var count int64
	result := r.dbConnection.GetDB().Model(&models.VoteGroup{}).Where("id = ?", id).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r *Repository) GetHouseByID(id uint) (bool, error) {
	var count int64
	result := r.dbConnection.GetDB().Model(&models.House{}).Where("id = ?", id).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count != 0, nil
}
func (r *Repository) GetHouseByName(name string) (bool, error) {
	var house models.House
	result := r.dbConnection.GetDB().Where("name = ?", name).First(&house)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (r *Repository) GetHouseAndVoteGroupInUser(houseID, voteGroupID uint) (bool, error) {
	var count int64
	result := r.dbConnection.GetDB().Model(&models.User{}).Where("house_id = ? AND vote_group_id = ?", houseID, voteGroupID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (r Repository) CreateHouse(house dtos.HouseDTO) (uint, error) {
	houseModel := models.House{
		Name: house.Name,
	}
	result := r.dbConnection.GetDB().Create(&houseModel)
	if result.Error != nil {
		return 0, result.Error
	}
	// Solo se crea la relación many-to-many con VoteGroup si se pasa VoteGroupID distinto de 0
	if house.VoteGroupID != 0 {
		voteGroup := models.VoteGroup{Model: gorm.Model{ID: house.VoteGroupID}}
		if err := r.dbConnection.GetDB().
			Model(&houseModel).
			Association("VoteGroups").
			Append(&voteGroup); err != nil {
			return 0, err
		}
	}
	return houseModel.ID, nil
}
