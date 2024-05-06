package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Inventory      Inventory `gorm:"foreignKey:IventoryId"`
	InventoryId	   int
	Creature	   Creature `gorm:"foreignKey:CreatureId"`
	CreatureId	   int
	Username       string
	User           string
	C_level        int
	C_experience   int
	C_health       int
	M_health       int
	B_health       int
	S_strength     int
	S_agility      int
	S_constitution int
	S_intelligence int
	S_wisdom       int
	W_s_sword      int
	W_e_sword      int
	W_s_axe        int
	W_e_axe        int
	W_s_spear      int
	W_e_spear      int
	P_state        string
	C_area         string
	C_e_weapon     int
	C_e_armor      int
}

type Inventory struct {
	gorm.Model
	I_apple      int
	I_potion     int
	I_potionPlus int
	C_gold       int
	B_gold       int
}

type Creature struct {
	gorm.Model
	C_name       string
	C_id     	 int
	C_level      int
	C_experience int
	C_c_health	 int
	C_m_health   int
	C_attack     int
	C_defense    int

}