package __Defer

import (
	"fmt"
	"io"
	"os"
)

func Defer() {
	defer fmt.Print("Bye\n")
	defer fmt.Print("\nGood ")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("Hello and ...")
}

func CopyFile(destination, start string) (int64, error) {
	startFile, err := os.Open(start)
	if err != nil {
		return 0, err
	}
	defer startFile.Close()

	endFile, err := os.Create(destination)
	if err != nil {
		return 0, err
	}
	defer endFile.Close()

	return io.Copy(startFile, endFile)
}
