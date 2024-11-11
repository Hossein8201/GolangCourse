# Lecture One : Get ready

> If you want to __get ready__ for this course, please [follow this link](https://github.com/naeemaei/Golang-Tutorial)

to run a go file use : 
> go run file_name.go

## Hints of packages and modules :
- You must import "fmt" for printing and other.
- the next step is functions.
- Remember the "init function" calls before "main function".
- You must __import__ a package in another file source if you want to use it.
  > import (package_name "path of that in your computer")
- packages is a set of files and modules is a set of packages.
- In modules and packages, you must __publish first your directory__ and then use it. 
  - `install module and import package`
- You can create a module use of :
  > go mod init ExampleModule

## Go CLI (command line interface) :
- `go version`
- `go help`
  - Show us all commands and their applications.
- `go bug`
  - Create a new page in google, to import a new bug in `golang repository`
- `go build file_name.go`
  - Create a `.exe` from file.go
- `go clean file_name.go`
  - Delete a `.exe` for file.go
- `go doc package_name`
  - Give us very information about the package
- `go mod module_name`
  - Create a module
- `go run file_name.go`
  - Run file
- `go install`
  - Compile and install packages and dependencies
- `go get`
  - Add dependencies to current module and install them
- `go list`
  - List packages or modules