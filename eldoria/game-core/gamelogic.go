package gamecore

import (
	"encoding/json"
	models "encore.app/eldoria/game-core/data"
	"fmt"
	"strconv"
	"strings"
	"errors"
	"math/rand/v2"
)

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
		Creature:	models.Creature{
			C_name:       "none",
			C_id:     	 1,
			C_level:     1,
			C_experience: 0,
			C_c_health:	  0,
			C_m_health:   0,
			C_attack:     0,
			C_defense:    0,

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
	msg1, msg2, _ = getStatus(&userStats, "check status")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2, _ = getStatus(&userStats, "check health")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2, _ = MoveArea(&userStats, "move town")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2, _ = Store(&userStats, "store buy apple")
	fmt.Println(msg1); fmt.Println(msg2)
	printCharacter(&userStats)
	msg1, msg2, _ = Store(&userStats, "store sell apple")
	fmt.Println(msg1); fmt.Println(msg2)
	printCharacter(&userStats)
	msg1, msg2, _ = CheatMode(&userStats, "cheat mode")
	fmt.Println(msg1); fmt.Println(msg2)
	msg1, msg2, _ = Inventory(&userStats, "inventory")
	fmt.Println(msg1); fmt.Println(msg2)
	printCharacter(&userStats)
	msg1, msg2, _ = Inventory(&userStats, "inventory use apple")
	fmt.Println(msg1); fmt.Println(msg2)
	printCharacter(&userStats)
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

func creatureHealthMsg(ch int, mh int) (int, string) {
	hmsg := ""

	switch {
		case ch / mh * 100 == 100:
			hmsg = "untouched"
		case ch / mh * 100 > 80:
			hmsg = "slightly wounded"
		case ch / mh * 100  > 50:
			hmsg = "fairly injured"
		case ch / mh * 100  > 30:
			hmsg = "critically wounded"
		case ch / mh * 100  > 15:
			hmsg = "near death"
		default:
			hmsg = "on the footsteps of oblivion"
	}
	return (ch / mh * 100), hmsg
}

func CreatePlayer(pc *models.Character, username string, name string) (string, string, models.Character, error) {
	//player creation

if pc != nil {
	return "", "", models.Character{}, errors.New("character already exists")
}


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

	return message, prompt, userStats, nil
}

func getStatus(pc *models.Character, command string) (string, string, error) {
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
		return "check subject not provided", "You though about checking something, but forgot it the moment you begin to act.  You hate it when that happens.", nil
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

	return message, prompt, nil
}

func MoveArea(pc *models.Character, command string) (string, string, error) {

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
		return "Check State Fail: " + msg, "I didn't work", nil
	}

	args := strings.Fields(command)
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "You try to move, but your inability to choose a direction fixes you in place.", nil
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
		return "", "", errors.New("invalid area")
	}

	message := "Move area: " + pc.Username + " to " + area
	fmt.Printf("Fields are: %q", args)

	prompt := "You collect your belongings, and set off for the " + area
	return message, prompt, nil
}

