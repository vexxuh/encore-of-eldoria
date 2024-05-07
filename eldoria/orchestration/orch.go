package orch

import (
	gamecore "encore.app/eldoria/game-core"
	"gorm.io/gorm"
)

func GenCharacter(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	// this will create a new and update the old character if a player wants to generate a new character
	response, err := createNewCharacter(db, novelApiKey, msg, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}

func GetCharacterSheetAndInventory(db *gorm.DB, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	response, err := characterSheetAndInvetoryCRUD(db, msg, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}

func TravelToLocation(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	response, err := travelToLocation(db, novelApiKey, msg, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}

func SearchForFight(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	response, err := searchForFight(db, novelApiKey, msg, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}

func AttackEnemy(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	response, err := attackEnemy(db, novelApiKey, msg, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}

func PlayerStore(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	response, err := playerStore(db, novelApiKey, msg, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}

func CheatMode(db *gorm.DB, playerId string) (gamecore.DiscordMessageResponse, error) {
	response, err := cheatMode(db, playerId)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}
	return response, nil
}
