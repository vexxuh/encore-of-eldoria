package orch

import (
	"errors"
	"fmt"

	"encore.app/game-core/data/models"
	"gorm.io/gorm"
)

func Attack(db *gorm.DB, attack_type string) (string, error) {
	if attack_type == "" {
		return "", errors.New("Not a valid attack type")
	}

	attack := models.Attacks{AttackType: attack_type}
	db.Save(&attack)

	return fmt.Sprintf("Ow that %s hurt!", attack_type), nil
}
