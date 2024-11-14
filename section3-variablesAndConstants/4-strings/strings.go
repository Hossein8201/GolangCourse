package __strings

import (
	"fmt"
	"strings"
)

func StringsAndMethods() {
	name := "hossein"
	studentNumber := 40133014
	myScore := 19.03
	print("my name is ", name, " and my student number is ", studentNumber, " . my average score until now ", myScore, "\n")
	print("my name is", name, "and my student number is", studentNumber, ". my average score until now", myScore, "\n")
	println("my name is ", name, " and my student number is ", studentNumber, " . my average score until now ", myScore)
	println("my name is", name, "and my student number is", studentNumber, ". my average score until now", myScore)
	//
	fmt.Printf("my name is %s and my student number is %d. my average score until now %f\n", name, studentNumber, myScore)
	fmt.Printf("the type of name is %T", name)
	fmt.Printf("the binary of studentNumber is %b\n", studentNumber)
	//
	fmt.Println("the string function : ")
	string := "this is a go test to learn better go. and be a go programmer."
	fmt.Println(strings.Contains(string, "go"))
	fmt.Println(strings.Contains(string, "go1"))
	fmt.Println(strings.ContainsAny(string, "go2"))
	fmt.Println(strings.ContainsAny(string, "q3"))

	fmt.Println(strings.Count(string, "go"))
	fmt.Println(strings.Count(string, "go1"))

	fmt.Println(strings.Cut(string, "go"))
	fmt.Println(strings.Cut(string, "go1"))

	words := strings.Split(string, " ")
	fmt.Println("the number of words :", len(words), "the words :", words)
	fmt.Println(strings.Join(words, ","))

	fmt.Println(strings.Repeat("Iran ", 10))
	fmt.Println(strings.Replace(string, "go", "golang", 1))
	fmt.Println(strings.ReplaceAll(string, "go", "golang"))

	fmt.Println(strings.Compare("golang", "golang")) // 0
	fmt.Println(strings.Compare("golang", "Golang")) // 1
	fmt.Println(strings.Compare("Golang", "golang")) // -1

	fmt.Println(strings.EqualFold("golAng", "Golang")) // True
	fmt.Println(strings.EqualFold("golAn", "Golang"))  //False

	fmt.Println(strings.HasPrefix("golAng", "go"))  // True
	fmt.Println(strings.HasPrefix("golAng", "go1")) //False

	fmt.Println(strings.HasSuffix("golAng", "Ang")) // True
	fmt.Println(strings.HasSuffix("golAng", "g"))   // True

	fmt.Println(strings.Index("Golang", "la"))          // 2
	fmt.Println(strings.Index("Golang", "l"))           // 2
	fmt.Println(strings.LastIndex("Golang land", "la")) // 7

	fmt.Println(strings.ToLower(string))
	fmt.Println(strings.ToUpper(string))
	fmt.Println(strings.Title(string))

	fmt.Println(strings.Trim("???golang language???", "?"))
	fmt.Println(strings.TrimLeft("???golang language???", " ?"))
	fmt.Println(strings.TrimRight("???golang language???", " ?"))

}
