package main

import (
	"encoding/json"
	models "encore.app/eldoria/game-core/data"
	"fmt"
	"strconv"
	"strings"
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
	userStats := models.Character{
		Inventory: models.Inventory{
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
	msg1, msg2 := " ", " "
	msg1, _ = checkState(userStats.P_state)
	fmt.Println(msg1)
	msg1, msg2 = getStatus(userStats, "check status")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2 = getStatus(userStats, "check health")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2 = moveArea(userStats, "move town")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2 = store(userStats, "store buy apple")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2 = store(userStats, "store sell apple")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2 = cheatMode(userStats, "cheat mode")
	fmt.Println(msg1); fmt.Println(msg2)
	

	
	


}

func printCharacter(pc *models.Character) {
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

func createPlayer(username string, name string) (string, string, models.Character) {
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

func getStatus(pc *models.Character, command string) (string, string) {
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

	strength := strconv.Itoa(pc.S_strength)
	agility := strconv.Itoa(pc.S_agility)
	constitution := strconv.Itoa(pc.S_constitution)
	intelligence := strconv.Itoa(pc.S_intelligence)
	wisdom := strconv.Itoa(pc.S_wisdom)
	C_health := strconv.Itoa(pc.C_health)
	B_health := strconv.Itoa(pc.B_health)
	M_health := strconv.Itoa(pc.M_health)
	message := "Get status: " + pc.Username

	switch noun := args[1]; noun {
	case "health":
		fmt.Println("you checked your health!")
		message = "your current HP is: " + C_health + " + (" + B_health + " bonus hp)/" + M_health
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

func moveArea(pc *models.Character, command string) (string, string) {

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

	msg, pass := checkState(pc.P_state)

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
		pc.C_area = "town"
	case "Plains":
		fmt.Println("you travel to the plains!")
		pc.C_area = "plains"
	case "Forest":
		fmt.Println("you travel to the forest!")
		pc.C_area = "forest"
	case "Cave":
		fmt.Println("you travel to the cave!")
		pc.C_area = "cave"
	case "Dungeon":
		fmt.Println("you travel to the dungeon!")
		pc.C_area = "dungeon"
	default:
		fmt.Println("Invalid location!")
	}

	message := "Move area: " + pc.Username + " to " + area
	fmt.Printf("Fields are: %q", args)

	prompt := "You collect your belongings, and set off for the " + area
	return message, prompt
}

func store(pc *models.Character, command string) (string, string) {
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/

	i1 := []string{"apple", "potion", "plus_potion"}
    i2 := []int{5, 50, 100}
    i3 := []int{2, 25, 50}

	items := ""
	amount:= 0

	msg, pass := checkState(pc.P_state)

	if pass == false {
		fmt.Println("User unable to do this command at this time. reason: " + msg)
		return "Check State Fail: " + msg, "User unable to do this command at this time."
	}

	args := strings.Fields(command)

	// if just 'store' command, show items for sale
	if len(args) == 1 {
		fmt.Println("list items")
		
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
	for i := range i1 {
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

		if amount * i2[itemIndex] > pc.Inventory.C_gold {
			return "not enough gold", "The shopkeeper says 'Sorry guy, it looks like you don't have enough gold for that.'"
		}

		pc.Inventory.C_gold = pc.Inventory.C_gold - amount * i2[itemIndex]

		switch itemIndex {
		case 1:
			pc.Inventory.I_apple = pc.Inventory.I_apple + amount
		case 2:
			pc.Inventory.I_potion = pc.Inventory.I_potion + amount
		case 3:
			pc.Inventory.I_potionPlus = pc.Inventory.I_potionPlus + amount
		}

	case "sell":
		fmt.Println("Selling")

		if itemIndex == -1 {
			return "item does not exist", "'I not sure what you want me to buy.' says the shopkeeper"
		}

		switch itemIndex {
		case 1:
			if pc.Inventory.I_apple < amount {
				return "not enough items", "The shopkeeper says 'Sorry guy, you don't have that many to sell."
			}
			if pc.Inventory.I_apple >= amount {
				pc.Inventory.I_apple = pc.Inventory.I_apple - amount
				pc.Inventory.C_gold = pc.Inventory.C_gold + amount * i3[itemIndex]

				return "sold items for " + strconv.Itoa(amount * i3[itemIndex]), "The shopkeeper says 'Thanks!"
			}

		case 2:
			if pc.Inventory.I_potion < amount {
				return "not enough items", "The shopkeeper says 'Sorry guy, you don't have that many to sell."
			}
			if pc.Inventory.I_potion >= amount {
				pc.Inventory.I_potion = pc.Inventory.I_potion - amount
				pc.Inventory.C_gold = pc.Inventory.C_gold + amount * i3[itemIndex]

				return "sold items for " + strconv.Itoa(amount * i3[itemIndex]), "The shopkeeper says 'Thanks!"
			}

		case 3:
			if pc.Inventory.I_potionPlus < amount {
				return "not enough items", "The shopkeeper says 'Sorry guy, you don't have that many to sell."
			}
			if pc.Inventory.I_potionPlus >= amount {
				pc.Inventory.I_potionPlus = pc.Inventory.I_potionPlus - amount
				pc.Inventory.C_gold = pc.Inventory.C_gold + amount * i3[itemIndex]

				return "sold items for " + strconv.Itoa(amount * i3[itemIndex]), "The shopkeeper says 'Thanks!"
			}
		}


	case "info":
		fmt.Println("Information")
		switch itemIndex {
		case 1:
			return "info apple", "The shopkeeper says 'I just got those fresh this morning! Eating one is sure to restore 10 health points.'"
		case 2:
			return "info potion", "The shopkeeper says 'Health potions are very popular with you adventurers. Drinking one is sure to restore 50 health points.'"
		case 3:
			return "info Plus potion", "The shopkeeper says 'You have a keen eye for quality! Drinking one is sure to restore 100 health points.'"
		}

	default:
		fmt.Println("The shopkeeper looks at you confused")
	}

	return "leaving store unexpectedly" , "You have somehow left the shop by mysterious means.  Maybe you blew out the window?"
}

func cheatMode(pc *models.Character, command string) (string, string) {
	fmt.Println("Cheat items added")

	pc.Inventory.C_gold = pc.Inventory.C_gold + 500
	pc.Inventory.I_apple = pc.Inventory.I_apple + 10
	pc.Inventory.I_potion = pc.Inventory.I_potion + 10
	pc.Inventory.I_potionPlus = pc.Inventory.I_potionPlus + 10

	message := "Cheat items added to inventory"
	prompt := "Your backpack suddenly becomes heavier, as if several new things just appeared in it."

	return message, prompt
}

func inventory(pc *models.Character, command string) (string, string) {

	args := strings.Fields(command)
	prompt := " "
	message := " "
	item := args[2]

	fmt.Printf("Fields are: %q", args)

	i1 := []string{"apple", "potion", "plus_potion"}
    i2 := []int{pc.Inventory.I_apple, pc.Inventory.I_potion, pc.Inventory.I_potionPlus}
    i3 := []int{10, 50, 100}

    var itemIndex int = -1
	for i := range i1 {
		if item == i1[i] {
			itemIndex = i
			break
		}
	}

	if len(args) == 1 {
		fmt.Println("Inventory check")
		contents := "Bag:\nApples: " + strconv.Itoa(i2[itemIndex]) + "\nPotions: " + strconv.Itoa(i2[itemIndex]) + "\nPlus_Potions: " + strconv.Itoa(i2[itemIndex])
		message = "you check your inventory" + contents
		prompt = "You look inside your bag. \n" + contents
	}

	if len(args) > 2 && args[1] == "use" && itemIndex > -1 {
		fmt.Println("use item index " + strconv.Itoa(itemIndex))

		switch {
		case i2[itemIndex] > 0:
			if pc.C_health + i3[itemIndex] > pc.M_health {
				pc.C_health = pc.M_health
				return "Healed to max", "You consume the " + i1[itemIndex] + " and heal to max health."
			} else {
				pc.C_health = pc.C_health + i3[itemIndex]
				return "Healed up " + strconv.Itoa(i3[itemIndex]), "You consume the " + i1[itemIndex] + " and heal " + strconv.Itoa(i3[itemIndex]) + " health points."
			}
		default:
			return "not enough item" + message, "You don't have one of those to use." + prompt
		}


	}
	return "leaving inventory unexpectedly" , "You have somehow left the inventory screen by mysterious means."

}
