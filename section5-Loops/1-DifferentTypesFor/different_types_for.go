package __DifferentTypesFor

func DifferentForLoop() {
	//
	for {
		print("hello word\n")
		break
	}
	//
	i := 0
	for i < 10 {
		println("Hi hossein", i)
		i++
	}
	//
	for j := 0; j < 10; j++ {
		println("Hi hossein", j)
	}
	//
	listNum := []int{11, 12, 13, 14, 15}
	for index, item := range listNum {
		println(index, item)
	}
}
