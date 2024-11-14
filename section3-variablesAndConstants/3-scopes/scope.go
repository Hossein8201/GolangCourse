package __scopes

var globalVariable = 125

func Scope() {
	println("In the global Scope : ", globalVariable)
	{
		globalVariable := 1
		println("In the local Scope : ", globalVariable)
	}
	println("In the global Scope : ", globalVariable)
}
