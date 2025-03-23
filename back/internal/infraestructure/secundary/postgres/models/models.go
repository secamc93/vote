package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string    `gorm:"column:name"`
	Dni         string    `gorm:"column:dni"`
	HouseID     uint      `gorm:"column:house_id"`
	House       House     `gorm:"foreignKey:HouseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VoteGroupID uint      `gorm:"column:vote_group_id"`
	VoteGroup   VoteGroup `gorm:"foreignKey:VoteGroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type House struct {
	gorm.Model
	Name       string      `gorm:"column:name"`
	VoteGroups []VoteGroup `gorm:"many2many:vote_group_houses;"` // relación many2many con VoteGroup
}

type VoteGroup struct {
	gorm.Model
	Name   string  `gorm:"column:name"`
	Houses []House `gorm:"many2many:vote_group_houses;"`
}

type Voting struct {
	gorm.Model
	VoteGroupID uint         `gorm:"column:vote_group_id"`
	Name        string       `gorm:"column:name"`
	VoteOptions []VoteOption `gorm:"foreignKey:VotingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Actualizado
}

type VoteOption struct {
	gorm.Model
	VotingID uint   // Clave foránea asignada
	Name     string `gorm:"column:name"`
	Vote     bool   `gorm:"column:is_winner"`
}
