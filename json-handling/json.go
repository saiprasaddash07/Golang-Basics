package jsonhandling

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	HairColor string `json:"hairColor"`
	HasDog    bool   `json:"hasDog"`
}

func main() {
	myJson := `
		[
			{
				"firstName" : "X",
				"lastName" : "Y",
				"hairColor" : "Z",
				"hasDog" : true
			},
			{
				"firstName" : "A",
				"lastName" : "B",
				"hairColor" : "C",
				"hasDog" : false
			}
		]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling the json", err)
	}

	log.Println(unmarshalled)
}
