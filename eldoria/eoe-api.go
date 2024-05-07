package eoe_api

import (
	"context"
	gamecore "encore.app/eldoria/game-core"
	orch "encore.app/eldoria/orchestration"
	"encore.dev/beta/errs"
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
	Migrations: "./migrations",
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

// CreateCharacter Command: Create - Create player character. Requires: character name. Return: Flavor text.
//   - Description: "Creates a new character with the given name. Replaces old character if one already exists with your Discord account."
//   - If no character attatched to userID, add character to the DB with that userID as key and character as value
//   - If a character is attatched to userID, replace character stored with userID key
//   - return extra data stating that your previous character has been removed
//   - Uses player name for flavor text generation
//   - Returns: flavor text
//
// - Future
//   - add option to pick a job (ie: Barbarian, Mage, etc.) which gives you different starting stat values and starting weapon and possible some consumables. (think souls games)
//   - add character cosmetics, either AI text describing the character and/or returning a generated image.
//
// encore:api public
func (s *Service) CreateCharacter(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.GenCharacter(s.db, secrets.NovelApiKey, params.Msg, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}

// Inventory Command: Inventory - Show the player Inventory including items, stats, weapons, armor, name, xp. Can be done in any player state. Requires: none. Optional: Action, Optional: Item. Return: formatted text of player's inventory.
// - Description: "Can show your name, stats, xp, weapons, armor, and items. Can also be used to equip items like weapons and armor. Can also be used to inspect gear."
// - Optional: Equip - Equips item in the item argument OR Use - Uses the item in the item argument.
// - Optional: Item - The item to be equiped.
//   - If called with Equip and Item, the item will be equipped if possible (weapons and armor) and stores in DB. Skips the player's action if in combat.
//   - If called with Equip and no Item, all currently equipped items will be unequipped and stores in DB. Skips the player's action if in combat.
//   - If called with Use and Item, the item will be used if possible (potions, etc.) and stores in DB. Skips the player's action if in combat.
//   - If called with Use and no item, nothing happens. Skips the player's action if in combat.
//   - If called with Item and no Equip or Item, returns flavor text of that item. (inspect) Skips the player's action if in combat.
//
// - No argument version can be called in Normal and Combat state and Dead state.
// - All argument version can only be called in Normal and Combat state.
// - Returns: text formatted in a list of the players inventory.
// - Future
//   - Return a generated image created from the item names and quantity as well as the area the player is currently in. (AI Vibe)
//   - If opened while dead, states the enemy that the player was killed by.
//
// encore:api public
func (s *Service) Inventory(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.GetCharacterSheetAndInventory(s.db, params.Msg, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}

// Travel Command: Travel - Take the player to the location they specify. Requires: A choice from the discord bot's list of places. Returns: Flavor text
// - Description: "Moves you to your desired location. Choose from the list Discord provides."
// - Can be called in Normal Player state.
// - Changes player location variable and stores it in the DB.
// - Uses new location (and possibly old location) for text generation
// - Returns: flavor text.
// - Future
//   - Return image generation.
//
// encore:api public
func (s *Service) Travel(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.TravelToLocation(s.db, secrets.NovelApiKey, params.Msg, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}

// Search Command: Search - Puts player into a fight with a static enemy depending on current location. Requires: none. Returns: flavor text of the enemy you found.
// - Description: "Searches the area for enemies to fight and loot to find."
// - Can be called in Normal Player State.
// - Puts the player into the combat state.
// - Adds enemy name and enemy health to struct and stores in DB.
// - Uses player name and enemy name for text generation.
// - Returns: flavor text.
// - Future
//   - Return image generation.
//   - Add more enemy types and varying levels.
//   - Add the chance to find items.
//   - Add the chance to find small quest. (Fight but also with gold as a reward. Let AI do the quest creation.)
//
// encore:api public
func (s *Service) Search(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.SearchForFight(s.db, secrets.NovelApiKey, params.Msg, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}

// Attack Command: Attack: - Enacts one cycle in the current fight if the player is in is one. Requires: Attack-type. Returns: flavor text.
// - Description: "Attacks the enemy."
// - Can be done in combat state.
// - Uses the player and enemy and weapon for player attack text gen.
// - Uses the player and enemy and armor for enemy attack text gen.
// - The player always moves before the enemy. If they kill the monster in that move, the player wins.
// - If the enemy is defeated, this is used in the text gen. The player has a chance of getting an item. If they do, this is also used in text gen and the item is added to their inventory. Xp is added to player struct. If level up, the get a message as well.
// - If the player is defeated, the player is now in the dead state and dying is used in the enemy attack text gen.
// - Every cycle stores in DB. Changing player health, enemy health and potentially items, enemy name, and player state.
// - Returns flavor text.
// - Future
//   - Image gen for the players attack/kill.
//   - Image gen for the enemies attack/kill.
//   - XP based weapon and armor you are using.
//
// encore:api public
func (s *Service) Attack(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.AttackEnemy(s.db, secrets.NovelApiKey, params.Msg, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}

// Store Command: Store - Opens the store if the player is in the town. Optional: Buy and Sell. Optional: Item. Optional: number of items to buy/sell. Returns: flavor text.
// - Description: "Browses the store when in town. Can also be used to buy and sell items."
// - Can be done in normal state.
// - Optional: Buy - the player chooses to buy the item they specify.
// - Optional: Sell - the player chooses to sell the item they specify.
//   - If the player does not provide any arguments, we return the store's current inventory. (everything at any quantity)
//   - If the player provides Buy or sell with no item, we reutrn funny flavor text.
//   - If the player provides item with no buy or sell, we return flavor text if the item is in the store.
//   - If the player provides Buy with an item:
//   - If the item is in shop:
//   - If they have enough gold for their request: they recieve a report on gaining the items they asked for and losing gold.
//   - If they do not have enough gold they get funny flavor text.
//   - If the item is not in the shop, they get funny flavor text.
//   - If the player provides Sell with an item:
//   - If they have the stated number of items: they recieve a report on losing the items they sold and gaining gold.
//   - If they do not have the stated number of items: they get funny flavor text.
//     *** If the player does not provide a number for buy or sell, the number 1 will be substituted.
//
// - Stores Player data in DB if transaction is made.
// - Uses the item and transaction data for text generation.
// - Returns: flavor text.
//
// encore:api public
func (s *Service) Store(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.PlayerStore(s.db, secrets.NovelApiKey, params.Msg, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}

// Cheatmode Command: Cheatmode - Adds weapons armor potions and potions+ to the players inv for testing Requires: None. Returns: text.
// Returns: static text
//
// encore:api public
func (s *Service) Cheatmode(ctx context.Context, params *gamecore.BackendProcessorRequest) (*gamecore.DiscordMessageResponse, error) {
	eb := errs.B()
	response, err := orch.CheatMode(s.db, params.PlayerId)
	if err != nil {
		return nil, eb.Code(errs.NotFound).Msg("Internal Server Error").Err()
	}
	return &response, nil
}
