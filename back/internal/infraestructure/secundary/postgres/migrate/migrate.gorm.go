package migrate

import (
	"voting/internal/infraestructure/secundary/postgres/connectpostgres"
	"voting/internal/infraestructure/secundary/postgres/models"
	"voting/pkg/logger"
)

func Migrate() {
	log := logger.NewLogger()
	db := connectpostgres.New().GetDB()
	err := db.AutoMigrate(
		&models.User{},
		&models.House{},
		&models.VoteGroup{},
		&models.Voting{},
		&models.VoteOption{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}
	log.Info("Database migrated")
}
