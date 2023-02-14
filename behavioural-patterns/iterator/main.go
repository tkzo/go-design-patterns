package main

import "fmt"

func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}
	user3 := &User{
		name: "c",
		age:  18,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2, user3},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
