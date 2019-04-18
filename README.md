# process


> Purpose

This framework has the purpose of simplifying the execution of processes and the evaluation of errors


> Description

 this framework allows to define a process as a set of steps that are executed until one of them fails. If one step fails the rest of the steps are not invoked


 > Example

 This Code (without using process)

```Go
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


```

Can be replaced by this

 ```Go

func main() {

	//Create Process
	proc := createProcess()

	//Init process
	person := &Person{Name: "Name", Age: 20}

    //Execute the process
    //RullAll execute Validate,Create and Audit
	if proc.Start(person).RunAll().Error() != nil {
		//there are mistakes?
		fmt.Println(proc.Error())
	} else {
		//Process finish Ok
		person := proc.Result().(*Person)
		fmt.Println(person)
	}

}


 ```