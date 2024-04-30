package gamecore
import("fmt")

func main(){
	fmt.Println("gamecore init")

}

// functions
	// command parser
		// playerID, command, arguments
		//check if open session


func create_player(user fstring){
	//player creation
	message:= "Create player: UserID passed in: " + user

	return message

}

func get_status(user fstring){
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started
	
		decide the format for this state tracking 
	*/

	message:= "Get status: " + user

	return message

}

func move_area(user fstring, area fstring){
	// navigation

	/*
		- check if the player can move (stuck in combat? middle of player creation or exchange?)
		- move to the selected area
		- provide flavor text for the area
	*/
	message:= "Move area: " + user + " to " + area

	return message

}

func store(user fstring, action fstring, count int){
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/

	message:= "Store menu: Action: " + action + "count: " + count

	return message		
}

func item(user fstring, action fstring, count int){
	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/
	message:= "Used item: " + action 

	return message
}

func combat(user fstrung, action fstring, target int){
	// fight

	/*
		- check if the player is in a valid state
		- search for an enemy
		- set the combat state
		- advance combat as the player selects the attack method
		- update the player with the turn results
		- complete combat and award experience
	*/
	message:= "Player is in combat: Player: " + user + " Attack Action: " + action + " Target: " + target

	return message

}


func in_team(user fstring){
	// is the user in a team currently?

		message:= "Team check: " + "user checked: " + user

	return message
}


func get_state(user fstring){
	//Is the user in town?  Is the user in combat?

	message:= "Check use state: " + user

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

