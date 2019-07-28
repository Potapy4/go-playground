package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {
	myUser := user{
		name:  "Nik",
		email: "test@email.com",
		age:   22,
	}

	fmt.Println("Name:", myUser.name)
	fmt.Println("Email:", myUser.email)
	fmt.Println("Age:", myUser.age)

	anonymousStruct := struct {
		name  string
		email string
		age   int64
	}{
		name:  "ABC",
		email: "test@test.test",
		age:   18,
	}

	fmt.Println("Anonymous name:", anonymousStruct.name)
	fmt.Println("Anonymous email:", anonymousStruct.email)
	fmt.Println("Anonymous age:", anonymousStruct.age)
}