func Store(pc *models.Character, command string) (string, string, error) {
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

	msg, pass := checkState(pc.P_state)

	if pass == false {
		fmt.Println("User unable to do this command at this time. reason: " + msg)
		return "Check State Fail: " + msg, "User unable to do this command at this time.", nil
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
		return "list of items requested", "'Welcome!' says the shopkeeper.  'Take a look around and let me know what you would like to buy!'" + items, nil
	}

	verb := args[1]
	item := args[2]
	amount := 1

		
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
			return "item does not exist", "'I don't think we have any of that.' says the shopkeeper", nil
		}

		if amount * i2[itemIndex] > pc.Inventory.C_gold {
			return "not enough gold", "The shopkeeper says 'Sorry guy, it looks like you don't have enough gold for that.'", nil
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
			return "item does not exist", "'I not sure what you want me to buy.' says the shopkeeper", nil
		}

		switch itemIndex {
		case 1:
			if pc.Inventory.I_apple < amount {
				return "not enough items", "The shopkeeper says 'Sorry guy, you don't have that many to sell.", nil
			}
			if pc.Inventory.I_apple >= amount {
				pc.Inventory.I_apple = pc.Inventory.I_apple - amount
				pc.Inventory.C_gold = pc.Inventory.C_gold + amount * i3[itemIndex]

				return "sold items for " + strconv.Itoa(amount * i3[itemIndex]), "The shopkeeper says 'Thanks!", nil
			}

		case 2:
			if pc.Inventory.I_potion < amount {
				return "not enough items", "The shopkeeper says 'Sorry guy, you don't have that many to sell.", nil
			}
			if pc.Inventory.I_potion >= amount {
				pc.Inventory.I_potion = pc.Inventory.I_potion - amount
				pc.Inventory.C_gold = pc.Inventory.C_gold + amount * i3[itemIndex]

				return "sold items for " + strconv.Itoa(amount * i3[itemIndex]), "The shopkeeper says 'Thanks!", nil
			}

		case 3:
			if pc.Inventory.I_potionPlus < amount {
				return "not enough items", "The shopkeeper says 'Sorry guy, you don't have that many to sell.", nil
			}
			if pc.Inventory.I_potionPlus >= amount {
				pc.Inventory.I_potionPlus = pc.Inventory.I_potionPlus - amount
				pc.Inventory.C_gold = pc.Inventory.C_gold + amount * i3[itemIndex]

				return "sold items for " + strconv.Itoa(amount * i3[itemIndex]), "The shopkeeper says 'Thanks!", nil
			}
		}


	case "info":
		fmt.Println("Information")
		switch itemIndex {
		case 1:
			return "info apple", "The shopkeeper says 'I just got those fresh this morning! Eating one is sure to restore 10 health points.'", nil
		case 2:
			return "info potion", "The shopkeeper says 'Health potions are very popular with you adventurers. Drinking one is sure to restore 50 health points.'", nil
		case 3:
			return "info Plus potion", "The shopkeeper says 'You have a keen eye for quality! Drinking one is sure to restore 100 health points.'", nil
		}

	default:
		fmt.Println("The shopkeeper looks at you confused")
	}

	return "leaving store unexpectedly" , "You have somehow left the shop by mysterious means.  Maybe you blew out the window?", nil
}

func CheatMode(pc *models.Character, command string) (string, string, error) {
	fmt.Println("Cheat items added")

	pc.Inventory.C_gold = pc.Inventory.C_gold + 500
	pc.Inventory.I_apple = pc.Inventory.I_apple + 10
	pc.Inventory.I_potion = pc.Inventory.I_potion + 10
	pc.Inventory.I_potionPlus = pc.Inventory.I_potionPlus + 10

	message := "Cheat items added to inventory"
	prompt := "Your backpack suddenly becomes heavier, as if several new things just appeared in it."

	return message, prompt, nil
}

func Inventory(pc *models.Character, command string) (string, string, error) {

	args := strings.Fields(command)
	prompt := " "
	message := " "
	item := ""

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
		contents := "Bag:\nApples: " + strconv.Itoa(i2[0]) + "\nPotions: " + strconv.Itoa(i2[1]) + "\nPlus_Potions: " + strconv.Itoa(i2[2])
		message = "you check your inventory" + contents
		prompt = "You look inside your bag. \n" + contents
	}

	if len(args) > 2 && args[1] == "use" && itemIndex > -1 {
		item = args[2]
		fmt.Println("use item index " + strconv.Itoa(itemIndex))

		switch {
		case i2[itemIndex] > 0:
			if pc.C_health + i3[itemIndex] > pc.M_health {
				pc.C_health = pc.M_health
				return "Healed to max", "You consume the " + i1[itemIndex] + " and heal to max health.", nil
			} else {
				pc.C_health = pc.C_health + i3[itemIndex]
				return "Healed up " + strconv.Itoa(i3[itemIndex]), "You consume the " + i1[itemIndex] + " and heal " + strconv.Itoa(i3[itemIndex]) + " health points.", nil
			}
		default:
			return "not enough item" + message, "You don't have one of those to use." + prompt, nil
		}


	}
	return "leaving inventory unexpectedly", "You have somehow left the inventory screen by mysterious means.", nil
}

