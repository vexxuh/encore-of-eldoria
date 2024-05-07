package orch

import (
	gamecore "encore.app/eldoria/game-core"
	novel_ai "encore.app/eldoria/game-core/ai/novel-ai"
	models "encore.app/eldoria/game-core/data"
	"gorm.io/gorm"
	"log"
)

func characterSheetAndInvetoryCRUD(db *gorm.DB, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)
	if c == nil {
		panic("AHHHHHHHH")
	}

	response := buildNoAiResponse(c)

	db.Save(&c)

	return response, nil
}

func searchForFight(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)
	// first parm is for logging
	_, prompt, err := gamecore.Combat(c, msg)
	if err != nil {
		panic("AHHHHHHH")
	}

	// do AI action
	aiData, err := novel_ai.GenerateAiText(novelApiKey, prompt)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}

	response := buildResponse(aiData, c)

	db.Save(&c)

	return response, nil
}

func attackEnemy(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)

	_, prompt, err := gamecore.Combat(c, msg)
	if err != nil {
		panic("AHHHHHHH")
	}
	// do AI action
	aiData, err := novel_ai.GenerateAiText(novelApiKey, prompt)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}

	response := buildResponse(aiData, c)

	db.Save(&c)

	return response, nil
}

func playerStore(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)

	_, prompt, err := gamecore.Store(c, msg)
	if err != nil {
		panic("AHHHHHHH")
	}

	// do AI action
	aiData, err := novel_ai.GenerateAiText(novelApiKey, prompt)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}

	response := buildResponse(aiData, c)

	db.Save(&c)

	return response, nil
}

func cheatMode(db *gorm.DB, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)

	// do game logic stuff here
	_, _, err := gamecore.CheatMode(c)
	if err != nil {
		panic("AHHHHHHH")
	}

	response := buildNoAiResponse(c)

	db.Save(&c)

	return response, nil
}

func travelToLocation(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)

	_, prompt, err := gamecore.MoveArea(c, msg)
	if err != nil {
		panic("AHHHH")
	}

	// do AI action
	aiData, err := novel_ai.GenerateAiText(novelApiKey, prompt)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}

	response := buildResponse(aiData, c)

	db.Save(&c)

	return response, nil
}

func createNewCharacter(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {
	c := findCharacter(db, playerId)

	_, prompt, char, err := gamecore.CreatePlayer(c, playerId, msg)
	if err != nil {
		panic("AHHHHH")
	}

	aiData, err := novel_ai.GenerateAiText(novelApiKey, prompt)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}

	response := buildResponse(aiData, &char)

	// do database action this will update and save the new character
	if err := db.Save(&char).Error; err != nil {
		// Handle error
		log.Println("Error saving data entry:", err)
	}

	return response, nil
}

func findCharacter(db *gorm.DB, playerId string) *models.Character {
	var c models.Character
	if err := db.First(&c, playerId).Error; err != nil {
		return nil
	}

	return &c
}

func buildResponse(aiData string, char *models.Character) gamecore.DiscordMessageResponse {
	player := convertCharacterToResponse(char)
	return gamecore.DiscordMessageResponse{
		P:       player,
		TextGen: aiData,
	}
}

func buildNoAiResponse(c *models.Character) gamecore.DiscordMessageResponse {
	player := convertCharacterToResponse(c)
	return gamecore.DiscordMessageResponse{
		P:       player,
		TextGen: "",
	}
}

func convertCharacterToResponse(char *models.Character) gamecore.Player {
	return gamecore.Player{
		Inventory: gamecore.PlayerInventoryResponse{
			I_apple:      char.Inventory.IApple,
			I_potion:     char.Inventory.IPotion,
			I_potionPlus: char.Inventory.IPotionplus,
			C_gold:       char.Inventory.CGold,
			B_gold:       char.Inventory.IApple,
		},
		User:           char.User,
		C_level:        char.CLevel,
		C_health:       char.CHealth,
		M_health:       char.MHealth,
		B_health:       char.BHealth,
		S_strength:     char.SStrength,
		S_agility:      char.SAgility,
		S_constitution: char.SConstitution,
		S_intelligence: char.SIntelligence,
		S_wisdom:       char.SWisdom,
		W_s_sword:      char.WSSword,
		W_s_axe:        char.WSAxe,
		W_s_spear:      char.WSSpear,
		P_state:        char.PState,
		C_area:         char.CArea,
	}
}

func getPlayerSession(playerId string) (*models.Character, error) {

	return nil, nil
}
