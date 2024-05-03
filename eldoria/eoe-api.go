package eoe_api

import (
	"context"
	"fmt"

	gamecore "encore.app/eldoria/game-core"
	orch "encore.app/eldoria/orchestration"
	"encore.dev/config"
	"encore.dev/storage/sqldb"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//encore:service
type Service struct {
	db *gorm.DB
}

type ApiConfig struct {
	NovelApiKey string
}

var secrets struct {
	NovelApiKey string
}

var cfg *ApiConfig = config.Load[*ApiConfig]()

var blogDB = sqldb.NewDatabase("game_db", sqldb.DatabaseConfig{
	Migrations: "./databases/migrations",
})

// initService initializes the site service(s).
func initService() (*Service, error) {
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
func (s *Service) CreateCharacter(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	response, err := orch.GenCharacter(s.db, secrets.NovelApiKey, params.Msg, params.PlayerId)
	print := fmt.Sprintf("User: %s, C_level: %d, C_health: %d", response.P.User, response.P.C_level, response.P.C_health)
	fmt.Println(print)
	if err != nil {
		fmt.Printf("msg: %vn", response)
	}
	return &response, nil
}
