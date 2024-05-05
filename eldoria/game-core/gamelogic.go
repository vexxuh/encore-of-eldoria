package gamecore

import (
	"encoding/json"
	models "encore.app/eldoria/game-core/data"
	"fmt"
	"strconv"
	"strings"
	"slices"
)

// https://www.geeksforgeeks.org/nested-structure-in-golang/

//Discord username | uid
//Player character name
//Character level
//Current Health
//Max Health
//Bonus Health
//Strength Stat
//Agility Stat
//Constitution Stat
//Intelligence Stat
//Wisdom Stat
//Sword type skill
//Axe type skill
//Spear type skill
//player state
//current location
//Equipped weapon index
//Equipped Armor
//apples in inventory
//potions in inventory
//Plus Potions in inventory
//gold in inventory
//gold in bank

func main() {
	userStats := Character{
		Inventory: Inventory{
			I_apple:      1,
			I_potion:     0,
			I_potionPlus: 0,
			C_gold:       0,
			B_gold:       0,
		},
		InventoryId:    1,
		Username:       "username",
		User:           "name",
		C_level:        1,
		C_experience:   0,
		C_health:       100,
		M_health:       100,
		B_health:       0,
		S_strength:     10,
		S_agility:      10,
		S_constitution: 10,
		S_intelligence: 10,
		S_wisdom:       10,
		W_s_sword:      0,
		W_e_sword:      0,
		W_s_axe:        0,
		W_e_axe:        0,
		W_s_spear:      0,
		W_e_spear:      0,
		P_state:        "normal",
		C_area:         "town",
		C_e_weapon:     0,
		C_e_armor:      0,
	}

	printCharacter(&userStats)
}

func printCharacter(pc *Character) {
	str, _ := json.MarshalIndent(pc, "", "\t")
	fmt.Println(string(str))
	fmt.Printf("%+v\n", pc)
}

func checkState(s string) (string, bool) {

	/*
		normal	Normal, non-combat, non-blocking
		combat	In combat
		dead	dead
	*/

	switch s {
	case "normal":
		return "check ok", true
	case "combat":
		return "player in combat", false
	case "dead":
		return "player is deceased", false
	default:
		return "default pass", true
	}
}

func createPlayer(username string, name string) (string, string, Character) {
	//player creation

	userStats := models.Character{
		Username:       username,
		User:           name,
		C_level:        1,
		C_experience:   0,
		C_health:       100,
		M_health:       100,
		B_health:       0,
		S_strength:     10,
		S_agility:      10,
		S_constitution: 10,
		S_intelligence: 10,
		S_wisdom:       10,
		W_s_sword:      0,
		W_e_sword:      0,
		W_s_axe:        0,
		W_e_axe:        0,
		W_s_spear:      0,
		W_e_spear:      0,
		P_state:        "normal",
		C_area:         "town",
		C_e_weapon:     0,
		C_e_armor:      0,
		Inventory: models.Inventory{
			I_apple:      1,
			I_potion:     0,
			I_potionPlus: 0,
			C_gold:       0,
			B_gold:       0,
		},
	}

	fmt.Printf("%+v\n", userStats)

	message := "Create player: UserID passed in: " + username
	prompt := "You awaken on a cushion of wildflowers in a small clearing near the edge of some woods.  In the distance you can see a small village.  You stand up, brush yourself off, and head into the village."

	return message, prompt, userStats
}

func getStatus(pc *Character, command string) (string, string) {
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started

		decide the format for this state tracking

		accepted nouns:
			health: return your health information
			stats: return your stats
			inventory: return your current inventory
	*/

	args := strings.Fields(command)

	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "check subject not provided", "You though about checking something, but forgot it the moment you begin to act.  You hate it when that happens."
	}

	strength := strconv.Itoa(pc.s_strength)
	agility := strconv.Itoa(pc.s_agility)
	constitution := strconv.Itoa(pc.s_constitution)
	intelligence := strconv.Itoa(pc.s_intelligence)
	wisdom := strconv.Itoa(pc.s_wisdom)
	c_health := strconv.Itoa(pc.c_health)
	b_health := strconv.Itoa(pc.b_health)
	m_health := strconv.Itoa(pc.m_health)
	message := "Get status: " + pc.username

	switch noun := args[1]; noun {
	case "health":
		fmt.Println("you checked your health!")
		message = "your current HP is: " + c_health + " + (" + b_health + " bonus hp)/" + m_health
	case "stats":
		fmt.Println("you checked your stats!")
		message = "current stats: Strength:" + strength + " Agility: " + agility + " Constitution: " + constitution + " Intelligence: " + intelligence + " Wisdom: " + wisdom
	default:
		fmt.Println("you didn't check anything!")
	}

	prompt := "You check your guts, looks like everything is there!"

	fmt.Printf("Fields are: %q", args)

	return message, prompt
}

