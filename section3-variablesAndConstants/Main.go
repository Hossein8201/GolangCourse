package main

import (
	"fmt"
	__defineVariables "variablesAndConstants/1-defineVariables"
	__Consts "variablesAndConstants/2-Consts"
	__scopes "variablesAndConstants/3-scopes"
	__strings "variablesAndConstants/4-strings"
)

func main() {
	__defineVariables.DefineVariables()
	fmt.Println()
	__Consts.Constants()
	fmt.Println()
	__scopes.Scope()
	fmt.Println()
	__strings.StringsAndMethods()
}
