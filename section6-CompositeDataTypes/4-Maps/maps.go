package __Maps

import "fmt"

type person struct {
	name string
	age  int
}

func MapExample() {
	myMap1 := make(map[string]string)
	myMap2 := map[string]string{}
	var myMap3 map[string]string = map[string]string{}
	myMap1["key1"] = "value1"
	myMap2["key2"] = "value2"
	myMap3["key3"] = "value3"
	fmt.Println(myMap1) // map[key1:value1]
	fmt.Println(myMap2) // map[key1:value1]
	fmt.Println(myMap3) // map[key1:value1]
	//
	myContacts := make(map[string]person)
	myContacts["001"] = person{"John", 42}
	myContacts["002"] = person{"Jane", 47}
	myContacts["003"] = person{"July", 32}
	fmt.Println(myContacts) // map[001:{John 42} 002:{Jane 47} 003:{July 32}]
	myContacts["001"] = person{"Max", 12}
	fmt.Println(myContacts) // map[001:{Max 12} 002:{Jane 47} 003:{July 32}]
	// len
	fmt.Println(len(myContacts))
	// orders
	for _, contact := range myContacts {
		fmt.Println(contact)
	}
	// for keeping the order we use array
	var orderContacts []string
	orderContacts = append(orderContacts, "001")
	orderContacts = append(orderContacts, "002")
	orderContacts = append(orderContacts, "003")
	for _, contact := range orderContacts {
		fmt.Println(myContacts[contact])
	}
	// how to copy:
	myContacts2 := map[string]person{}
	for key, value := range myContacts {
		myContacts2[key] = value
	}
	// delete
	delete(myContacts, "002")
	fmt.Println(myContacts)
	// get a member
	Max := myContacts["001"]
	fmt.Println(Max)
	July := myContacts["004"]
	fmt.Println(July)
	// get and check existing
	july, ok := myContacts["004"]
	if ok {
		fmt.Println(july)
	} else {
		fmt.Println("not found")
	}
}
