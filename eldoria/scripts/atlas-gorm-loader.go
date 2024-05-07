package main

import (
	"fmt"
	"io"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	models "encore.app/eldoria/game-core/data"
)

// Define the models to generate migrations for.
var ms = []any{
	&models.Character{},
	&models.Inventory{},
	&models.Armor{},
	&models.Weapon{},
	&models.Creature{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(ms...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
