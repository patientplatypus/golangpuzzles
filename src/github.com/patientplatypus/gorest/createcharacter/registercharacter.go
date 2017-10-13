package createcharacter

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/patientplatypus/gorest/config"
)

func RegisterCharacter(charactername string) {
	log.Print("inside register character")
	log.Print("Charactername:", charactername)

	log.Print("Waiting for create table")
	finished := make(chan bool)
	go Create_Table(finished)
	<-finished
	log.Print("Continuing with the rest of the program")

	var cnt int
	config.DB.QueryRow(`select count(*) from usercharacters`).Scan(&cnt)
	log.Print("count: ", cnt)
	if cnt == 0 {
		log.Print("username is not in usercharacters adding now")

		usercharacter := UserCharacters{
			Character: []ACharacter{
				{CharactersName: charactername},
			}}

		datacharacter, err := json.Marshal(usercharacter)
		if err != nil {
			fmt.Println("Oh no! There was an error:", err)
			return
		}

		config.DB.QueryRow("insert into usercharacters (usernames, characters) values ($1, $2);", Username, datacharacter)
	} else {
		var databasecharacters UserCharacters
		var datarows []byte
		config.DB.QueryRow("SELECT characters FROM usercharacters where usernames = $1", Username).Scan(&datarows)
		json.Unmarshal(datarows, &databasecharacters)
		log.Print("Here are the usernames in the database:", databasecharacters)

		var namealreadyused bool

		for j := 0; j < len(databasecharacters.Character); j++ {
			if databasecharacters.Character[j].CharactersName == charactername {
				namealreadyused = true
			}
			if j == len(databasecharacters.Character) && namealreadyused != true {
				namealreadyused = false
			}
		}
		if namealreadyused != true {
			newcharacter := ACharacter{
				CharactersName: charactername,
			}
			databasecharacters.Character = append(databasecharacters.Character, newcharacter)

			appendcharacter, err := json.Marshal(databasecharacters)
			if err != nil {
				fmt.Println("Oh no! There was an error:", err)
				return
			}

			config.DB.Query("UPDATE usercharacters SET characters=$1 where usernames=$2;", appendcharacter, Username)
		} else {
			log.Print("Character name already in use!")
		}

	}

	//   UPDATE Customers
	// SET ContactName = 'Alfred Schmidt', City= 'Frankfurt'
	// WHERE CustomerID = 1;

	// usercharacter := UserCharacters{
	// 	Character: []ACharacter{
	// 		{CharactersName: "John"},
	// 	}}
	//
	// newcharacter := ACharacter{
	// 	CharactersName: "Dave",
	// }
	//
	// usercharacter.Character = append(usercharacter.Character, newcharacter)

	// log.Print(usercharacter)
}
