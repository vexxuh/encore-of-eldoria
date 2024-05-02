package main
//package gamecore
import(
	"fmt"
	"strconv"
	"strings"
)
// https://www.geeksforgeeks.org/nested-structure-in-golang/

func main() {

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

	userStats := &Character{
		username: "Yaac-itysnorsh",		
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
	//command := "shop buy potion 10"
	//message, prompt := store(userStats, command)
	//fmt.Println("\n" + message + prompt)

	/*
	checkstate(username string) ( bool )
	createPlayer(username string) ( string, string )
	getStatus(username string, command string)  ( string, string )
	moveArea(username string, command string) ( string, string )
	store(username string, command string) ( string, string )
	item(username string, command string) ( string, string )
	combat(username string, command string) ( string, string )
	inTeam(username string) ( string )
	getState(username string, command string) ( string, string )
	*/
}



// functions
	// command parser
		// playerID, command, arguments
		//check if open session


func checkState(username string) bool {

	if username == "Ham" {
		return true
	} else {
		return false
	}
}

func createPlayer(username string) (string, string) {
	//player creation
	message := "Create player: UserID passed in: " + username
	prompt := "You awaken on a cushion of wildflowers in a small clearing near the edge of some woods.  In the distance you can see a small village.  You stand up, brush yourself off, and head into the village."

	return message, prompt

}

func getStatus(username string, command string) ( string, string ) {
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started
	
		decide the format for this state tracking 
	*/

	err := checkState(username)

	if err == false {
		fmt.Println("User unable to do this command at this time.")
		return "bad state", "bad state"
	}

	args := strings.Fields(command)

	if len(args) == 0 {
		fmt.Println("no commands passed")
		return "invalid" , "invalid"
	}
	
	switch verb := args[1]; verb {
		case "health":
			fmt.Println("you checked your health!") 
		default:
			fmt.Println("you didn't check anything!")
	}
	
	message:= "Get status: " + username
	prompt := "You check your backpack and see what items you have."

	
	fmt.Printf("Fields are: %q", args)

	return message, prompt
}


func moveArea(username string, command string) ( string, string) {
	// navigation
	/*
		- check if the player can move (stuck in combat? middle of player creation or exchange?)
		- move to the selected area
		- provide flavor text for the area

		Areas:
			Town
			Plains
			Forest
			Cave
			Dungeon
	*/

	err := checkState(username)
	if err == false {
		fmt.Println("User unable to do this command at this time.")
		return "bad state", "bad state"
	}

	args := strings.Fields(command)
	if len(args) == 0 {
		fmt.Println("no commands passed")
		return "invalid", "invalid"
	}
	
	area := args[1]

	switch area {
		case "Town":
			fmt.Println("you travel to the town!") 
		case "Plains":
			fmt.Println("you travel to the plains!") 
		case "Forest":
			fmt.Println("you travel to the forest!") 
		case "Cave":
			fmt.Println("you travel to the cave!") 
		case "Dungeon":
			fmt.Println("you travel to the dungeon!") 
		default:
			fmt.Println("Invalid location!")
	}

	message:= "Move area: " + username + " to " + area
	fmt.Printf("Fields are: %q", args)

	prompt := "You collect your belongings, and set off for the " + area
	return message, prompt
}


func store(pc *userStats, command string) ( string, string ) {
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
		return "invalid", "invalid"
	}
	
	verb := args[1]
	item := args[2]
	count := "3"
//	amount, e1 := strconv.Atoi(count);
//
//	if e1 != nil {
//		amount := 1	
//	}

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

	fmt.Printf("Fields are: %q", args)

	return message, prompt
}

func item(username string, command string) ( string, string ) {
	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	args := strings.Fields(command)
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		return "invalid", "invalid"
	}
	prompt := "You use a " + args[2]
	message:= "Used item: " + args[1]

	fmt.Printf("Fields are: %q", args)

	return message, prompt
}

func combat(username string, command string) (string, string) {
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
		return "invalid", "invalid"
	}

	message:= "Player is in combat:"
	prompt := "You attack the monster"
	fmt.Printf("Fields are: %q", args)

	return message, prompt

}

// internal function to determine if teamwork should be used in combat

func inTeam(username string) string {
	// is the user in a team currently?

	message:= "Team check: user checked: " + username

	return message
}


func getState(username string, command string) ( string, string ) {
	//Is the user in town?  Is the user in combat?

	args := strings.Fields(command)
	fmt.Printf("Fields are: %q", args)
	if len(args) < 1 {
		fmt.Println("not enough arguments")
		return "invalid", "invalid"
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

