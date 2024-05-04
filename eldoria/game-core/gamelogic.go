package gamelogic

import (
	"encoding/json"
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

type Inventory struct {
	i_apple      int
	i_potion     int
	i_potionPlus int
	c_gold       int
	b_gold       int
}

type Character struct {
	username       string
	user           string
	c_level        int
	c_experience   int
	c_health       int
	m_health       int
	b_health       int
	s_strength     int
	s_agility      int
	s_constitution int
	s_intelligence int
	s_wisdom       int
	w_s_sword      int
	w_e_sword      int
	w_s_axe        int
	w_e_axe        int
	w_s_spear      int
	w_e_spear      int
	p_state        string
	c_area         string
	c_e_weapon     int
	c_e_armor      int
	inventory      Inventory
}

type Weapon struct {
	index   int
	name    string
	atk_mod int
	s_mod   int
	c_mod   int
	agi_mod int
	cost    int
}

type Armor struct {
	index   int
	name    string
	atk_mod int
	s_mod   int
	c_mod   int
	agi_mod int
	cost    int
}

// func main() {
	// userStats := Character{
		// username:       "username",
		// user:           "name",
		// c_level:        1,
		// c_experience:   0,
		// c_health:       100,
		// m_health:       100,
		// b_health:       0,
		// s_strength:     10,
		// s_agility:      10,
		// s_constitution: 10,
		// s_intelligence: 10,
		// s_wisdom:       10,
		// w_s_sword:      0,
		// w_e_sword:      0,
		// w_s_axe:        0,
		// w_e_axe:        0,
		// w_s_spear:      0,
		// w_e_spear:      0,
		// p_state:        "normal",
		// c_area:         "town",
		// c_e_weapon:     0,
		// c_e_armor:      0,
		// inventory: Inventory{
			// i_apple:      1,
			// i_potion:     0,
			// i_potionPlus: 0,
			// c_gold:       0,
			// b_gold:       0,
		// },
	// }
// 
	// printCharacter(&userStats)
// }

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

	userStats := Character{
		username:       username,
		user:           name,
		c_level:        1,
		c_experience:   0,
		c_health:       100,
		m_health:       100,
		b_health:       0,
		s_strength:     10,
		s_agility:      10,
		s_constitution: 10,
		s_intelligence: 10,
		s_wisdom:       10,
		w_s_sword:      0,
		w_e_sword:      0,
		w_s_axe:        0,
		w_e_axe:        0,
		w_s_spear:      0,
		w_e_spear:      0,
		p_state:        "normal",
		c_area:         "town",
		c_e_weapon:     0,
		c_e_armor:      0,
		inventory: Inventory{
			i_apple:      1,
			i_potion:     0,
			i_potionPlus: 0,
			c_gold:       0,
			b_gold:       0,
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

	msg, pass := checkState(pc.p_state)

	if pass == false {
		fmt.Println("User unable to do this command at this time. reason: " + msg)
		return "Check State Fail: " + msg, "User unable to do this command at this time."
	}

	args := strings.Fields(command)
	if len(args) < 3 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "The shopkeeper looks bored with your window shopping."
	}

	verb := args[1]
	item := args[2]
	amount, e1 := strconv.Atoi(args[3])

	if e1 == nil {
		fmt.Printf("%T \n %v", amount, amount)
		amount = 1

	}

	// if just 'store' command, show items for sale

	switch verb := args[1]; verb {
	case "buy":
		fmt.Println("Buying")
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
