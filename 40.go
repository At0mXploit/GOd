package main

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	user := User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
		},
	}
	
	if membershipType == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}
	
	return user
}
