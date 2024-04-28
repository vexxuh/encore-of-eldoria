package eoe_api

import (
	"context"
	"fmt"

	orch "encore.app/orchestration"
	"encore.dev/storage/sqldb"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

var blogDB = sqldb.NewDatabase("game_db", sqldb.DatabaseConfig{
	Migrations: "./databases/migrations",
})

// initService initializes the site service(s).
func initService() (*Service, error) {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// load database service
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: blogDB.Stdlib(),
	}))
	if err != nil {
		return nil, err
	}
	return &Service{db: db}, nil
}

// Attack defines a testing endpoint for connection testing
//
// encore:api public
func (s *Service) Attack(ctx context.Context, params *AttackParam) (*AttackResponse, error) {
	msg, err := orch.Attack(s.db, params.Type)
	if err != nil {
		return &AttackResponse{Message: fmt.Sprintf("Hmm seem like %s is not a valid attack... please try this again "+err.Error(), params.Type)}, nil
	}
	return &AttackResponse{Message: msg}, nil
}
