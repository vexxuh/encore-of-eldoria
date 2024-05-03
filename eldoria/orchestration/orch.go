package orch

import (
	gamecore "encore.app/eldoria/game-core"
	novel_ai "encore.app/eldoria/game-core/ai/novel-ai"
	models "encore.app/eldoria/game-core/data"
	"gorm.io/gorm"
	"fmt"
)

func GenCharacter(db *gorm.DB, novelApiKey string, msg string, playerId string) (gamecore.DiscordMessageResponse, error) {

	// do game logic jaction
	_, prompt, char := gamecore.CreatePlayer(playerId)
	print := fmt.Sprintf("User: %s, C_level: %d, C_health: %d", char.User, char.C_level, char.C_health)
	fmt.Println(print)

	// do AI action
	aiData, err := novel_ai.GenerateAiText(novelApiKey, prompt)
	if err != nil {
		return gamecore.DiscordMessageResponse{}, err
	}

	response := buildResponse(aiData, &char)
	print = fmt.Sprintf("User: %s, C_level: %d, C_health: %d", response.P.User, response.P.C_level, response.P.C_health)
	fmt.Println(print)

	// do database action
	db.Save(&char)

	return response, nil
}

func buildResponse(aiData string, char *models.Character) gamecore.DiscordMessageResponse {
	player := convertCharacterToResponse(char)
	return gamecore.DiscordMessageResponse{
		P:       player,
		TextGen: aiData,
	}
}

func convertCharacterToResponse(char *models.Character) gamecore.Player {
	return gamecore.Player{
		Inventory: gamecore.PlayerInventoryResponse{
			I_apple:      char.Inventory.I_apple,
			I_potion:     char.Inventory.I_potion,
			I_potionPlus: char.Inventory.I_potionPlus,
			C_gold:       char.Inventory.C_gold,
			B_gold:       char.Inventory.I_apple,
		},
		User:           char.User,
		C_level:        char.C_level,
		C_health:       char.C_health,
		M_health:       char.M_health,
		B_health:       char.B_health,
		S_strength:     char.S_strength,
		S_agility:      char.S_agility,
		S_constitution: char.S_constitution,
		S_intelligence: char.S_intelligence,
		S_wisdom:       char.S_wisdom,
		W_s_sword:      char.W_s_sword,
		W_s_axe:        char.W_s_axe,
		W_s_spear:      char.W_s_spear,
		P_state:        char.P_state,
		C_area:         char.C_area,
	}
}

func getPlayerSession(playerId string) (*models.Character, error) {

	return nil, nil
}
