package gamecore

import (
	"strconv"
)

/*
func main() {
	user := "Ham"
	message := create_player(user)
	fmt.Println(message)
}

*/

// functions
// command parser
// playerID, command, arguments
//check if open session

func create_player(user string) string {
	//player creation
	message := "Create player: UserID passed in: " + user

	return message

}

func get_status(user string) string {
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started

		decide the format for this state tracking
	*/

	message := "Get status: " + user

	return message

}

func move_area(user string, area string) string {
	// navigation

	/*
		- check if the player can move (stuck in combat? middle of player creation or exchange?)
		- move to the selected area
		- provide flavor text for the area
	*/
	message := "Move area: " + user + " to " + area

	return message

}

func store(user string, action string, count int) string {
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/

	message := "Store menu: Action: " + action + "count: " + strconv.Itoa(count)

	return message
}

func item(user string, action string, count int) string {
	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	message := "Used item: " + action

	return message
}

func combat(user string, action string, target int) string {
	// fight

	/*
		- check if the player is in a valid state
		- search for an enemy
		- set the combat state
		- advance combat as the player selects the attack method
		- update the player with the turn results
		- complete combat and award experience
	*/
	message := "Player is in combat: Player: " + user + " Attack Action: " + action + " Target: " + strconv.Itoa(target)

	return message

}

func in_team(user string) string {
	// is the user in a team currently?

	message := "Team check: " + "user checked: " + user

	return message
}

func get_state(user string) string {
	//Is the user in town?  Is the user in combat?

	message := "Check use state: " + user

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
