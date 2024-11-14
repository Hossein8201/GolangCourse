package main

import (
	ifElse "ConditionalStatements/1-If-Else"
	switchCase "ConditionalStatements/2-SwitchCase"
	breakFallthrough "ConditionalStatements/3-BreakFallthrough"
	"fmt"
)

func main() {
	ifElse.IfElseStatement()
	fmt.Println()
	switchCase.SwitchCaseStatement()
	fmt.Println()
	breakFallthrough.BreakFallthroughStatement()
}
