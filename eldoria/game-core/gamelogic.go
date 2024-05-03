package gamecore
import(
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
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
	i_apple			int
	i_potion		int
	i_potionPlus	int
	c_gold			int
	b_gold			int
}

type Character struct {
	username			string
	user				string
	c_level				int
	c_health			int
	m_health			int
	b_health			int
	s_strength			int	
	s_agility			int
	s_constitution		int
	s_intelligence		int
	s_wisdom			int
	w_s_sword			int
	w_s_axe				int
	w_s_spear			int
	p_state				string
	c_area				string
	c_e_weapon			int
	c_e_armor			int
	inventory 			Inventory
}

type Weapon struct {
	index				int
	name				string
	atk_mod				int
	s_mod				int
	c_mod				int
	agi_mod				int
	cost				int

}

type Armor struct {
	index				int
	name				string
	atk_mod				int
	s_mod				int
	c_mod				int
	agi_mod				int
	cost				int

}


func printCharacter(pc *Character){
	str, _ := json.MarshalIndent(pc, "", "\t")
	fmt.Println(string(str))
	fmt.Printf("%+v\n", pc)
}


func checkState(s string) ( string, bool ) {

	/*
	normal	Normal, non-combat, non-blocking
	combat	In combat
	dead	dead
	*/

	switch s:
	"normal":
		return "check ok",  true
	"combat":
		return "player in combat", false
	"dead":
		return "player is deceased", false
	default:
		return "default pass", true
}


func createPlayer(username string) (string, string, Character) {
	//player creation

		userStats := Character{
		username: username,		
		user: "Ham",					
		c_level: 1,						
		c_health: 100,					
		m_health: 100,					
		b_health: 0,					
		s_strength: 10,					
		s_agility: 10,					
		s_constitution: 10,				
		s_intelligence: 10,				
		s_wisdom: 10,					
		w_s_sword: 10,					
		w_s_axe: 0,						
		w_s_spear: 0,					
		p_state: "normal",				
		c_area: "town",					
		c_e_weapon: 0,					
		c_e_armor: 0,				
		inventory: Inventory{
			i_apple: 1,						
			i_potion: 0,					
			i_potionPlus: 0,				
			c_gold: 0,						
			b_gold: 0,						

		},
	}
	
	fmt.Printf("%+v\n", userStats)
	
	message := "Create player: UserID passed in: " + username
	prompt := "You awaken on a cushion of wildflowers in a small clearing near the edge of some woods.  In the distance you can see a small village.  You stand up, brush yourself off, and head into the village."

	return message, prompt, userStats

}


func getStatus(pc *Character, command string) ( string, string ) {
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started
	
		decide the format for this state tracking 
	*/

	msg, pass := checkState(pc.p_state)

	if pass == false {
		fmt.Println("User unable to do this command at this time. reason: " + msg)
		return "Check State Fail: " + msg
	}

	args := strings.Fields(command)

	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "check subject not provided" , "You though about checking something, but forgot it the moment you begin to act.  You hate it when that happens."
	}
	
	switch verb := args[1]; verb {
		case "health":
			fmt.Println("you checked your health!")
			message := "your current HP is: " + pc.c_health + " + (" + pc.b_health + " bonus hp)/" + pc.m_health
		case "stats":
			fmt.Println("you checked your stats!")
			message := "current stats: Strength:" + strconv.Itoa(pc.s_strength) + " Agility: " + strconv.Itoa(pc.s_agility) + " Constitution: " + strconv.Itoa(pc.s_constitution) + " Intelligence: " + strconv.Itoa(pc.s_intelligence) + " Wisdom: " + strconv.Itoa(pc.s_wisdom)
		case "inventory":
			fmt.Println("you checked your inventory!")
			message := "Your backpack contains: " + strconv.Itoa(pc.inventory.i_apple) + " Apples, " + strconv.Itoa(pc.inventory.i_potion) + " Potions, " + strconv.Itoa(pc.inventory.i_potionPlus) + " Plus Potions. You have " + strconv.Itoa(pc.inventory.c_gold) + " gold."
		default:
			fmt.Println("you didn't check anything!")
	}
	
	message:= "Get status: " + pc.username
	prompt := "You check your guts, looks like everything is there!"
	
	fmt.Printf("Fields are: %q", args)

	return message, prompt
}


