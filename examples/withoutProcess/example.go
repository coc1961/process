package main

import (
	"errors"
	"fmt"
)

//Person test struct
type Person struct {
	Name string
	Age  int
}

func validate(person *Person) (*Person, error) {
	if person.Name == "" {
		return nil, errors.New("Name is Null")
	}
	return person, nil
}
func create(person *Person) (*Person, error) {
	// Save person in database
	return person, nil
}

func audit(person *Person) (*Person, error) {
	// Audit
	return person, nil
}

func main() {

	person := &Person{Name: "Name", Age: 20}
	var err error

	if person, err = validate(person); err != nil {
		fmt.Println(err)
		return
	}

	if person, err = create(person); err != nil {
		fmt.Println(err)
		return
	}

	if person, err = audit(person); err != nil {
		fmt.Println(err)
		return
	}
	//Ok!!
	fmt.Println(person)
}
