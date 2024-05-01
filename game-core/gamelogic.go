package main
//package gamecore
import(
	"fmt"
	"strconv"
	"strings"
)


func main() {
	user := "Ham"
	command := "shop buy potion 10"
	message := store(user, command)
	fmt.Println("\n" + message)

	/*
	checkstate(user string)
	create_player(user string)
	get_status(user string, command string)
	move_area(user string, command string)
	store(user string, command string)
	item(user string, action string, count int)
	combat(user string, action string, target int)
	in_team(user string)
	get_state(user string)
	*/
}



// functions
	// command parser
		// playerID, command, arguments
		//check if open session


func checkstate(user string) bool {

	if user == "Ham" {
		return true
	} else {
		return false
	}
}

func create_player(user string) string {
	//player creation
	message:= "Create player: UserID passed in: " + user

	return message

}

func get_status(user string, command string) string {
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started
	
		decide the format for this state tracking 
	*/

	err := checkstate(user)

	if err == false {
		fmt.Println("User unable to do this command at this time.")
		return "bad state"
	}

	args := strings.Fields(command)

	if len(args) == 0 {
		fmt.Println("no commands passed")
		return "invalid"
	}
	
	switch verb := args[1]; verb {
		case "health":
			fmt.Println("you checked your health!") 
		default:
			fmt.Println("you didn't check anything!")
	}
	
	message:= "Get status: " + user
	
	fmt.Printf("Fields are: %q", args)

	return message
}


func move_area(user string, command string) string {
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

	err := checkstate(user)
	if err == false {
		fmt.Println("User unable to do this command at this time.")
		return "bad state"
	}

	args := strings.Fields(command)
	if len(args) == 0 {
		fmt.Println("no commands passed")
		return "invalid"
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

	message:= "Move area: " + user + " to " + area
	fmt.Printf("Fields are: %q", args)

	return message
}


func store(user string, command string) string {
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	err := checkstate(user)
	if err == false {
		fmt.Println("User unable to do this command at this time.")
		return "bad state"
	}

	args := strings.Fields(command)
	if len(args) < 3 {
		fmt.Println("not enough arguments")
		return "invalid"
	}
	
	verb := args[1]
	item := args[2]
	count := "3"
	amount, e1 := strconv.Atoi(count);

	if e1 != nil {
		amount := 1	
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

	message:= "Store menu: Action: " + verb + "item:" + item + "count: " + args[3]
	fmt.Printf("Fields are: %q", args)

	return message
}

func item(user string, action string, count int) string {
	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	message:= "Used item: " + action
	fmt.Printf("Fields are: %q", args)

	return message
}

func combat(user string, command string) string {
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

	message:= "Player is in combat:"
	fmt.Printf("Fields are: %q", args)

	return message

}


func in_team(user string) string {
	// is the user in a team currently?

		message:= "Team check: " + "user checked: " + user

	return message
}


func get_state(user string, command string) string {
	//Is the user in town?  Is the user in combat?

	message:= "Check use state: " + user
	args := strings.Fields(command)
	fmt.Printf("Fields are: %q", args)

	return message
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