func moveArea(pc *Character, command string) (string, string) {

	// navigation
	/*
		- check if the player can move (stuck in combat? middle of player creation or exchange?)
		- move to the selected area
		- provide flavor text for the area
	*/

	// err := checkState(username)
	// if err == false {
	// fmt.Println("User unable to do this command at this time.")
	// return "bad state", "bad state"
	// }

	msg, pass := checkState(pc.p_state)

	if pass == false {
		fmt.Println("User unable to do this command at this time. reason: " + msg)
		return "Check State Fail: " + msg, "I didn't work"
	}

	args := strings.Fields(command)
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "You try to move, but your inability to choose a direction fixes you in place."
	}

	area := args[1]

	switch area {
	case "Town":
		fmt.Println("you travel to the town!")
		pc.c_area = "town"
	case "Plains":
		fmt.Println("you travel to the plains!")
		pc.c_area = "plains"
	case "Forest":
		fmt.Println("you travel to the forest!")
		pc.c_area = "forest"
	case "Cave":
		fmt.Println("you travel to the cave!")
		pc.c_area = "cave"
	case "Dungeon":
		fmt.Println("you travel to the dungeon!")
		pc.c_area = "dungeon"
	default:
		fmt.Println("Invalid location!")
	}

	message := "Move area: " + pc.username + " to " + area
	fmt.Printf("Fields are: %q", args)

	prompt := "You collect your belongings, and set off for the " + area
	return message, prompt
}

func store(pc *Character, command string) (string, string) {
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/

	i1 := []string{"apple", "potion", "plus potion"}
    i2 := []int{5, 50, 100}
	items := ""
	amount:= 0

	msg, pass := checkState(pc.p_state)

	if pass == false {
		fmt.Println("User unable to do this command at this time. reason: " + msg)
		return "Check State Fail: " + msg, "User unable to do this command at this time."
	}

	args := strings.Fields(command)

	// if just 'store' command, show items for sale
	if len(args) == 1 {
		fmt.Println("list items")
		items := ""
		
		// Print store items
		for i := range i1 {
            fmt.Println(i1[i])
            fmt.Println(i2[i])
            items = items + i1[i] + " for " + strconv.Itoa(i2[i]) + "gp\n"
           }
		return "list of items requested", "'Welcome!' says the shopkeeper.  'Take a look around and let me know what you would like to buy!'" + items
	}

	verb := args[1]
	item := args[2]

	amount, e1 := strconv.Atoi(args[3])

	if e1 == nil {
		fmt.Printf("%T \n %v", amount, amount)
		amount = 1
	}

	// find the index for the requested item
	var itemIndex int = -1
	for i, item := range i1 {
		if item == i1[i] {
		itemIndex = i
		break
		}
	}

	

	switch verb {
	case "buy":
		fmt.Println("Buying")

		if itemIndex == -1 {
			return "item does not exist", "'I don't think we have any of that.' says the shopkeeper"
		}

		if amount * i2[itemIndex] > pc.inventory.c_gold {
			return "not enough gold", "The shopkeeper says 'Sorry guy, it looks like you don't have enough gold for that.'"
		}

		pc.inventory.c_gold = pc.inventory.c_gold - amount * i2[itemIndex]

		switch itemIndex {
		case 1:
			pc.inventory.i_apple = pc.inventory.i_apple + amount
		case 2:
			pc.inventory.i_potion = pc.inventory.i_potion + amount
		case 3:
			pc.inventory.i_potionPlus = pc.inventory.i_potionPlus + amount
		}

	}




	case "sell":
		fmt.Println("Selling")
	case "info":
		fmt.Println("Information")
	default:
		fmt.Println("The shopkeeper looks at you confused")
	}

	prompt := "You ask the shopkeeper about buying a" + args[2]
	message := "Store menu: Action: " + verb + "item:" + item + "count: " + args[3]

	return message, prompt
}


func cheatMode(pc *Character, command string) (string, string) {
	fmt.Println("Cheat items added")

	pc.inventory.c_gold = pc.inventory.c_gold + 500
	pc.inventory.i_apple = pc.inventory.i_apple + 10
	pc.inventory.i_potion = pc.inventory.i_potion + 10
	pc.inventory.i_potionPlus = pc.inventory.i_potionPlus + 10

	message := "Cheat items added to inventory"
	prompt := "Your backpack suddenly becomes heavier, as if several new things just appeared in it."

	return message, prompt
}

func inventory(pc *Character, command string) (string, string) {

	args := strings.Fields(command)
	prompt := " - "
	message := " - "

	apples := strconv.Itoa(pc.inventory.i_apple)
	potion := strconv.Itoa(pc.inventory.i_potion)
	potionPlus := strconv.Itoa(pc.inventory.i_potionPlus)

	if len(args) == 1 {
		fmt.Println("Inventory check")
		contents := "Bag:\nApples: " + apples + "\nPotions: " + potion + "\nPlusPotions: " + potionPlus
		message = "you check your inventory" + contents
		prompt = "You look inside your bag. \n" + contents
	}

	if len(args) == 2 && args[1] == "use" {

		//equip

		if len(args) == 3 && args[1] == "use" {
			fmt.Println("not enough arguments")
			return "not enough arguments", "You look at your empty hands, unsure of what you were trying to do."
		}

		prompt = "You use a " + args[2]
		message = "Used item: " + args[1]

		fmt.Printf("Fields are: %q", args)

		return message, prompt
	}
	return message, prompt

}