func moveArea(pc *Character, command string) ( string, string) {
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

	args := strings.Fields(command)
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "You try to move, but your inability to choose a direction fixes you in place."
	}
	
	area := args[1]

	switch area {
		case "Town":
			fmt.Println("you travel to the town!"); pc.c_area = "town"
		case "Plains":
			fmt.Println("you travel to the plains!"); pc.c_area = "plains"
		case "Forest":
			fmt.Println("you travel to the forest!"); pc.c_area = "forest"
		case "Cave":
			fmt.Println("you travel to the cave!"); pc.c_area = "cave"
		case "Dungeon":
			fmt.Println("you travel to the dungeon!"); pc.c_area = "dungeon"
		default:
			fmt.Println("Invalid location!")
	}

	message:= "Move area: " + username + " to " + area
	fmt.Printf("Fields are: %q", args)

	prompt := "You collect your belongings, and set off for the " + area
	return message, prompt
}


func store(pc *Character, command string) ( string, string ) {
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	err := checkState( pc.username )
	if err == false {
		fmt.Println("User unable to do this command at this time.")
		return "bad state", "bad state"
	}

	args := strings.Fields(command)
	if len(args) < 3 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "The shopkeeper looks bored with your window shopping."
	}
	
	verb := args[1]
	item := args[2]
	amount, e1 := strconv.Atoi(args[3]);

	if e1 == nil {
		fmt.Printf("%T \n %v", amount, amount)
	
	}

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
	message:= "Store menu: Action: " + verb + "item:" + item + "count: " + count

	return message, prompt
}


func cheatMode(pc *Character, command string) (string, string) {
	fmt.Println("Cheat items added")

	pc.inventory.c_gold = pc.inventory.c_gold + 500
	pc.inventory.i_apple = pc.inventory.i_apple + 10
	pc.inventory.i_potion = pc.inventory.i_potion + 10
	pc.inventory.i_potionPlus = pc.inventory.i_potionPlus + 10

	message := "Cheat items added to inventroy"
	prompt := "Your backpack suddenly becomes heavier, as if several new things just appeared in it."

	return message, prompt
}


func item(pc *Character, command string) ( string, string ) {
	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	args := strings.Fields(command)
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "You look at your empty hands, unsure of what you were trying to do."
	}
	prompt := "You use a " + args[2]
	message:= "Used item: " + args[1]

	fmt.Printf("Fields are: %q", args)

	return message, prompt
}


func combat(pc *Character, command string) (string, string) {
	// fight

	/*
		- check if the player is in a valid state
		- search for an enemy
		- set the combat state
		- advance combat as the player selects the attack method
		- update the player with the turn results
		- complete combat and award experience
	*/

	args := strings.Fields(command)
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "You feel like punching something, but don't know what."
	}

	message:= "Player is in combat:"
	prompt := "You attack the monster"

	return message, prompt
}


// internal function to determine if teamwork should be used in combat
func inTeam(username string) string {
	// is the user in a team currently?

	message:= "Team check: user checked: " + username

	return message
}


func getState(pc *Character, command string) ( string, string ) {
	//Is the user in town?  Is the user in combat?

	args := strings.Fields(command)
	fmt.Printf("Fields are: %q", args)
	if len(args) < 1 {
		fmt.Println("not enough arguments")
		return "not enough arguments", "You don't even know what state you are in."
	}

	message:= "Check use state: " + username
	prompt := "You look down at yourself and assess the damage."

	return message, prompt
}

	// creature list

/*
	Field
		Giant Rat
			bite
			scratch
			infection
	
		Feral Hound
			bite
			scratch
			tear

		Kobold
			slash
			bite
			call for aid

	Forest
		Slime
			slam
			acid
			slow

		Wolf
			bite
			scratch
			tear

		Goblin
			slash
			kick
			rush attack


*/

