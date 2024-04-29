package gamecore
import("fmt")

func main(){
	fmt.Println("Yo!")



}

// functions
	// command parser
		// playerID, command, arguments
		//check if open session

func check_session(user fstring){

}

func parse_command(cmd fstring){
	//parse command
	switch cmd {
	case "attack", "fight":
		combat()
	case "store", "shop", "vendor":
		store()
//	case "move":
//		store()
//	case "status":
//		store()
//	case "item":
//		store()
//	case "store":
//		store()
//	case "store":
//		store()
//	case "store":
//		store()
	default:
		fmt.Println("this is not a valid command!")
	}
}

func create_player(user fstring){
	//player creation

}

func get_status(user fstring){
	// player details

	/*
		lookup the player and get their user state. return the action they are currently in the middle of.
		0 = does not exist
		1 = player creation started
	
		decide the format for this state tracking 
	*/

}

func move_area(user fstring, area fstring){
	// navigation

	/*
		- check if the player can move (stuck in combat? middle of player creation or exchange?)
		- move to the selected area
		- provide flavor text for the area
	*/
}

func store(user fstring, action fstring, count int){
	// store

	/*
		- check if the player is in a valid state
		- determine if the request is valid
		- update inventory
		- complete the transaction and notify the player
	*/		
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
}


func in_team(user fstring){
	// is the user in a team currently?
}


func get_state(user fstring){
	//Is the user in town?  Is the user in combat?
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

