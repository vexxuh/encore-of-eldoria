package discordbot

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get .env parameters
func getBotParams(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// HTTP request handler
func handleHttpRequest(route string, action string) {

}

/*
{
	"field": "data"
	"attack": "damage"
	"player": "username"
	"verb": "value"
}

1. receive http response back from api
2. serialize into a go-readable datatype -> string (stringify js)
JSON.stringify => { "string": "string" }
3. serialized_json.map((verb, value) => player -> discord.mesasge.senderID + verb + value)
4. build 'base syntax' of how the bot will relay API responses -> map those responses agnostic of the actual text
i.e.
you BUY an ITEM for X VALUE
you PERFORM AN ACTION with WEAPON for X DAMAGE
you TRAVEL to a LOCATION in PLACE (idk)

ALTERNATIVELY
novel ai will, based on values passed from discord user input, be prompted with a schema on David's end,
and all we pass back is the text response from novelai

discord -> /attack enemy: wolf with: sword
david api -> game logic performed on input -> generates values -> result of action
result of action -> passed to novelai -> returns a player state and result of action with lore text
-> passed BACK to http handler { "result": "You attacked a wolf with a sword for 11 dmg. You stand triumphant", "image": "url.com"}
-> this response is kicked back to the discord response handler to simply inject this text into the bot's message
*/
