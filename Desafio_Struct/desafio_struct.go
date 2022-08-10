// faça duas Struct, uma usando "herança"

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Carros struct {
	Marca  string  `json:"Marca"`
	Modelo string  `json:"Modelo"`
	Valor  float64 `json:"Valor"`
	Turbo  bool    `json:"Turbo"`
}

type Caminhonete struct {
	Carros
	Cabine    string `json:"Cabine"`
	Importada bool   `json:"Importada"`
}

func main() {

	Cruze := Carros{"Chevrolet", "Cruze Ltz", 98.060, true}

	Ranger := Caminhonete{
		Carros:    Carros{"Ford", "Ranger", 130.980, true},
		Cabine:    "Dupla",
		Importada: false,
	}

	cruzeJson, err := json.Marshal(Cruze)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cruzeJson))

	rangerJson, err := json.Marshal(Ranger)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rangerJson))

	fmt.Printf("Fizemos 2 structs %v, e %v", Cruze, Ranger)

}
