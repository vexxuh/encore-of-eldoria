package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Inventory Inventory `gorm:"foreignKey:CharacterId"`

	Creature Creature `gorm:"foreignKey:CharacterId"`

	Weapon Weapon `gorm:"foreignKey:CharacterId"`

	Armor Armor `gorm:"foreignKey:CharacterId"`

	Username      string
	User          string
	CLevel        int
	CExperience   int
	CHealth       int
	MHealth       int
	BHealth       int
	SStrength     int
	SAgility      int
	SConstitution int
	SIntelligence int
	SWisdom       int
	WSMelee       int
	WEMelee       int
	WSSword       int
	WESword       int
	WSAxe         int
	WEAxe         int
	WSSpear       int
	WESpear       int
	PState        string
	CArea         string
	CEWeapon      int
	CEArmor       int
}

type Inventory struct {
	gorm.Model
	CharacterId uint
	IApple      int
	IPotion     int
	IPotionplus int
	CGold       int
	BGold       int
}

type Creature struct {
	gorm.Model
	CharacterId int
	CName       string
	CId         int
	CLevel      int
	CExperience int
	CCHealth    int
	CMHealth    int
	CAttack     int
	CDefense    int
}

type Weapon struct {
	gorm.Model
	CharacterId   int
	IName         string
	IId           int
	IAttack       int
	IStrength     int
	IDefense      int
	IAgility      int
	IConstitution int
	IType         string
}

type Armor struct {
	gorm.Model
	CharacterId   int
	IName         string
	IId           int
	IAttack       int
	IStrength     int
	IDefense      int
	IAgility      int
	IConstitution int
	IType         string
}