func Combat(pc *models.Character, command string) (string, string, error) {

	args := strings.Fields(command)
	prompt := ""
	allowed := true

	switch pc.C_area {
	case "town":
		allowed = false
	default:
		allowed = true
	}

	c1 := []string{"turkey","kobold",	"wolf",	"bandit",	"goblin",	"slime",	"bear",	"spider",	"creeping mass",	"skeleton",	"imp",	"wyrm"}
	c2 := []int{	1, 		2,			3,		4,			5,			6,			7,		8,			9,					10,			11,		12}
	c3 := []int{	1, 		3,			5,		8,			12,			15,			20,		25,			30,					35,			40,		50}
	c4 := []int{	1, 		3,			5,		8,			12,			15,			20,		25,			30,					35,			40,		50}
	c5 := []int{	15,		25,			35,		35,			40,			45,			70,		40,			80,					70,			80,		100}
	c6 := []int{	15,		25,			35,		35,			40,			45,			70,		40,			80,					70,			80,		100}
	c7 := []int{	2, 		4,			6,		10,			7,			5,			20,		10,			18,					12,			15,		30}
	c8 := []int{	2, 		3,			5,		7,			7,			5,			15,		12,			20,					7,			18,		25}
	//c9 := []string{"plains", "plains", "plains", "forest", "forest",	"forest",	"cave",	"cave",		"cave",				"dungeon",	"dungeon","dungeon"}

/*
	C_name       string
	C_id     	 int
	C_level      int
	C_experience int
	C_c_health	 int
	C_m_health   int
	C_attack     int
	C_defense    int
	c9 
		plains
		forest
		cave
		dungeon
*/

	if len(args) == 1 {
		// if in combat, show status
		//if not in combat, trigger event for the area (combat, special event like a shop...)
			if pc.P_state == "normal" && allowed {
				fmt.Println("Combat: no arguments")
				return "not in combat.", "you are located in " + pc.C_area + ". Use the command 'combat search' to search for a monster", nil
			}
			
			if pc.P_state == "combat" {
				fmt.Println("Combat: no arguments. Combat engaged.")

				_, hmsg := creatureHealthMsg(pc.Creature.C_c_health, pc.Creature.C_c_health)

				return "You are in combat", "You are in combat.  You are being attacked by a " + pc.Creature.C_name + "(id:" + strconv.Itoa(pc.Creature.C_id) + "). It is " + hmsg + ". use the command 'combat attack monsterID' to attack that monster.", nil
			}

			if pc.P_state == "dead" {
				fmt.Println("Combat: action failed, player is dead.")
				return "combat failed, player is dead", "Try as you might, not a dead person yet has managed the will to ressurect themselves. You must create a new character to continue.", nil
			}		
	}

	if len(args) == 2 {
		//combat search
		//combat attack **future**  auto attack the weakest/last targeted enemy?
			if pc.P_state == "normal" && allowed && args[1] == "search" {
				fmt.Println("Combat: searching")
				c := rand.IntN(3)
				fmt.Println("random number: " + strconv.Itoa(c))
				
				multi := 1
				switch pc.C_area {
					case "plains":
						multi = 1
					case "forest":
						multi = 2
					case "cave":
						multi = 3
					case "dungeon":
						multi = 4
					default:
						multi = 1
					}

				pc.Creature.C_name = c1[c * multi]
				pc.Creature.C_id = c2[c * multi]
				pc.Creature.C_level = c3[c * multi]
				pc.Creature.C_experience = c4[c * multi]
				pc.Creature.C_c_health = c5[c * multi]
				pc.Creature.C_m_health = c6[c * multi]
				pc.Creature.C_attack = c7[c * multi]
				pc.Creature.C_defense = c8[c * multi]
				pc.P_state = "combat"
				fmt.Println("you found a " + pc.Creature.C_name)

				return "encounted enemy", "you come across a " + pc.Creature.C_name +". Get ready to fight!", nil
				}
	}

	if pc.P_state == "combat" && args[1] == "search" {
		fmt.Println("Combat: search attempted. Already in combat.")
		return "search failed, already in combat", "You are already in an event. use the command 'combat attack monsterID' or 'combat status'", nil		
	}

	if len(args) == 3 && args[1] == "attack" {
		weapon := "Ya' Mitts"

		fmt.Println("Combat: attacking with " + weapon)

		pSuccess := false
		pCrit := false
		pAttackRoll := rand.IntN(20) + 1
		pResult := "you missed"
		pReward := 0.0
		
		cSuccess := false
		cCrit := false
		cAttackRoll := rand.IntN(20) + 1
		cResult := "it missed"

		combatComplete := false

		pAttackMod := float64(pc.C_level) - (float64(pc.Creature.C_level) * 0.03) * 100
		// (difference in levels * 3%) + (difference in pc agility to creature speed * 3%) + ()

		cAttackMod := float64(pc.Creature.C_level) - (float64(pc.C_level) * 0.03 ) * 100
		// (difference in levels * 3%) + (difference in pc agility to creature speed * 3%) + ()


		pDamage := (rand.Float64() * 100) * float64(pc.S_strength)  - float64(pc.Creature.C_defense)
		// strength + weapon damage * 1-100% - creature defense 

		if pAttackMod < 0 {
			pAttackMod = 0
		}

		cDamage := (rand.Float64() * 100) * float64(pc.Creature.C_attack) - float64(pc.S_constitution) * (rand.Float64())
		// strength + weapon damage * 1-100% - creature defense 

		if pAttackMod < 0 {
			pAttackMod = 0
		}

		if ((float64(pAttackRoll) / 20 * 100) + pAttackMod) > 50 {
			pSuccess = true
		} 

		if pAttackRoll == 20 {
			pSuccess = true
			pCrit = true
			pDamage = float64(pc.S_strength) * 1.2
		}
			fmt.Println("Player critical hit: " + strconv.FormatBool(pCrit))

		if ((float64(cAttackRoll) / 20 * 100) + cAttackMod) > 50 {
			cSuccess = true
		} 

		if cAttackRoll == 20 {
			cSuccess = true
			cCrit = true
			cDamage = float64(pc.Creature.C_attack) * 1.2
		}
			fmt.Println("Creature critical hit: " + strconv.FormatBool(cCrit))

		if pSuccess {
			pc.Creature.C_c_health = pc.Creature.C_c_health - int(pDamage)
			pResult = "you hit for " + strconv.Itoa(int(pDamage))
		}

		if pc.Creature.C_c_health < 1 {
			combatComplete = true
			pReward = float64(pc.Creature.C_level) * (rand.Float64() * .2 )
			pResult = pResult + "\nYou slayed the creature! You collected enough resources to earn " + strconv.Itoa(int(pReward)) + " gold."
			cResult = ""
			pc.Inventory.C_gold = pc.Inventory.C_gold + int(pReward)
		}

		if combatComplete == false && cSuccess == true {
			pc.C_health = pc.C_health - int(cDamage)
			cResult = "the creature hits for " + strconv.Itoa(int(cDamage))
		}

		if pc.C_health < 1 {
			pc.P_state = "dead"
			cResult = cResult + "\nyou were slain."
		}

		return pResult + "\n" + cResult, pResult + "\n" + cResult, nil
	}

	fmt.Println("Combat: EOF")
	return "Combat EOF", prompt, nil
}

