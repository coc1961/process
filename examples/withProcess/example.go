package main

import (
	"errors"
	"fmt"

	"github.com/coc1961/process"
)

//Person test struct
type Person struct {
	Name string
	Age  int
}

func validate(p process.Context) (interface{}, error) {
	person := p.(*Person)
	if person.Name == "" {
		return nil, errors.New("Name is Null")
	}
	return person, nil
}
func create(p process.Context) (interface{}, error) {
	person := p.(*Person)
	// Save person in database
	return person, nil
}

func audit(p process.Context) (interface{}, error) {
	person := p.(*Person)
	// Audit
	return person, nil
}

func createProcess() *process.Process {
	// Create and init Process
	proc := process.New()
	proc.AddStep(validate)
	proc.AddStep(create)
	proc.AddStep(audit)

	return proc
}

func main() {

	//Create Process
	proc := createProcess()

	//Init process
	person := &Person{Name: "Name", Age: 20}

	//Execute the process
	if proc.Start(person).RunAll().Error() != nil {
		//there are mistakes?
		fmt.Println(proc.Error())
	} else {
		//Process finish Ok
		person := proc.Result().(*Person)
		fmt.Println(person)
	}

}
