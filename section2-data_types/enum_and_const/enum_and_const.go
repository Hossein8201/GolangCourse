package enum_and_const

import (
	"encoding/json"
	"fmt"
)

type state struct {
	number       int
	presentState status
}

// Enum:
type status int

const (
	status1   status = 1
	pay       status = 2
	strongest status = 3
	status4   status = 4
	status5   status = 5
)

// Const:
const author = "hossein Tatar"
const old = 20

func GiveInformation() *map[string]int {
	state1 := state{number: 1, presentState: pay}
	state2 := state{number: 2, presentState: strongest}
	fmt.Println(state1)

	stateJson1, _ := json.Marshal(state1)
	stateJson2, _ := json.Marshal(state2)

	fmt.Println(string(stateJson1))
	fmt.Println(string(stateJson2))
	fmt.Println()
	fmt.Println(status1)
	fmt.Println(pay)
	fmt.Println(strongest)
	fmt.Println(status4)
	fmt.Println(status5)
	fmt.Println()
	//author = "ali"    // Can't do this
	const favorite = "apple"
	const much = 75
	//favorite = "banana"	// Can't do this
	map1 := map[string]int{author: old, favorite: much}
	fmt.Println(map1)
	return &map1
}
