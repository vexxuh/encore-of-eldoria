package orch

import (
	"errors"
	"fmt"

	novel_ai "encore.app/game-core/ai/novel-ai"
	"encore.app/game-core/data/models"
	"gorm.io/gorm"
)

func Attack(db *gorm.DB, novelApiKey string, attack_type string) (string, error) {
	if attack_type == "" {
		return "", errors.New("Not a valid attack type")
	}

	// do game logic action

	// do AI action
	aiData, err := novel_ai.GenerateAiText(novelApiKey, "He started to attack me with a "+attack_type)
	if err != nil {
		return "", err
	}

	// do database action
	attack := models.Attacks{AttackType: attack_type}
	db.Save(&attack)

	return fmt.Sprintf("Ow that %s hurt! "+aiData, attack_type), nil
}
